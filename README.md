# goitopdf

A simple CLI tool written in Go to convert images in a specified directory to pdf

Keep Going ;)

## Usage

### As an executable

`goitopdf images_dir output`

- images_dir - path to your images
- output - path and filename to save your created pdf to

#### Example

`goitopdf "/home/user/my_favorite_manga" ./my_favourite_manga.pdf`

### As a module

The module provides `itopdf` package which contains a structure that implements a `IToPDF` interface.

Basic usage:

```golang
package main

import "github.com/br3w0r/goitopdf/itopdf"

func main() {
    pdf := itopdf.NewInstance()
    
    err := pdf.WalkDir("/images/dir/", nil)
    if err != nil {
        panic(err)
    }

    err = pdf.Save("/output/dir/output_name.pdf")
    if err != nil {
        panic(err)
    }
}
```

You can also implement your images iteration algorithm by using `IToPDF`'s `AddImage` method.

## License

goitopdf is released under the MIT License and copyrighted by Daniil Dirun.
gofpdf is released under the MIT License. It is copyrighted by Dave Barnes and the contributors acknowledged in [documentation to the library](https://github.com/phpdave11/gofpdf/blob/master/README.md).
