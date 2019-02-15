package main

import (
	"net/http"
	"learngo/errhandling/filelinstingserver/filelisting"
	"os"
	"log"
)
type appHandle func (writer http.ResponseWriter ,request *http.Request) error
func errWrappper(handler appHandle) func(http.ResponseWriter,*http.Request){
	return func (writer http.ResponseWriter,request *http.Request){
		defer func(){
			if r := recover() ;r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer,request)
		if err !=nil{
			log.Printf("Error occurred handling request: %s",err.Error())
			//if userErr,ok := err.(userError); ok{
			//	http.Error(writer,userErr.Message(),http.StatusBadRequest)
			//	return
			//}
			code := http.StatusOK
			switch  {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/",errWrappper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888",nil)
	if err!= nil{
		panic(err)
	}
}
