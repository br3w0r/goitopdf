package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
	"path/filepath"
	"github.com/phpdave11/gofpdf"
)

type PdfGen struct {
	inst *gofpdf.Fpdf
	width float64
}

func CreatePdf(width int, height int) (pdf *PdfGen) {
	size := gofpdf.SizeType{
		Wd: float64(width),
		Ht: float64(height),
	}
	initType := gofpdf.InitType{
		OrientationStr: "P",
		UnitStr: "mm",
		SizeStr: "",
		Size: size,
		FontDirStr: "",
	}

	pdf = &PdfGen{
		inst: gofpdf.NewCustom(&initType),
		width: float64(width),
	}
	return
}

func (pdf *PdfGen) AddImage(path string) error {
	pdf.inst.AddPage()
	pdf.inst.Image(path, 0, 0, pdf.width, 0, false, "", 0, "")

	return nil
}

func (pdf *PdfGen) Save(path string) error {
	return pdf.inst.OutputFileAndClose(path)
}

func main() {
	if os.Args[1] == "-h" {
		fmt.Println("pdf-gen images_dir page_width page_height output_filename")
		return
	}

	width, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	height, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	pdf := CreatePdf(width, height)
	first := true
	err = filepath.Walk(os.Args[1], func (path string, info os.FileInfo, err error) error {
		// Skipping first because it's the dir itself
		if first {
			first = false
			return nil	
		}

		if err != nil {
			return err
		}

		fmt.Println("Adding new image: " + path);
		
		err = pdf.AddImage(path)
		if err != nil {
			return err
		}

		return nil
	});

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Saving file...")
	err = pdf.Save(os.Args[4])

	if err != nil {
		log.Fatal(err)
	}
}
