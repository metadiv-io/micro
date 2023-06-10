package micro

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginmid"
	"github.com/metadiv-io/micro/auth"
	"github.com/metadiv-io/micro/usage"
	"github.com/robfig/cron"
)

type Engine struct {
	GinEngine  *gin.Engine
	CronWorker *cron.Cron
}

func NewEngine() *Engine {
	e := &Engine{
		GinEngine:  gin.Default(),
		CronWorker: cron.New(),
	}
	return e
}

func (e *Engine) SetupCORS(config *cors.Config) {
	e.GinEngine.Use(cors.New(*config))
}

func (e *Engine) SetupGeneralRateLimit(duration time.Duration, rate int64) {
	e.GinEngine.Use(ginmid.RateLimited(duration, rate))
}

func (e *Engine) SetupAuthCron() {
	CRON(e, "@every 1m", auth.RegisterCron)
}

func (e *Engine) SetupUsageCron() {
	CRON(e, "@every 1m", usage.SendConsumptionCron)
}

func (e *Engine) Run(addr string) {
	GET(e, "/ping", PingHandler)
	e.CronWorker.Start()
	e.GinEngine.Run(addr)
}
