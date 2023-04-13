package Controller

import (
	"fmt"
	Manejoarchivo "proyecto/ManejoArchivo"
	Modelo "proyecto/Modelo"
)

func primeros() (map[string][]string, error) {
	gramatica, err := Manejoarchivo.DecodificarContenido()
	if err != nil {
		fmt.Println("Error al decodificar", err)
	}
	var primeros = make(map[string][]string)
	for i := 0; i < len(gramatica.NoTerminales); i++ {
		NoTerminal := gramatica.NoTerminales[i]

		primer, err := obtenerPrimerosNoTer(NoTerminal, gramatica)
		if err != nil {
			fmt.Println(" obtener prim no ter")
		}
		primeros[NoTerminal] = primer
	}
	return primeros, nil
}

func obtenerPrimerosNoTer(NoTerminal string, gramatica Modelo.Gramatica) ([]string, error) {
	Prodiccion := gramatica.Expresiones[NoTerminal]
	var primeros []string

	if len(Prodiccion[1]) > 0 {
		if contiene(Prodiccion[0][0], gramatica.NoTerminales) {
			primer, err := obtenerPrimerosNoTer(Prodiccion[0][0], gramatica)
			for i := 0; i < len(primer); i++ {
				primeros = append(primeros, primer[i])
			}
			if err != nil {
				fmt.Println(err)
			}
		} else if contiene(Prodiccion[1][0], gramatica.NoTerminales) {
			primer, err := obtenerPrimerosNoTer(Prodiccion[1][0], gramatica)
			for i := 0; i < len(primer); i++ {
				primeros = append(primeros, primer[i])
			}
			if err != nil {
				fmt.Println(err)
			}
		} else {
			primeros = append(primeros, Prodiccion[0][0])
			primeros = append(primeros, Prodiccion[1][0])
		}

	} else {
		if contiene(Prodiccion[0][0], gramatica.NoTerminales) {
			primer, err := obtenerPrimerosNoTer(Prodiccion[0][0], gramatica)
			for i := 0; i < len(primer); i++ {
				primeros = append(primeros, primer[i])
			}
			if err != nil {
				fmt.Println(err)
			}
		} else {
			primeros = append(primeros, Prodiccion[0][0])
		}
	}
	return primeros, nil
}

func contiene(s string, lista []string) bool {
	for _, v := range lista {
		if v == s {
			return true
		}
	}
	return false
}

func ObtenerCadenas() string {
	primeros, err := primeros()
	if err != nil {
		fmt.Println("Error al obtener primeros")
	}
	var primero string
	for clave, valor := range primeros {
		var primer string
		for i := 0; i < len(valor); i++ {
			primer += valor[i] + "  "
		}
		primero += "Prim(" + clave + ") = { " + primer + "}" + "\n"
		primer = ""
	}
	return primero
}
