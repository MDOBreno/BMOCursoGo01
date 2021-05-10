package main

import (
	"fmt" // Chama-se "Format", mas todo mundo pronuncia "Fâmiti"
)

var y09 = "olá bom dia" // Tem uma abrangencia "PackageLevelScope", ou seja é Global

func main() { // Obs: '{}' == Escopo == Bloco
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
	x = 10
	y = "olá bom dia"

	fmt.Println(x, y)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	// Repare que, apesar do ':=' Ser apenas para declaração(sem ser atribuição), é possível contato que pelo -1 das variáveis esteja sendo declarada:
	x, x09 := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("x09: %v, %T\n", x09, x09)

	//Não podemos usar keywords(Palavras reservadas) como Identificadores(variaveis,classes,pacotes,etc..):
	//package := 10
	//fmt.Printf(package)

	z09 := 10 + 10
	fmt.Println(z09, "\n")

	//Expressões:
	a09 := 10 == 10
	fmt.Println(a09)
	a09 = 10 < 10
	fmt.Println(a09, "\n")

	//[Video 10]
	y10 := 10
	qualquerCoisa10(y10)

	// aqui termina
}

func qualquerCoisa10(x int) {
	fmt.Println(y09)
	fmt.Println(x)
}
