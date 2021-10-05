package main

import (
	"fmt"
	"os"

	"github.com/br3w0r/goitopdf/itopdf"
)

// prints help to stdout
func displayHelp() {

        fmt.Println("goitopdf images_dir output_filename")
}

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
        // Display help if no argument is provided
        if len(os.Args) < 2 {
                displayHelp()
                return
        }

        // support --help as well
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		displayHelp()
		return
	}

	err := walkAndSave(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("Failed to create pdf: %v\n", err)
	}
}
