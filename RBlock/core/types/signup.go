package types

import (
	"github.com/buger/jsonparser"
	"errors"
)

type SignUp struct {
	Email       string
	Password    string
}

func (self *SignUp) FromJSON(json []byte) error {
	var err error = nil
	var confirmedPassword string

	if self.Email, err = jsonparser.GetString(json, "email"); err != nil {
		return err
	}
	if self.Password, err = jsonparser.GetString(json, "password"); err != nil {
		return err
	}

	if confirmedPassword, err = jsonparser.GetString(json, "confirmedPassword"); err != nil {
		return err
	}

	if confirmedPassword != self.Password {
		return errors.New("Password not equals confirmed password!")
	}

	return err
}
