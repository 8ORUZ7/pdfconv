package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	unipdf "github.com/unidoc/unipdf/v3/model"
	"github.com/baliance/gooxml/document"
)

func extractTextFromPDF(pdfPath string) (string, error) {
	file, err := os.Open(pdfPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	pdfReader, err := unipdf.NewPdfReader(file)
	if err != nil {
		return "", err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return "", err
	}
	if isEncrypted {
		if ok, _ := pdfReader.Decrypt([]byte("")); !ok {
			return "", fmt.Errorf("could not decrypt PDF")
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return "", err
	}

	var textBuilder strings.Builder
	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			return "", err
		}
		ex, err := page.ExtractText()
		if err != nil {
			return "", err
		}
		textBuilder.WriteString(ex)
		textBuilder.WriteString("\n")
	}

	return textBuilder.String(), nil
}

func saveAsTxt(text, filename string) error {
	path := filepath.Join("output", filename)
	return ioutil.WriteFile(path, []byte(text), 0644)
}

func saveAsDocx(text, filename string) error {
	doc := document.New()
	for _, line := range strings.Split(text, "\n") {
		doc.AddParagraph().AddRun().AddText(line)
	}
	return doc.SaveToFile(filepath.Join("output", filename))
}

func convertPDF(pdfPath, outputFormat string) error {
	text, err := extractTextFromPDF(pdfPath)
	if err != nil {
		return err
	}

	os.MkdirAll("output", os.ModePerm)

	base := strings.TrimSuffix(filepath.Base(pdfPath), filepath.Ext(pdfPath))
	switch strings.ToLower(outputFormat) {
	case "txt":
		return saveAsTxt(text, base+".txt")
	case "docx":
		return saveAsDocx(text, base+".docx")
	default:
		return fmt.Errorf("unsupported format: %s", outputFormat)
	}
}

func main() {
	var pdfPath, outputFormat string

	fmt.Print("Enter path to PDF (e.g., example.pdf): ")
	fmt.Scanln(&pdfPath)

	fmt.Print("Convert to (txt/docx): ")
	fmt.Scanln(&outputFormat)

	err := convertPDF(pdfPath, outputFormat)
	if err != nil {
		log.Fatalf("Conversion failed: %v", err)
	} else {
		fmt.Println("Conversion successful! Check the output/ folder.")
	}
}
