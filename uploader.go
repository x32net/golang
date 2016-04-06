// $ go get -u all
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// to get file's size
type Size interface {
	Size() int64
}

//var uploadTemplate = template.Must(template.ParseFiles("page.html"))

func indexHandle(w http.ResponseWriter, r *http.Request) {
	//	if err := uploadTemplate.Execute(w, nil); err != nil {
	//		log.Fatal("Execute: ", err.Error())
	//		return
	//	}
	w.Write([]byte(fmt.Sprintf(`<html> <head> <title>Golang upload</title> </head> <body> <form id="uploadForm" method="POST" enctype="multipart/form-data" action="/upload"> <p>Golang upload</p> <input type="file" id="file" name="file" /> <input type="submit" value="upload"> </form> </body> </html>`)))
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5000)
	file, header, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(header.Filename)
	//	if sizeInterface, ok := file.(Size); ok {
	//		fmt.Fprintf(w, "size is: %d", sizeInterface.Size())
	//		return
	//	} else {
	//		fmt.Println("can't get")
	//	}

	t, err := os.Create("./" + header.Filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer t.Close()

	io.Copy(t, file)
	http.Redirect(w, r, "/", 303)
	return
}

func main() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/upload", uploadHandle)
	http.ListenAndServe(":8080", nil)
}
