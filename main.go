package main

import (
	"flag"
	"fmt"
	"gencert/cert"
	"gencert/html"
	"gencert/pdf"
	"os"
)

func main() {
	csvFile := flag.String("file", "", "CSV file to parse.")
	outputType := flag.String("type", "pdf", "Output type of the certificate.")
	flag.Parse()

	if len(*csvFile) <= 0 {
		fmt.Printf("Invalid file (len=%v)", len(*csvFile))
		os.Exit(1)
	}

	certs, err := cert.ParseCSV(*csvFile)
	if nil != err {
		fmt.Printf("Failed to parse CSV file '%v'. err=%v\n", *csvFile, err)
		os.Exit(1)
	}

	var saver cert.Saver
	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		fmt.Printf("Unknown type: %v", *outputType)
	}
	if nil != err {
		fmt.Println("Error while saving certificate")
		os.Exit(1)
	}

	for _, c := range certs {
		saver.Save(*c)
	}

}
