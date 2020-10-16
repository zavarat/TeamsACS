package config

import (
	"gopkg.in/yaml.v2"
	"testing"
)

func TestEncodeYaml(t *testing.T) {
	r, _ := yaml.Marshal(DefaultAppConfig)
	t.Log(string(r))
}

func TestDecodeYaml(t *testing.T) {
	conf := &AppConfig{}
	ystr := `system:
  appid: TeamsACS
  workdir: /var/teamsacs
  syslog_addr: ""
  debug: false
web:
  host: 0.0.0.0
  port: 18998
  debug: true
  jwt_secret: 9b6de5cc-0738-4bf1-acs1-0f568ac9da37
mongodb:
  url: mongodb://127.0.0.1:27017
  user: ""
  passwd: ""
grpc:
  host: 0.0.0.0
  port: 18999
  debug: true
radiusd:
  host: 0.0.0.0
  auth_port: 1812
  acct_port: 1813
  debug: true`
	yaml.Unmarshal([]byte(ystr), conf)

	t.Log(conf)
}
