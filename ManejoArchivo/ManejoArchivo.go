package Manejoarchivo

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	Modelo "proyecto/Modelo"

	"github.com/sqweek/dialog"
)

func seleccionarArchivo() (string, error) {
	file, err := dialog.File().Title("Seleccionar archivo").Filter("Todos los archivos", "*.*").Load()
	if err != nil {
		return "", err
	}
	return filepath.ToSlash(file), nil
}

func obtenerArchivo(ruta string) (*os.File, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	return archivo, nil
}

func obtenerContenido(archivo *os.File) ([]byte, error) {
	contenido, err := io.ReadAll(archivo)
	if err != nil {
		return nil, err
	}
	return contenido, nil
}

func DecodificarContenido() (Modelo.Gramatica, error) {
	ruta, err := seleccionarArchivo()
	archivo, err := obtenerArchivo(ruta)
	defer archivo.Close()
	contenido, err := obtenerContenido(archivo)
	var gramatica Modelo.Gramatica
	error := json.Unmarshal(contenido, &gramatica)
	if error != nil {
		fmt.Println("Error: ", err)
	}
	return gramatica, nil
}
