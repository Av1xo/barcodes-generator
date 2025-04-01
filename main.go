package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Page struct {
	Title string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/generator", viewCodeHandler)
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "Barcode Generator"}
	t, _ := template.ParseFiles("generator.html")
	t.Execute(w, p)
}

func viewCodeHandler(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")
	barcodeType := r.FormValue("barcodeType")

	switch barcodeType {
	case "qr":
		qrGenerator(w, dataString)
	case "codabar":
		codabarGenerator(w, dataString)
	case "aztec":
		aztecGenerator(w, dataString)
	case "code128":
		code128Generator(w, dataString)
	case "code93":
		code93Generator(w, dataString)
	case "code32":
		code32Generator(w, dataString)
	case "ean":
		eanGenerator(w, dataString)
	case "datamatrix":
		datamatrixGenerator(w, dataString)
	case "pdf417":
		pdf417Generator(w, dataString)
	case "twooffive":
		twooffiveGenerator(w, dataString)
	default:
	}
}
