package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Av1xo/barcode/utils"
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

	scaleWidth, err := strconv.Atoi(r.FormValue("width"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	scaleHeight, err := strconv.Atoi(r.FormValue("height"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	size := utils.Size{
		Width: scaleWidth,
		Height: scaleHeight,
	}

	switch barcodeType {
	case "qr":
		barcode, err := utils.QrGenerator(dataString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		utils.PngEncode(barcode, size, w)

	case "codabar":
		barcode, err := utils.CodabarGenerator(dataString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		utils.PngEncode(barcode, size, w)
	case "aztec":
		barcode, err := utils.AztecGenerator(dataString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		utils.PngEncode(barcode, size, w)
	case "code128":
		barcode, err := utils.Code128Generator(dataString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		utils.PngEncode(barcode, size, w)
	case "code93":
		barcode, err := utils.Code93Generator(dataString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		utils.PngEncode(barcode, size, w)
	case "code39":
		barcode, err := utils.Code39Generator(dataString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		utils.PngEncode(barcode, size, w)
	case "ean":
		barcode, err := utils.EanGenerator(dataString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		utils.PngEncode(barcode, size, w)
	case "datamatrix":
		barcode, err := utils.DatamatrixGenerator(dataString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		utils.PngEncode(barcode, size, w)
	case "pdf417":
		barcode, err := utils.Pdf417Generator(dataString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		utils.PngEncode(barcode, size, w)
	case "twooffive":
		barcode, err := utils.TwooffiveGenerator(dataString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		utils.PngEncode(barcode, size, w)
	default:
	}
}
