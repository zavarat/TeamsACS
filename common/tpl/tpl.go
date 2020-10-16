package tpl

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/labstack/echo/v4"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/log"
	"github.com/ca17/teamsacs/common/resources"
)

type CommonTemplate struct {
	Templates *template.Template
}

func NewCommonTemplate(dirs []string, devmode bool, funcMap map[string]interface{}) *CommonTemplate {
	var templates = template.New("GlobalTemplate").Funcs(funcMap)
	var ct = &CommonTemplate{Templates: templates}
	for _, d := range dirs {
		ct.parseDir(d, devmode)
	}
	return ct
}

func (ct *CommonTemplate) parseDir(dir string, devmode bool) {
	if devmode {
		dir2 := strings.TrimLeft(dir, "/")
		fss := common.Must2(ioutil.ReadDir(dir2)).([]os.FileInfo)
		for _, item := range fss {
			if item.IsDir() {
				continue
			}
			c, err := ioutil.ReadFile(path.Join(dir2, item.Name()))
			if err == nil {
				ct.parseItem(item, c, ct.Templates)
			}
		}
	} else {
		fs := resources.FS(devmode)
		fd := common.Must2(fs.Open(dir)).(http.File)
		fss := common.Must2(fd.Readdir(0)).([]os.FileInfo)
		for _, item := range fss {
			if item.IsDir() {
				continue
			}
			c, err := resources.FSByte(false, path.Join(dir, item.Name()))
			if err == nil {
				ct.parseItem(item, c, ct.Templates)
			}
		}
	}
}

func (ct *CommonTemplate) parseItem(item os.FileInfo, c []byte, templates *template.Template) {
	name := strings.TrimSuffix(item.Name(), path.Ext(item.Name()))
	if templates.Lookup(name) != nil {
		return
	}
	tplstr := fmt.Sprintf(`{{define "%s"}}%s{{end}}`, name, c)
	ct.Templates = template.Must(templates.Parse(tplstr))
	if log.IsDebug() {
		log.Debugf("parse template %s", name)
	}
}

func (ct *CommonTemplate) ParseDir(dir string, nameprifix string) error {
	fss, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, item := range fss {
		if item.IsDir() {
			continue
		}
		c, err := ioutil.ReadFile(path.Join(dir, item.Name()))
		if err != nil {
			log.Error(err)
			continue
		}
		name := nameprifix + strings.TrimSuffix(item.Name(), path.Ext(item.Name()))
		tplstr := fmt.Sprintf(`{{define "%s"}}%s{{end}}`, name, c)
		t, err := ct.Templates.Parse(tplstr)
		if err != nil {
			log.Error(err)
			continue
		}
		ct.Templates = t
		return nil
	}
	return nil
}

func (t *CommonTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
