package main

//break        default      func         interface    select
//case         defer        go           map          struct
//chan         else         goto         package      switch
//const        fallthrough  if           range        type
//continue     for          import       return       var
//+    &     +=    &=     &&    ==    !=    (    )
//-    |     -=    |=     ||    <     <=    [    ]
//*    ^     *=    ^=     <-    >     >=    {    }
///    <<    /=    <<=    ++    =     :=    ,    ;
//%    >>    %=    >>=    --    !     ...   .    :
//     &^          &^=

import (
	"YAKALENDARPEREVERNY/pkg/logging"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	var resultConf ResultConf
	var err error
	resultConf, err = LoadConfig()
	if err != nil {
		fmt.Printf("Это моя ошибка: %v", err)
	}
	fmt.Printf("Это моя бд: %s", resultConf.Baseconnector.Dbname)

}

var (
	configFullPath string
	configOpenFile *os.File
	readConfigFile []byte
	err            error
)

type ResultConf struct {
	Baseconnector struct {
		Dbname     string `json:"dbname"`
		Dbuser     string `json:"dbuser"`
		Dbpassword string `json:"dbpassword"`
	} `json:"baseconnector"`
	Logcon struct {
		Loglvl string `json:"loglvl"`
	} `json:"logcon"`
}

func LoadConfig() (result ResultConf, err error) {
	mainpass := "мой главный pass: %s"
	configFullPath, err = os.Getwd()
	logging.GlobalLog(logging.Logger{
		LoggerLevel:            "ОТВЕЛ ТВОЕГО ОТЦА ЗА ХЛЕБОМ",
		LoggerMessage:          fmt.Sprintf(mainpass, configFullPath),
		LoggerError:            nil,
		LoggerSubLoggerMessage: logging.SubLogger{},
	})
	ValidateError(err)
	mainpass = "мой главный file: %s"
	configOpenFile, err = os.Open(configFullPath + "\\config.json")
	ValidateError(err)
	readConfigFile, err = io.ReadAll(configOpenFile)
	ValidateError(err)
	validFile := readConfigFile
	err = json.Unmarshal(validFile, &result)
	err = ValidateError(err)
	return
}

func ValidateError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
