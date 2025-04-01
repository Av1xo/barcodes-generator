package main

import (
	"fmt"
	"image/png"
	"net/http"
	"regexp"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/aztec"
	"github.com/boombuler/barcode/codabar"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/code39"
	"github.com/boombuler/barcode/code93"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/ean"
	"github.com/boombuler/barcode/pdf417"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode/twooffive"
)

func twooffiveGenerator(w http.ResponseWriter, dataString string) {
	if !validateNumbers(dataString) {
		http.Error(w, "Opps input numbers must ne a numbers", http.StatusBadRequest)
		return
	}
	checkSumData, err := twooffive.AddCheckSum(dataString)
	if err != nil {
		fmt.Println("CheckSum Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	twooffiveCode, err := twooffive.Encode(checkSumData, true)
	if err != nil {
		fmt.Println("Encode Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	twooffiveCode, err = barcode.Scale(twooffiveCode, 512, 256)
	if err != nil {
		fmt.Println("Scale Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
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
	if len(dataString) != 7 && len(dataString) != 12 && !validateNumbers(dataString) {
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

	scaledCode, err := barcode.Scale(eanCode, 512, 512)
	if err != nil {
		fmt.Println("Scale Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := png.Encode(w, scaledCode); err != nil {
		fmt.Println("PNG Error:", err)
		http.Error(w, "cannot encode to png: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func code39Generator(w http.ResponseWriter, dataString string) {
	if !validateCode39(dataString) {
		http.Error(w, "Opps invalid input", http.StatusBadRequest)
		return
	}

	code39Code, err := code39.Encode(dataString, true, true)
	if err != nil {
		fmt.Println("Encode Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	scaledCode, err := barcode.Scale(code39Code, 512, 512)
	if err != nil {
		fmt.Println("Scale Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := png.Encode(w, scaledCode); err != nil {
		http.Error(w, "cannot encode to png", http.StatusInternalServerError)
		return
	}
}

func code93Generator(w http.ResponseWriter, dataString string) {
	if !validateCode93(dataString) {
		http.Error(w, "Opps Opps invalid input", http.StatusBadRequest)
		return
	}

	code93Code, err := code93.Encode(dataString, true, true)
	if err != nil {
		fmt.Println("Encode Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	scaledCode, err := barcode.Scale(code93Code, 512, 512)
	if err != nil {
		fmt.Println("Scale Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := png.Encode(w, scaledCode); err != nil {
		http.Error(w, "cannot encode to png", http.StatusInternalServerError)
		return
	}
}

func code128Generator(w http.ResponseWriter, dataString string) {
	if !validateASCII(dataString) {
		http.Error(w, "Opps Opps invalid input", http.StatusBadRequest)
		return
	}

	code128Code, err := code128.Encode(dataString)
	if err != nil {
		fmt.Println("Encode Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	scaledCode, err := barcode.Scale(code128Code, 512, 512)
	if err != nil {
		fmt.Println("Scale Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := png.Encode(w, scaledCode); err != nil {
		http.Error(w, "cannot encode to png", http.StatusInternalServerError)
		return
	}
}

func aztecGenerator(w http.ResponseWriter, dataString string) {
	if !validateAztec(dataString) {
		http.Error(w, "Opps Opps invalid input", http.StatusBadRequest)
		return
	}

	aztecCode, err := aztec.Encode([]byte(dataString), aztec.DEFAULT_EC_PERCENT, aztec.DEFAULT_LAYERS)
	if err != nil {
		fmt.Println("Encode Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	scaledCode, err := barcode.Scale(aztecCode, 512, 512)
	if err != nil {
		fmt.Println("Scale Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := png.Encode(w, scaledCode); err != nil {
		http.Error(w, "cannot encode to png", http.StatusInternalServerError)
		return
	}
}

func codabarGenerator(w http.ResponseWriter, dataString string) {
	if !validateCodabar(dataString) {
		http.Error(w, "Opps Opps invalid input", http.StatusBadRequest)
		return
	}

	codabarCode, err := codabar.Encode(dataString)
	if err != nil {
		fmt.Println("Encode Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	scaledCode, err := barcode.Scale(codabarCode, 512, 512)
	if err != nil {
		fmt.Println("Scale Error:", err)
		http.Error(w, "Opps something went wrong"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := png.Encode(w, scaledCode); err != nil {
		http.Error(w, "cannot encode to png", http.StatusInternalServerError)
		return
	}
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

func validateNumbers(input string) bool {
	regex := "^[0-9]+$"
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}

func validateCode39(input string) bool {
	regex := "^[A-Z0-9\\-\\.\\$\\/\\+\\%]+$"
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}

func validateCode93(input string) bool {
	regex := "^[A-Z0-9!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~]+$"
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}

func validateASCII(input string) bool {
	regex := "^[\\x00-\\x7F]+$"
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}

func validateAztec(input string) bool {
	// (A-Z, a-z), (0-9)
	// space, +, -, ., $, /, :, ;, etc.
	regex := "^[A-Za-z0-9\\-\\.\\$\\/\\:\\+\\,\\?\\!\\*\\(\\)]+$"
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}

func validateCodabar(input string) bool {
	// Codabar: start and end [A, B, C, D], in the mid nums and sym
	regex := "^[ABCD][0-9\\-\\.\\$\\/\\:\\+]+[ABCD]$"
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}
