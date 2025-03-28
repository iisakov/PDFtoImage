package main

import (
	"PDFtoImage/internal/directorer"
	"PDFtoImage/internal/imager"
	"log"
)

func main() {
	fPath := "example.pdf"
	dPath := "source/output"

	ok, err := directorer.Create(fPath, dPath)
	if !ok {
		log.Printf("Что то пошло не так: %v\n", err)
	}

	ok, err = imager.Create(fPath, dPath)
	if !ok {
		log.Printf("Что то пошло не так: %v\n", err)
	}

}
