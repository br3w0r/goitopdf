package main

import (
	"fmt"
	"github.com/phpdave11/gofpdf"
	"log"
	"os"
	"path/filepath"
	"strings"
	"image"
	_ "image/jpeg"
	_ "image/png"
)

var (
	imagetypes = []string{".jpg", ".png"}
)

func checkTypes(filename string) bool {

	for _, k := range imagetypes {

		if strings.HasSuffix(filename, k) {
			return true
		}

	}
	return false
}

type PdfGen struct {
	inst  *gofpdf.Fpdf
}

func CreatePdf() (pdf *PdfGen) {
	pdf = &PdfGen{
		inst:  gofpdf.New("P", "mm", "", ""),
	}
	return
}

func (pdf *PdfGen) AddImage(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return err
	}

	sizeType := gofpdf.SizeType{
		Wd: float64(image.Width),
		Ht: float64(image.Height),
	}

	pdf.inst.AddPageFormat("P", sizeType)
	pdf.inst.Image(path, 0, 0, sizeType.Wd, sizeType.Ht, false, "", 0, "")

	return nil
}

func (pdf *PdfGen) Save(path string) error {
	return pdf.inst.OutputFileAndClose(path)
}

func main() {
	if os.Args[1] == "-h" {
		fmt.Println("goitopdf images_dir output_filename")
		return
	}

	pdf := CreatePdf()
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		// Ignoring directories and files that are not images
		if info.IsDir() || !checkTypes(info.Name()) {
			return nil
		}

		fmt.Println("Adding new image: " + path)

		err = pdf.AddImage(path)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Saving file...")
	err = pdf.Save(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}
}
