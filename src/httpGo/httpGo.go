package httpGo

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
)

func StartServer(attr string) {
	http.ListenAndServe(attr, nil)
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

func GetUploadedFile(request *http.Request, name string) (multipart.File, int64) {
	request.ParseMultipartForm(1024 * 1024 * 10)
	fileheader := request.MultipartForm.File[name][0]
	length := fileheader.Size
	file, err := fileheader.Open()
	if err == nil {
		log.Println("Uploaded file name: " + fileheader.Filename + ", Size: " + strconv.FormatInt(length, 10))
		return file, length
	} else {
		log.Fatal("Error Opening File")
		return nil, 0
	}
}

func ResponseJSON(response http.ResponseWriter, jsonObject interface{}) {
	response.Header().Set("Content-Type", "application/json")
	jsonObj, _ := json.Marshal(jsonObject)
	response.Write(jsonObj)
}
