package main

import (
	"fmt" // Chama-se "Format", mas todo mundo pronuncia "Fâmiti"
)

var y09 = "olá bom dia" // Tem uma abrangencia "PackageLevelScope", ou seja é Global
var y11 int

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

	//[Video 11]
	// Tipos em Go são estaticos
	var x11 int = 10 //Obs: Normalmente não se faz uma declaração 'var' dentro de um bloco,Porém farei apenas por didática
	// x11 = 20.5 //Descomentar esta linha causa erro, pois na linha anterior declaramos o tipo como inteiro
	fmt.Printf("%v, %T\n\n", x11, x11)

	y11 = 11 // Obs: Quando declaramos algo globalmente(ou seja, "PackageLevelScope") sem valor, so poderemos atribuir algo dentro de um escopo

	//[Video 12]
	var x12 int
	x12 = 10
	fmt.Printf("%v, %T\n", x12, x12)
	x12 = 20
	fmt.Printf("%v, %T\n", x12, x12)

	// O "ValorZero" das variaveis é o que ela 'armazenam' entre serem declaradas e inicializadas, e para cada tipo vale:
	/*
		int:                                                      0
		float:                                                    0.0
		booleans:                                                 false
		strings:                                                  ""
		ponteiros, funcoes, interfaces, slices, channels, maps:   nil
	*/
	var a12 int
	var b12 float64
	var c12 string
	var d12 bool
	fmt.Printf("%v, %T\n", a12, a12)
	fmt.Printf("%v, %T\n", b12, b12)
	fmt.Printf("%v, %T\n", c12, c12)
	fmt.Printf("%v, %T\n", d12, d12)

	//[Video 13]
	x13 := "oi bom dia\ncomo vai?\tespero que \"tudo bem\"."
	fmt.Printf("%v, %T\n", x13, x13)
	x13 = `oi bom dia
como vai?	espero que "tudo bem".`
	fmt.Printf("%v, %T\n", x13, x13)

	k := "oi"
	fmt.Print(k)   //Print sem adicionar uma nova linha no final ( \n )
	fmt.Println(k) //Print adicionando uma nova linha no final ( \n )

	l := "bom dia"
	fmt.Println(l)

	z13 := fmt.Sprint(k, " ", l)
	fmt.Println(z13)

	//[Video 14]
	type hotdog int
	var b14 hotdog = 10
	fmt.Printf("%T %v\n", b14, b14)
	x14 := 10
	fmt.Printf("%T %v\n", x14, x14)
	//b14 = x14 //Essa linha da erro, pois são tipos diferentes, e variaveis instanciadas com um certo tipo nunca mudam até sua "morte"

	//[Video 15]
	x14 = int(b14)
	fmt.Printf("%T %v\n", x14, x14)

	// aqui termina
}

func qualquerCoisa10(x int) {
	fmt.Println(y09)
	fmt.Println(x)
}
