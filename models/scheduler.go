package models

import "github.com/go-co-op/gocron"

func (m *ModelManager) StartScheduler()  {
	m.Sched = gocron.NewScheduler(m.Location)
	<-m.Sched.Start()
}

