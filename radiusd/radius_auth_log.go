package radiusd

import (
	"time"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/pool.v3"

	"github.com/ca17/teamsacs/constant"
)

func (s *RadiusService) addAuthlog(start time.Time, username string, nasip string, result string, reason string) {
	gpool.Queue(func() pool.WorkFunc {
		return func(pool.WorkUnit) (interface{}, error) {
			err := s.Manager.GetRadiusManager().AddRadiusAuthLog(username, nasip, result, reason, time.Since(start).Milliseconds())
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
	}())
}

func (s *RadiusService) CheckRadAuthError(start time.Time,username, nasip string, err error) {
	if err != nil {
		logLevel := s.GetStringConfig(constant.RadiusAuthlogLevel, RadiusAuthlogAll)
		if logLevel != RadiusAuthlogNone && (logLevel == RadiusAuthlogAll || logLevel == RadiusAuthFailure) {
			s.addAuthlog(start, username, nasip, RadiusAuthFailure, err.Error())
		}
		if IsDebug() {
			panic(errors.WithStack(err))
		} else {
			panic(err)
		}
	}
}

func (s *RadiusService) LogAuthSucess(start time.Time,username, nasip string) {
	logLevel := s.GetStringConfig(constant.RadiusAuthlogLevel, RadiusAuthlogAll)
	if logLevel != RadiusAuthlogNone && (logLevel == RadiusAuthlogAll || logLevel == RadiusAuthSucces) {
		s.addAuthlog(start, username, nasip, RadiusAuthSucces, RadiusAuthSucces)
	}
}

