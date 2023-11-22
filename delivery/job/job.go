package job

import "github.com/teq-quocbang/arrows/config"

type IJob interface {
	Run()
}

type Jobs []IJob

func (js Jobs) Run() {
	for _, j := range js {
		go j.Run()
	}
}

func New() Jobs {
	return Jobs{
		NewInsufficientCreditsCheck(config.GetConfig().CheckInsufficientCreditsEndPoint),
		NewHealthChecks(config.GetConfig().HealthCheck.HealthCheckEndPoint),
	}
}
