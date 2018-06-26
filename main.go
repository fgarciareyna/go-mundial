package main

import (
	"./seleccion"
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	rand.Seed(time.Now().Unix())

	golesArgentina := rand.Intn(5)
	golesNigeria := rand.Intn(5)

	fmt.Println(fmt.Sprintf("Argentina %v - Nigeria %v", golesArgentina, golesNigeria))

	argentina := seleccion.Doparti{
		Gana: golesArgentina > golesNigeria,
		DiferenciaGoles: golesArgentina - golesNigeria,
	}

	golesIslandia := rand.Intn(5)
	golesCroacia := rand.Intn(5)

	fmt.Println(fmt.Sprintf("Islandia %v - Croacia %v", golesIslandia, golesCroacia))

	islandia := seleccion.Doparti{
		Gana: golesIslandia > golesCroacia,
		DiferenciaGoles: golesIslandia - golesCroacia,
	}

	clasificamos := seleccion.Clasificamos(argentina, islandia)

	fmt.Println()

	if clasificamos {
		fmt.Println("Grande Sampaoli, siempre te tuve fe")
	} else {
		fmt.Println("Sampaoli mejor quedate en Rusia")
	}
}
