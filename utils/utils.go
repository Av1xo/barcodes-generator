package utils

import (
	"errors"
	"fmt"
	"log"

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

func TwooffiveGenerator(dataString string) (barcode.Barcode, error) {
	if !validateNumbers(dataString) {

		return nil, errors.New("opps input numbers must ne a numbers")
	}
	checkSumData, err := twooffive.AddCheckSum(dataString)
	if err != nil {
		fmt.Println("CheckSum Error:", err)
		return nil, errors.New("opps something went wrong")
	}

	twooffiveCode, err := twooffive.Encode(checkSumData, true)
	if err != nil {
		fmt.Println("Encode Error:", err)
		return nil, errors.New("opps something went wrong")
	}

	return twooffiveCode, nil
}

func Pdf417Generator(dataString string) (barcode.Barcode, error) {
	pdf417Code, err := pdf417.Encode(dataString, 8)
	if err != nil {
		fmt.Println("Encode Error:", err)
		return nil, errors.New("opps something went wrong")
	}

	return pdf417Code, nil
}

func DatamatrixGenerator(dataString string) (barcode.Barcode, error) {
	datamatrixCode, err := datamatrix.Encode(dataString)
	if err != nil {
		fmt.Println("Encode Error:", err)
		return nil, errors.New("opps something went wrong")
	}

	return datamatrixCode, nil
}

func EanGenerator(dataString string) (barcode.Barcode, error) {
	if len(dataString) != 7 && len(dataString) != 12 && !validateNumbers(dataString) {
		fmt.Println("Invalid data")
		return nil, errors.New("opps ean must be a 7 or 12 numbers")
	}

	eanCode, err := ean.Encode(dataString)
	if err != nil {
		fmt.Println("Encode Error:", err)
		return nil, errors.New("opps something went wrong")
	}

	return eanCode, nil
}

func Code39Generator(dataString string) (barcode.Barcode, error) {
	if !validateCode39(dataString) {
		return nil, errors.New("opps invalid input")
	}

	code39Code, err := code39.Encode(dataString, true, true)
	if err != nil {
		fmt.Println("Encode Error:", err)
		return nil, errors.New("opps invalid input")
	}

	return code39Code, nil
}

func Code93Generator(dataString string) (barcode.Barcode, error) {
	if !validateCode93(dataString) {
		return nil, errors.New("opps invalid input")
	}

	code93Code, err := code93.Encode(dataString, true, true)
	if err != nil {
		fmt.Println("Encode Error:", err)
		return nil, errors.New("opps something went wrong")
	}

	return code93Code, nil
}

func Code128Generator(dataString string) (barcode.Barcode, error) {
	if !validateASCII(dataString) {
		return nil, errors.New("opps invalid input")
	}

	code128Code, err := code128.Encode(dataString)
	if err != nil {
		fmt.Println("Encode Error:", err)
		return nil, errors.New("opps something went wrong")
	}

	return code128Code, nil
}

func AztecGenerator(dataString string) (barcode.Barcode, error) {
	if !validateAztec(dataString) {
		return nil, errors.New("opps invalid input")
	}

	aztecCode, err := aztec.Encode([]byte(dataString), aztec.DEFAULT_EC_PERCENT, aztec.DEFAULT_LAYERS)
	if err != nil {
		fmt.Println("Encode Error:", err)
		return nil, errors.New("opps something went wrong")
	}

	return aztecCode, nil
}

func CodabarGenerator(dataString string) (barcode.Barcode, error) {
	if !validateCodabar(dataString) {
		return nil, errors.New("opps invalid input")
	}

	codabarCode, err := codabar.Encode(dataString)
	if err != nil {
		fmt.Println("Encode Error:", err)
		return nil, errors.New("opps something went wrong")
	}

	return codabarCode, nil
}

func QrGenerator(dataString string) (barcode.Barcode, error) {
	qrCode, err := qr.Encode(dataString, qr.L, qr.Auto)
	if err != nil {
		log.Println("Error Encode: " + err.Error())
		return nil, errors.New("opps something went wrong")
	}

	return qrCode, nil
}
