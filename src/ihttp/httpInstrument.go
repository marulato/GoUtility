package ihttp

import (
	"log"
	"net/http"
	"os"
)

func StartServer(attr string) {
	err := http.ListenAndServe(attr, nil)
	if err == nil {
		log.Println("Go Server Started up at " + attr)
	} else {
		log.Fatal(err)
	}
}

func GetFormValue(request *http.Request, key string) string {
	return request.PostFormValue(key)
}

func GetUrlValue(request *http.Request, key string) string {
	return request.FormValue(key)
}

func GetData(request *http.Request) []byte {
	length := request.ContentLength
	data := make([]byte, length)
	request.Body.Read(data)
	return data
}

func GetUploadedFile(request *http.Request, name string) os.File {
	request.ParseMultipartForm(1024 * 1024 * 10)
	fileheader := request.MultipartForm.File[name][0]
	file, err := fileheader.Open()
	if err == nil {
		return file
	} else {
		log.Fatal("Error Opening File")
		return nil
	}
}
