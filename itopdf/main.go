package itopdf

import (
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/phpdave11/gofpdf"
)

type IToPDF interface {
	AddImage(path string) error
	Save(path string) error

	// WalkDir uses AddImage and Save methods to iterate over a directory and add all images in it to pdf.
	// iterCallback - function that takes full path of an image that is being saved
	WalkDir(dir string, iterCallback func(path string)) error
}

type iToPDF struct {
	inst *gofpdf.Fpdf
}

var (
	imagetypes              = []string{".jpg", ".png"}
	ErrWrongImageType error = errors.New("wrong image type")
)

func NewInstance() (pdf IToPDF) {
	pdf = &iToPDF{
		inst: gofpdf.New("P", "mm", "", ""),
	}
	return
}

func (pdf *iToPDF) AddImage(path string) error {
	if !checkTypes(path) {
		return ErrWrongImageType
	}

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

	if pdf.inst.Err() {
		return errors.New("Error adding image")
	}
	return nil
}

func (pdf *iToPDF) Save(path string) error {
	return pdf.inst.OutputFileAndClose(path)
}

func (pdf *iToPDF) walkFunc(iterCallback func(path string)) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignoring directories
		if info.IsDir() {
			return nil
		}

		if iterCallback != nil {
			iterCallback(path)
		}

		err = pdf.AddImage(path)

		// Ignoring files of wrong extention
		if err == ErrWrongImageType {
			return nil
		}

		if err != nil {
			return err
		}

		return nil
	}
}

func (pdf *iToPDF) WalkDir(dir string, iterCallback func(path string)) error {
	err := filepath.Walk(dir, pdf.walkFunc(iterCallback))

	if err != nil {
		return err
	}

	return nil
}

func checkTypes(filename string) bool {
	for _, k := range imagetypes {
		if strings.HasSuffix(filename, k) {
			return true
		}
	}

	return false
}
