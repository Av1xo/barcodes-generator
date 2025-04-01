package utils

import (
	"fmt"
	"image/png"
	"net/http"

	"github.com/boombuler/barcode"
)

type Size struct {
	Width  int
	Height int
}

func PngEncode(bcode barcode.Barcode, size Size, w http.ResponseWriter) {
	scaledCode, err := barcode.Scale(bcode, size.Width, size.Height)
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
