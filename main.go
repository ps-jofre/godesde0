package main

import (
	"fmt"

	"github.com/ps-jofre/godesde0/ejercicios"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariable()

	/*
		estado, texto := variables.ConvietoText(2578)
		fmt.Println(estado)
		fmt.Println(texto)
	*/
	/*
		if os := runtime.GOOS; os == "linux" || os == "OS X." {
			fmt.Println("esto no es Windows, es ", os)
		} else {
			fmt.Println("Esto es Windows")
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es Linux")
		case "darwin":
			fmt.Println("Esto es Darwin")
		default:
			fmt.Printf("%s \n", os)
		}
	*/

	numero, texto := ejercicios.FuncionEjercicio("asdads")
	fmt.Println(numero)
	fmt.Print(texto)
}
