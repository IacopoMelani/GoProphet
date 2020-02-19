package file

import (
	"io"
	"net/http"
	"os"
)

// CheckAndCreateDir - Si occupa di controllare se esiste una cartella, se non fosse la crea
func CheckAndCreateDir(path string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

// GetFileFromURL - Si occupa di scaricare un file e spostarlo nella destinazione scelta
func GetFileFromURL(url string, dest string) {

	out, err := os.Create(dest)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
