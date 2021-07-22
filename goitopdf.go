package main

import (
	"fmt"
	"os"

	"github.com/br3w0r/goitopdf/itopdf"
)

func main() {
	if os.Args[1] == "-h" {
		fmt.Println("goitopdf images_dir output_filename")
		return
	}

	pdf := itopdf.NewInstance()
	err := pdf.WalkDir(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("Failed to create pdf: %v\n", err)
	}
}
