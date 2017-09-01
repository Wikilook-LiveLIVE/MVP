package test

import (
	"testing"
	"RBlock/utils"
	"os"
	"fmt"
)

func TestGetNewAESKey_16Bytes(t *testing.T) {
	encryptionService := utils.NewEncryptionService("", "")
	key, err := encryptionService.GenerateAESKey(16)
	if err !=nil {
		fmt.Println("Error: ",err.Error())
		os.Exit(1)
	}
	fmt.Println("128bitAES Key:" ,key)
}

func TestGetNewAESKey_24Bytes(t *testing.T) {
	encryptionService := utils.NewEncryptionService("", "")
	key, err := encryptionService.GenerateAESKey(24)
	if err !=nil {
		fmt.Println("Error: ",err.Error())
		os.Exit(1)
	}
	fmt.Println("192bitAES Key:" ,key)
}

func TestGetNewAESKey_32Bytes(t *testing.T) {
	encryptionService := utils.NewEncryptionService("", "")
	key, err := encryptionService.GenerateAESKey(32)
	if err !=nil {
		fmt.Println("Error: ",err.Error())
		os.Exit(1)
	}
	fmt.Println("256bitAES Key:" ,key)
}