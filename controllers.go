package controllers

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func UplaodFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return 
	}

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return 
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
		case "image/png", "image/jpg":
			saveFile(w, file, handle)
		default:
			jsonResponse(w, http.StatusBadRequest, "The format file is not valid.")
	}
}

func saveFile(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(file)
	if err !=  nil {
		fmtFprintf(w, "%v", err)
		return 
	}

	err = ioutil.WriteFile("./file/"+handle.Filename, data, 0666)
	if err !=  nil {
		fmt.Fprintf(w, "%v", err)
		return 
	}
	jsonResponse(w, http.StatusCreated, "File upload successfully!")
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, message)
}


