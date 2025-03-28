package imager

import (
	"fmt"
	"github.com/gen2brain/go-fitz"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Create(fPath, dPath string) (bool, error) {
	fnl := strings.Split(fPath, ".")
	fName := fnl[len(fnl)-2]

	// Открываем PDF-файл
	doc, err := fitz.New(fPath)
	if err != nil {
		log.Printf("Ошибка при открытии файла: %v\n", err)
		return false, fmt.Errorf("ошибка при открытии файла: %v\n", err)
	}
	defer func() { _ = doc.Close() }()

	// Обрабатываем каждую страницу
	for i := 0; i < doc.NumPage(); i++ {
		var img image.Image
		var file *os.File

		// Рендерим страницу в изображение
		img, err = doc.Image(i)
		if err != nil {
			log.Printf("Ошибка рендера страницы %d: %v\n", i, err)
			continue
		}

		// Сохраняем в PNG
		var outputPath string
		if doc.NumPage() == 1 {
			outputPath = filepath.Join(dPath, fmt.Sprintf("%s/%s.png", fName, fName))
		} else {
			outputPath = filepath.Join(dPath, fmt.Sprintf("%s/%s_%d.png", fName, fName, i+1))
		}

		file, err = os.Create(outputPath)
		if err != nil {
			log.Printf("Ошибка при создании картинки страницы: %v\n", err)
			continue
		}

		if err = png.Encode(file, img); err != nil {
			log.Printf("Ошибка кодирования в PNG: %v\n", err)
		} else {
			log.Printf("Сохранили: %s\n", outputPath)
		}

		_ = file.Close()
	}

	return true, err
}
