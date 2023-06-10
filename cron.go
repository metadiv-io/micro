package micro

func CRON(engine *Engine, spec string, job func()) {
	engine.CronWorker.AddFunc(spec, job)
}
