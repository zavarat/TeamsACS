package web

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	elog "github.com/labstack/gommon/log"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/log"
	"github.com/ca17/teamsacs/common/static"
	"github.com/ca17/teamsacs/common/tpl"
	"github.com/ca17/teamsacs/models"
)

// 运行管理系统
func ListenAdminServer(manager *models.ModelManager) error {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(ServerRecover(manager.Config.Web.Debug))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "admin ${time_rfc3339} ${remote_ip} ${method} ${uri} ${protocol} ${status} ${id} ${user_agent} ${latency} ${bytes_in} ${bytes_out} ${error}\n",
		Output: os.Stdout,
	}))
	manager.WebJwtConfig = &middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey:    []byte(manager.Config.Web.JwtSecret),
		Skipper: func(c echo.Context) bool {
			skips := []string{"", "/", "/login", "/verifymfa", "/opr/verifymfa"}
			if common.InSlice(c.Request().RequestURI, skips) ||
				strings.HasPrefix(c.Path(), "/static") ||
				strings.HasSuffix(c.Path(), ".css") ||
				strings.HasSuffix(c.Path(), ".css") ||
				strings.HasSuffix(c.Path(), ".gif") ||
				strings.HasSuffix(c.Path(), ".png") {
				return true
			}
			return false
		},
		ErrorHandler: func(err error) error {
			return NewHTTPError(http.StatusBadRequest, "Missing tokens, limited access to resources")
		},
	}
	e.Use(middleware.JWTWithConfig(*manager.WebJwtConfig))

	// Init Handlers
	InitAllRouter(e)

	manager.TplRender = tpl.NewCommonTemplate([]string{"/resources/templates"}, manager.Dev, manager.GetTemplateFuncMap())
	e.Renderer = manager.TplRender
	if manager.Dev {
		e.Static("/static", "static")
	} else {
		e.GET("/static/*", echo.WrapHandler(http.FileServer(static.FS(false))))
	}

	e.Pre(middleware.Rewrite(map[string]string{
		"/freeradius/*": "/api/freeradius/$1",
		"/favicon.ico":  "/static/favicon.ico",
	}))

	e.HideBanner = true
	e.Logger.SetLevel(common.If(manager.Config.Web.Debug, elog.DEBUG, elog.INFO).(elog.Lvl))
	e.Debug = manager.Config.Web.Debug
	log.Info("try start tls web server")
	err := e.StartTLS(fmt.Sprintf("%s:%d", manager.Config.Web.Host, manager.Config.Web.Port),
		path.Join(manager.Config.GetPrivateDir(), "teamsacs-web.tls.crt"), path.Join(manager.Config.GetPrivateDir(), "teamsacs-web.tls.key"))
	if err != nil {
		log.Warningf("start tls server error %s", err)
		log.Info("start web server")
		err = e.Start(fmt.Sprintf("%s:%d", manager.Config.Web.Host, manager.Config.Web.Port))
	}
	return err
}

func ServerRecover(debug bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					if debug {
						log.Errorf("%+v", r)
					}
					c.Error(echo.NewHTTPError(http.StatusInternalServerError, err.Error()))
				}
			}()
			return next(c)
		}
	}
}
