package main

import (
	"fmt" // Chama-se "Format", mas todo mundo pronuncia "Fâmiti"
)

func main() {
	// aqui começa

	//[Video 07]
	//pacote.identificador("Hello, World!")
	fmt.Println("Hello, World!")

	//[Video 08]
	numeroDeBytes, erros := fmt.Println("Hello, World!", "Oi galera!", 100)
	fmt.Println(numeroDeBytes, erros) //Como GO não permite a instanciacao de variaveis que não estejam sendo usadas estou "printando-os"

	// Outra forma de lidar com isso quando queremos apenas usar um dos valores retornados é usar o '_' dessa forma:
	_, numErros := fmt.Println("Hello, World!", "Oi galera!", 100)
	fmt.Println(numErros)

	x := 16
	y := "string"
	z := true
	fmt.Println(x, y, z)

	//[Video 09]

	// aqui termina
}
