package structs

import (
	"github.com/artjoma/flog"
	"RBlock/core/api"
)

type AppContext struct {
	AppFolder         string
	LogManager        *flog.LogManager
	Logger            *flog.Logger //main logger
	ConfigFile        *ConfigFile
	HttpApiController *api.HttpApiController
}