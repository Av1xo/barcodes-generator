package main

import (
	"fmt"
	"image/png"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/ean"
	"github.com/boombuler/barcode/pdf417"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode/twooffive"
)

func twooffiveGenerator(w http.ResponseWriter, dataString string) {
	checkSumData, err := twooffive.AddCheckSum(dataString)
	if err != nil {
		fmt.Println("CheckSum Error:", err)
		http.Error(w, "Opps twooffive encode only digits"+err.Error(), http.StatusInternalServerError)
		return
	}

	twooffiveCode, err := twooffive.Encode(checkSumData, true)
	if err != nil {
		fmt.Println("Encode Error:", err)
		http.Error(w, "Opps twooffive encode only digits"+err.Error(), http.StatusInternalServerError)
		return
	}

	twooffiveCode, err = barcode.Scale(twooffiveCode, 512, 256)
	if err != nil {
		fmt.Println("Scale Error:", err)
		http.Error(w, "Opps twooffive encode only digits"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := png.Encode(w, twooffiveCode); err != nil {
		fmt.Println("PNG Error:", err)
		http.Error(w, "cannot encode to png: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func pdf417Generator(w http.ResponseWriter, dataString string) {
	pdf417Code, err := pdf417.Encode(dataString, 8)
	if err != nil {
		fmt.Println("Encode Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	pdf417Code, err = barcode.Scale(pdf417Code, 1024, 256)
	if err != nil {
		fmt.Println("Scale Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := png.Encode(w, pdf417Code); err != nil {
		http.Error(w, "cannot encode to png", http.StatusInternalServerError)
		return
	}
}

func datamatrixGenerator(w http.ResponseWriter, dataString string) {
	datamatrixCode, err := datamatrix.Encode(dataString)
	if err != nil {
		fmt.Println("Encode Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	datamatrixCode, err = barcode.Scale(datamatrixCode, 512, 512)
	if err != nil {
		fmt.Println("Scale Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := png.Encode(w, datamatrixCode); err != nil {
		http.Error(w, "cannot encode to png", http.StatusInternalServerError)
		return
	}
}

func eanGenerator(w http.ResponseWriter, dataString string) {
	if len(dataString) != 7 && len(dataString) != 12 {
		fmt.Println("Invalid data")
		http.Error(w, "Opps ean must be a 7 or 12 numbers", http.StatusInternalServerError)
		return
	}

	eanCode, err := ean.Encode(dataString)
	if err != nil {
		fmt.Println("Encode Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	scaledCode, _ := barcode.Scale(eanCode, 512, 512)

	if err := png.Encode(w, scaledCode); err != nil {
		fmt.Println("PNG Error:", err)
		http.Error(w, "cannot encode to png: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func code32Generator(w http.ResponseWriter, dataString string) {
	panic("unimplemented")
}

func code93Generator(w http.ResponseWriter, dataString string) {
	panic("unimplemented")
}

func code128Generator(w http.ResponseWriter, dataString string) {
	panic("unimplemented")
}

func aztecGenerator(w http.ResponseWriter, dataString string) {
	panic("unimplemented")
}

func codabarGenerator(w http.ResponseWriter, dataString string) {
	panic("unimplemented")
}

func qrGenerator(w http.ResponseWriter, data string) {
	qrCode, err := qr.Encode(data, qr.L, qr.Auto)
	if err != nil {
		http.Error(w, "cannot create qr code", http.StatusInternalServerError)
		return
	}

	qrCode, err = barcode.Scale(qrCode, 512, 512)
	if err != nil {
		http.Error(w, "cannot scale qr code", http.StatusInternalServerError)
		return
	}

	if err := png.Encode(w, qrCode); err != nil {
		http.Error(w, "cannot encode to png", http.StatusInternalServerError)
		return
	}
}
