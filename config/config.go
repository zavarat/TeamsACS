package config

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"

	"gopkg.in/yaml.v2"

	"github.com/ca17/teamsacs/common"
)

type MongodbConfig struct {
	Url    string `yaml:"url" json:"url"`
	User   string `yaml:"user" json:"user"`
	Passwd string `yaml:"passwd" json:"passwd"`
}

type SysConfig struct {
	Appid      string `yaml:"appid" json:"appid"`
	Workdir    string `yaml:"workdir" json:"workdir"`
	SyslogAddr string `yaml:"syslog_addr" json:"syslog_addr"`
	Location   string `yaml:"location" json:"location"`
	Aeskey     string `yaml:"aeskey" json:"aeskey"`
	Debug      bool   `yaml:"debug" json:"debug"`
}

type WebConfig struct {
	Host      string `yaml:"host" json:"host"`
	Port      int    `yaml:"port" json:"port"`
	Debug     bool   `yaml:"debug" json:"debug"`
	JwtSecret string `yaml:"jwt_secret" json:"jwt_secret"`
}

type GrpcConfig struct {
	Host  string `yaml:"host" json:"host"`
	Port  int    `yaml:"port" json:"port"`
	Debug bool   `yaml:"debug" json:"debug"`
}

type RadiusdConfig struct {
	Host     string `yaml:"host" json:"host"`
	AuthPort int    `yaml:"auth_port" json:"auth_port"`
	AcctPort int    `yaml:"acct_port" json:"acct_port"`
	Debug    bool   `yaml:"debug" json:"debug"`
}

type AppConfig struct {
	System  SysConfig     `yaml:"system" json:"system"`
	Web     WebConfig     `yaml:"web" json:"web"`
	Mongodb MongodbConfig `yaml:"mongodb" json:"mongodb"`
	Grpc    GrpcConfig    `yaml:"grpc" json:"grpc"`
	Radiusd RadiusdConfig `yaml:"radiusd" json:"radiusd"`
}

func (c *AppConfig) GetLogDir() string {
	return path.Join(c.System.Workdir, "logs")
}

func (c *AppConfig) GetDataDir() string {
	return path.Join(c.System.Workdir, "data")
}

func (c *AppConfig) GetRadiusDir() string {
	return path.Join(c.System.Workdir, "radius")
}

func (c *AppConfig) GetPrivateDir() string {
	return path.Join(c.System.Workdir, "private")
}

func (c *AppConfig) GetResourceDir() string {
	return path.Join(c.System.Workdir, "resource")
}

func (c *AppConfig) GetBackupDir() string {
	return path.Join(c.System.Workdir, "backup")
}

func (c *AppConfig) InitDirs() {
	os.MkdirAll(path.Join(c.System.Workdir, "logs"), 0700)
	os.MkdirAll(path.Join(c.System.Workdir, "radius"), 0700)
	os.MkdirAll(path.Join(c.System.Workdir, "data"), 0700)
	os.MkdirAll(path.Join(c.System.Workdir, "public"), 0700)
	os.MkdirAll(path.Join(c.System.Workdir, "private"), 0700)
	os.MkdirAll(path.Join(c.System.Workdir, "resource"), 0700)
	os.MkdirAll(path.Join(c.System.Workdir, "backup"), 0644)
}

var DefaultAppConfig = &AppConfig{
	System: SysConfig{
		Appid:      "TeamsACS",
		Workdir:    "/var/teamsacs",
		SyslogAddr: "",
		Location:   "Asia/Shanghai",
		Aeskey: "5f8923be3da19452d3acdc9e69fa24e6",
	},
	Web: WebConfig{
		Host:      "0.0.0.0",
		Port:      18998,
		Debug:     true,
		JwtSecret: "9b6de5cc07384bf1acs10f568ac9da37",
	},
	Grpc: GrpcConfig{
		Host:  "0.0.0.0",
		Port:  18999,
		Debug: true,
	},
	Radiusd: RadiusdConfig{
		Host:     "0.0.0.0",
		AuthPort: 1812,
		AcctPort: 1813,
		Debug:    true,
	},
	Mongodb: MongodbConfig{
		Url:    "mongodb://127.0.0.1:27017",
		User:   "",
		Passwd: "",
	},
}

func setEnvValue(name string, f func(v string)) {
	var evalue = os.Getenv(name)
	if evalue != "" {
		f(evalue)
	}
}
func setEnvInt64Value(name string, f func(v int64)) {
	var evalue = os.Getenv(name)
	if evalue == "" {
		return
	}
	p, err := strconv.ParseInt(evalue, 10, 64)
	if err == nil {
		f(p)
	}
}

func LoadConfig(cfile string) *AppConfig {

	cfg := new(AppConfig)
	if common.FileExists(cfile) {
		data := common.Must2(ioutil.ReadFile(cfile))
		common.Must(yaml.Unmarshal(data.([]byte), cfg))
	} else {
		cfg = DefaultAppConfig
	}

	setEnvValue("TEAMSACS_WORKER_DIR", func(v string) {
		cfg.System.Workdir = v
	})
	setEnvValue("TEAMSACS_WEB_HOST", func(v string) {
		cfg.Web.Host = v
	})
	setEnvValue("TEAMSACS_WEB_DEBUG", func(v string) {
		cfg.Web.Debug = v == "true"
	})
	setEnvValue("TEAMSACS_WEB_SECRET", func(v string) {
		cfg.Web.JwtSecret = v
	})
	setEnvInt64Value("TEAMSACS_WEB_PORT", func(v int64) {
		cfg.Web.Port = int(v)
	})

	setEnvValue("TEAMSACS_MONGODB_URL", func(v string) {
		cfg.Mongodb.Url = v
	})
	setEnvValue("TEAMSACS_MONGODB_USER", func(v string) {
		cfg.Mongodb.User = v
	})
	setEnvValue("TEAMSACS_MONGODB_PASSWD", func(v string) {
		cfg.Mongodb.Passwd = v
	})

	setEnvValue("TEAMSACS_GRPC_HOST", func(v string) {
		cfg.Grpc.Host = v
	})
	setEnvInt64Value("TEAMSACS_GRPC_PORT", func(v int64) {
		cfg.Grpc.Port = int(v)
	})

	setEnvInt64Value("TEAMSACS_GRPC_DEBUG", func(v int64) {
		cfg.Grpc.Debug = v == 1
	})

	setEnvInt64Value("TEAMSACS_RADIUS_AUTH_PORT", func(v int64) {
		cfg.Radiusd.AuthPort = int(v)
	})

	setEnvInt64Value("TEAMSACS_RADIUS_ACCT_PORT", func(v int64) {
		cfg.Radiusd.AcctPort = int(v)
	})

	setEnvInt64Value("TEAMSACS_RADIUS_DEBUG", func(v int64) {
		cfg.Radiusd.Debug = v == 1
	})

	return cfg
}
