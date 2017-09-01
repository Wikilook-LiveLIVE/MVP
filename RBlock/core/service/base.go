package service

import (
	"github.com/artjoma/flog"
	"RBlock/utils"
	"RBlock/templateengine"
	"RBlock/dao"
)

//global variables for this package
var (
	log          *flog.Logger //main logger
	siteLocation string

	templateService        *templateengine.TemplateEngine
	encryptionService      *utils.EncryptionService
	mailService       		*MailService
	customerService      	*CustomerService
	customerDAO     		*dao.CustomerDAO

)

