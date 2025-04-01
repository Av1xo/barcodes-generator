package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
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

	barCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	barCode, _ = barcode.Scale(barCode, 512, 512)
	png.Encode(w, barCode)
}
