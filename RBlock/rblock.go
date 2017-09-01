package main

import (
	"fmt"
	"os"
	"RBlock/core/structs"
	"path/filepath"
	"github.com/artjoma/flog"
	"encoding/json"
	"os/signal"
	"syscall"
	"time"
	"RBlock/core/api"
	"RBlock/core/service"
	"RBlock/utils"
	"RBlock/dao"
)

const (
	APP_CONFIG_FILE_NAME = "config.json"
	APP_MAIN_LOGGER      = "MAIN_LOGGER"
)

type AppContext structs.AppContext

func main() {


}

func NewContext(processArgs []string) *AppContext {
	fmt.Println("Start create app context")

	workPath, err := os.Getwd()
	if err != nil {
		panic("Can't get working path")
	}
	fmt.Println("Working path: ", workPath)

	//reading config file
	configFile, err := os.Open(filepath.Join(workPath, APP_CONFIG_FILE_NAME))
	if err != nil {
		panic("Application config file doen't exist !")
	}
	defer configFile.Close()
	config := &structs.ConfigFile{}
	err = json.NewDecoder(configFile).Decode(config)
	if err != nil {
		panic("Invalid config file! " + err.Error())
	}
	//init logger manager
	fmt.Println("Logger console: ", config.LoggerConsole)
	var logManager *flog.LogManager = nil
	if config.LoggerConsole {
		logManager = flog.NewLogManagerConsole()
	} else {
		logManager = flog.NewLogManagerFile(workPath, config.LogFileSize)
	}
	fmt.Printf("LogThreshold: %q\n", config.LogThreshold)
	//create main logger
	logger := logManager.NewLogger(APP_MAIN_LOGGER, config.LogThreshold)


	return &AppContext{AppFolder: workPath, LogManager: logManager, Logger: logger, ConfigFile: config}
}



func (self *AppContext) createServices() {

}







//Free resources
func (self *AppContext) destroy() {


}