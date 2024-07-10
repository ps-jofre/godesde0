package ejercicios

import (
	"strconv"
)

func FuncionEjercicio(valor1 string) (int, string) {

	entero, err := strconv.Atoi(valor1)
	if err != nil {
		return 0, "error" + err.Error()
	}
	if entero > 100 {
		return entero, "es mayor a 100"
	} else {
		return entero, "es menor a 100"
	}
}
