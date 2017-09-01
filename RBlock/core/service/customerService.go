package service

import (
	"fmt"
	"RBlock/core/types"
	"RBlock/errors"
	"RBlock/utils"
)

type CustomerService struct {
}

func NewCustomerSerivce() *CustomerService {
	return &CustomerService{}
}

/*
SignUp
 */
func (self *CustomerService) SignUp(signUp *types.SignUp) *errors.AppErr {
	log.Info(fmt.Sprint(signUp.Email))


	mailEnc, err := encryptionService.EncryptDataToHex([]byte(signUp.Email))
	if err != nil {
		log.Err("encryptionService.EncryptDataToHex " + err.Error())
	}
	passHash := utils.ToHash([]byte(signUp.Password))
	activationCode := utils.RandString(16)

	if _, err := customerDAO.SaveSignUp(mailEnc, passHash, activationCode); err != nil {
		log.Err("Err customerDAO.SaveSignUp " + err.Error())
		return errors.ERR_METHOD_UNAVAILABLE
	}
	paramMap := map[string]string{"siteLocation": siteLocation, "activationCode": activationCode}
	buffer, err := templateService.Execute("activation_mail", paramMap)

	if err != nil {
		log.Err("Err templateService.Execute. " + err.Error())
		return errors.ERR_METHOD_UNAVAILABLE
	}

	if err := mailService.SendMail(signUp.Email, "Your RBlock Account Activation!", buffer.Bytes()); err != nil {
		log.Err("Err mailService.SendMail. " + err.Error())
		return errors.ERR_METHOD_UNAVAILABLE
	}
	return nil
}
/*
SignUpActivation
 */
func (self *CustomerService) SignUpActivation(acivationCode string) *errors.AppErr {
	log.Info("Acivation Code: " + acivationCode)
	customerId, err := customerDAO.ActivateByRegKey(acivationCode)
	if err != nil {
		log.Err("customerDAO.ActivateByRegKey " + err.Error())
		return errors.ERR_METHOD_UNAVAILABLE
	}
	if customerId == 0 {
		return errors.ERR_CUSTOMER_NOT_FOUND
	}
	return nil
}
