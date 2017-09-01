package dao

import (
	"encoding/json"
)

type CustomerEntity struct {

}

func (self *CustomerEntity) FromJSON(jsonData []byte) error {
	return json.Unmarshal(jsonData, self)
}
