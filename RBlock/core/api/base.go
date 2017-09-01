package api

import (
	"github.com/artjoma/flog"
	"RBlock/utils"
	"RBlock/core/service"
)

var (
	log               *flog.Logger //main logger
	encryptionService *utils.EncryptionService
	customerService   *service.CustomerService
)


