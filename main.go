package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

func generateDoc(w io.Writer) (err error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("dejavu", "", "../gofpdf/font/DejaVuSansCondensed.ttf")
	pdf.SetFont("dejavu", "", 12)
	pdf.Cell(40, 10, "Привет, мир!")
	err = pdf.Output(w)
	pdf.Close()
	return
}

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	reportHandler := func(w http.ResponseWriter, req *http.Request) {

		var err error
		var buf bytes.Buffer

		err = generateDoc(&buf)
		if err == nil {
			w.Header().Set("Content-Type", "application/pdf")
			buf.WriteTo(w)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return

		/*
		 */

		/*
			fileD, err := os.Open("hello.pdf")
			if err != nil {
				log.Panic(err)
			}
			file_bytes, err := io.ReadAll(fileD)
			if err != nil {
				log.Panic(err)
			}

			w.Header().Set("Content-type", "application/octet-stream")
			w.Write(file_bytes)
		*/
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/report", reportHandler)
	log.Println("Listing for requests at http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
