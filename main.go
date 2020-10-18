package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/op/go-logging"
	"golang.org/x/sync/errgroup"

	"github.com/ca17/teamsacs/common/installer"
	"github.com/ca17/teamsacs/common/log"
	"github.com/ca17/teamsacs/config"
	"github.com/ca17/teamsacs/grpcservice"
	"github.com/ca17/teamsacs/models"
	"github.com/ca17/teamsacs/radiusd"
	"github.com/ca17/teamsacs/radiusd/radlog"
	"github.com/ca17/teamsacs/web"
)

var (
	g errgroup.Group

	BuildVersion   string
	ReleaseVersion string
	BuildTime      string
	BuildName      string
	CommitID       string
	CommitDate     string
	CommitUser     string
	CommitSubject  string
)

//go:generate esc -o common/static/static.go -pkg static -ignore=".DS_Store,webix_debug.*,webix.js" static
//go:generate esc -o common/resources/resources.go -pkg resources -ignore=".DS_Store" resources
//go:generate protoc -I ./grpcservice --go_out=plugins=grpc:./grpcservice  ./grpcservice/service.proto

// Command line definition
var (
	h          = flag.Bool("h", false, "help usage")
	showVer    = flag.Bool("v", false, "show version")
	debug      = flag.Bool("X", false, "run debug level")
	syslogaddr = flag.String("syslog", "", "syslog addr x.x.x.x:x")
	conffile   = flag.String("c", "/etc/teamsacs.yaml", "config yaml/json file")
	dev        = flag.Bool("dev", false, "run develop mode")
	port       = flag.Int("p", 0, "web port")
	install    = flag.Bool("install", false, "run install")
	uninstall  = flag.Bool("uninstall", false, "run uninstall")
	initcfg    = flag.Bool("initcfg", false, "write default config > /etc/teamsacs.yaml")
	initSuper  = flag.Bool("initsuper", false, "init super password to 'Teams@Acs' ")
)

// Print version information
func PrintVersion() {
	fmt.Fprintf(os.Stdout, "build name:\t%s\n", BuildName)
	fmt.Fprintf(os.Stdout, "build version:\t%s\n", BuildVersion)
	fmt.Fprintf(os.Stdout, "build time:\t%s\n", BuildTime)
	fmt.Fprintf(os.Stdout, "release version:\t%s\n", ReleaseVersion)
	fmt.Fprintf(os.Stdout, "Commit ID:\t%s\n", CommitID)
	fmt.Fprintf(os.Stdout, "Commit Date:\t%s\n", CommitDate)
	fmt.Fprintf(os.Stdout, "Commit Username:\t%s\n", CommitUser)
	fmt.Fprintf(os.Stdout, "Commit Subject:\t%s\n", CommitSubject)
}

func printHelp() {
	if *h {
		ustr := fmt.Sprintf("%s version: %s, Usage:%s -h\nOptions:", BuildName, BuildVersion, BuildName)
		fmt.Fprintf(os.Stderr, ustr)
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func setupAppconfig() *config.AppConfig {
	appconfig := config.LoadConfig(*conffile)
	if *port > 0 {
		appconfig.Web.Port = *port
	}

	if *syslogaddr != "" {
		appconfig.System.SyslogAddr = *syslogaddr
	}
	if *debug {
		appconfig.Web.Debug = *debug
		appconfig.Radiusd.Debug = *debug
		appconfig.Grpc.Debug = *debug
	}
	appconfig.InitDirs()
	return appconfig
}

func setupLogging(appconfig *config.AppConfig) {
	// system logging
	level := logging.INFO
	if appconfig.Web.Debug {
		level = logging.DEBUG
	}
	log.SetupLog(level, appconfig.System.SyslogAddr, appconfig.GetLogDir(), appconfig.System.Appid)

	// radius logging
	radlevel := logging.INFO
	if appconfig.Radiusd.Debug {
		radlevel = logging.DEBUG
	}
	radlog.SetupLog(radlevel, appconfig.System.SyslogAddr, appconfig.GetLogDir(), "Radiusd")

}

func installService(appconfig *config.AppConfig) bool {
	// 安装为系统服务
	if *install {
		err := installer.Install(appconfig)
		if err != nil {
			log.Error(err)
		}
		return true
	}

	// 卸载
	if *uninstall {
		installer.Uninstall()
		return true
	}
	return false
}

func ionitConfig(appconfig *config.AppConfig) bool {
	if *initcfg {
		err := installer.InitConfig(appconfig)
		if err != nil {
			log.Error(err)
		}
		return true
	}
	return false
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	if *showVer {
		PrintVersion()
		os.Exit(0)
	}

	printHelp()

	appconfig := setupAppconfig()

	// set logging level
	setupLogging(appconfig)

	if installService(appconfig) {
		return
	}

	if ionitConfig(appconfig) {
		return
	}


	manager := models.NewModelManager(appconfig, *dev)

	if *initSuper {
		return
	}

	if *dev {
		log.Debug("Running for Dev Mode")
	}

	g.Go(func() error {
		log.Info("Start Radius auth Server ...")
		return radiusd.ListenRadiusAuthServer(manager)
	})

	g.Go(func() error {
		log.Info("Start Radius acct Server ...")
		return radiusd.ListenRadiusAcctServer(manager)
	})

	time.Sleep(time.Millisecond * 50)

	g.Go(func() error {
		log.Info("Start Admin Server ...")
		return web.ListenAdminServer(manager)
	})

	g.Go(func() error {
		log.Info("Start Grpc Server ...")
		return grpcservice.StartGrpcServer(manager)
	})

	time.Sleep(time.Millisecond * 50)

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
