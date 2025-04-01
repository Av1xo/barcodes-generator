package main

import (
	"fmt"
	"image/png"
	"net/http"

	"github.com/boombuler/barcode"
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
	
}

func datamatrixGenerator(w http.ResponseWriter, dataString string) {
	panic("unimplemented")
}

func eanGenerator(w http.ResponseWriter, dataString string) {
	panic("unimplemented")
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
