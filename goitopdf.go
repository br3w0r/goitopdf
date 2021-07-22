package main

import (
	"fmt"
	"os"

	"github.com/br3w0r/goitopdf/itopdf"
)

func walkAndSave(dir string, out string) error {
	pdf := itopdf.NewInstance()

	err := pdf.WalkDir(dir, func(path string) {
		fmt.Println("Adding new image: " + path)
	})
	if err != nil {
		return err
	}

	fmt.Println("Saving file...")

	err = pdf.Save(out)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if os.Args[1] == "-h" {
		fmt.Println("goitopdf images_dir output_filename")
		return
	}

	err := walkAndSave(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("Failed to create pdf: %v\n", err)
	}
}
