package directorer

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Create(fPath, dPath string) (bool, error) {
	fnl := strings.Split(fPath, ".")
	fileName := fnl[len(fnl)-2]

	if err := os.MkdirAll(dPath+"/"+fileName, 0755); err != nil {
		log.Printf("Ошибка при создании папки %s: %v\n", dPath, err)
		return false, fmt.Errorf("ошибка при создании папки %s: %v\n", dPath, err)
	}

	return true, nil
}
