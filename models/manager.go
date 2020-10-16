package models

import (
	"time"

	"github.com/allegro/bigcache"
	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4/middleware"
	cmap "github.com/orcaman/concurrent-map"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ca17/teamsacs/common"
	"github.com/ca17/teamsacs/common/gmail"
	"github.com/ca17/teamsacs/common/mongodb"
	"github.com/ca17/teamsacs/common/tpl"
	"github.com/ca17/teamsacs/config"
)

const (
	MDBTeamsacs        = "teamsacs"
	MDBGenieacs        = "genieacs"
	TeamsacsConfig     = "config"
	TeamsacsOperator   = "operator"
	TeamsacsSubscribe  = "subscribe"
	TeamsacsOprlog     = "oprlog"
	TeamsacsLdap       = "ldap"
	TeamsacsVpe       = "vpe"
	TeamsacsOnline     = "online"
	TeamsacsProfile    = "profile"
	TeamsacsAccounting = "accounting"
	TeamsacsAuthlog = "authlog"

	GenieacsDevices = "devices"
	GenieacsFaults  = "faults"
	GenieacsTasks   = "tasks"
	GenieacsPresets = "presets"
)
type Doc = map[string]interface{}

type ModelManager struct {
	Config       *config.AppConfig
	Mongo        *mongo.Client
	Sched        *gocron.Scheduler
	LongCache    *bigcache.BigCache
	M1Cache      *bigcache.BigCache
	M5Cache      *bigcache.BigCache
	TplRender    *tpl.CommonTemplate
	Location     *time.Location
	WebJwtConfig *middleware.JWTConfig
	MailSender   *gmail.MailSender
	ManagerMap   cmap.ConcurrentMap
	Dev          bool
}

func NewModelManager(appconfig *config.AppConfig, dev bool) *ModelManager {
	m := &ModelManager{Config: appconfig, Dev: dev}
	m.ManagerMap = cmap.New()
	_mongodb, err := mongodb.GetMongodbClient(appconfig.Mongodb)
	common.Must(err)
	m.Mongo = _mongodb
	loc, err := time.LoadLocation(appconfig.System.Location)
	common.Must(err)
	m.Location = loc
	m.registerManagers()
	m.TplRender = tpl.NewCommonTemplate([]string{"/resources/templates"}, m.Dev, m.GetTemplateFuncMap())
	go m.StartScheduler()
	return m
}

func (m *ModelManager) registerManagers()  {
	m.ManagerMap.Set("SubscribeManager", &SubscribeManager{m})
}


func (m *ModelManager) GetTeamsAcsCollection(coll string) *mongo.Collection {
	return m.Mongo.Database(MDBTeamsacs).Collection(coll)
}

func (m *ModelManager) GetGenieAcsCollection(coll string) *mongo.Collection {
	return m.Mongo.Database(MDBTeamsacs).Collection(coll)
}

func (m *ModelManager) GetTemplateFuncMap() map[string]interface{} {
	return map[string]interface{}{
		"Pagever": func() int64 {
			return time.Now().Unix()
		},
	}
}

