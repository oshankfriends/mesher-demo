package common

import (
	"os"
	"log"
	"encoding/json"
	"net/http"
)

type (
	configuration struct {
		Server      string     `json:"server"`
	}
	errorResource struct {
		Data        appError    `json:"data"`
	}
	appError struct {
		Error       string	`json:"error"`
		Message     string	`json:"message"`
		HttpStatus  int         `json:"status"`
	}
)

var AppConfig configuration

func DisplayError(w http.ResponseWriter, handlerError error,message string, code int){
	errObj := appError{
		Error:   handlerError.Error(),
		Message:  message,
		HttpStatus: code,
	}
	log.Printf("[AppError]: %s",handlerError)
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(code)
	if js,err := json.Marshal(errorResource{Data:errObj}); err == nil {
		w.Write(js)
	}
}

func init(){
	file,err := os.Open("conf/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n",err)
	}
	AppConfig = configuration{}
	if err := json.NewDecoder(file).Decode(&AppConfig); err != nil {
		log.Fatalf("[loadAppConfig]: %s",err)
	}

}