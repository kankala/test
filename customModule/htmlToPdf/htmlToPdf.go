package htmlToPdf

import (
	"fmt"
	"log"
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func ToPdf() {

	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	htmlStr := `<html><body><h1 style="color:red;">This is an html
 		from pdf to test color<h1><img src="http://api.qrserver.com/v1/create-qr-
		code/?data=HelloWorld" alt="img" height="420" width="420"></img></body></html>`

	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	//Your Pdf Name
	err = pdfg.WriteFile("./123.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
