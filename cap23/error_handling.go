package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// É boa prática lidar com o erro (if err != nil) logo após a declaração do erro (result, err := someFunc())
// Essa regra só não se aplica ao fmt.Println (que pode retornar numero de bytes e erro)
// 		-> n, err := fmt.Println() [NOK]
// 		-> n, _ := fmt.Println() [OK]
// 		-> fmt.Println() [OK]
// Recomendação: usar sempre log.*

func main() {
	writeOnFile()
	printLine()
	// OUTPUT:
	// $ go run cap23/error_handling.go
	// #=> arquivo names.txt gravado
	// $ cat names.txt
	// #=> alala ô lalaô lalaô eô
	// OUTPUT COM ERRO:
	// $ chmod 000 names.txt
	// $ go run cap23/error_handling.go
	// #=> open names.txt: permission denied

	dealingWithErrors_Simple()
	printLine()
	// OUTPUT:
	// a wild error appeared (fmt.Println): open inexistent-file.txt: no such file or directory
	// 2020/10/06 20:33:48 a wild error appeared (log.Println): open inexistent-file.txt: no such file or directory

	dealingWithErrors_SetOutput()
	printLine()
	// OUTPUT:
	// $ go run cap23/error_handling.go
	// #=> Dá uma olhadela no arquivo log.txt
	// $ cat log.txt
	// #=> 2020/10/06 20:38:14 a wild error appeared (log.Println): open inexistent-file.txt: no such file or directory

	canRunCmd := canRun("fatal")
	if canRunCmd {
		dealingWithErrors_Fatalln()
		printLine()
		// OUTPUT:
		// 2020/10/06 20:58:04 a wild error appeared (log.Fatalln): open inexistent-file.txt: no such file or directory
		// exit status 1
	}

	canRunCmd = canRun("panicln")
	if canRunCmd {
		dealingWithErrors_Panicln()
		printLine()
		// OUTPUT:
		// 2020/10/06 20:55:11 a wild error appeared (log.Println): open inexistent-file.txt: no such file or directory
		// Quando rola um fatal (os.Exit(2)), as defer funcs rodam
		// panic: a wild error appeared (log.Panicln): open inexistent-file.txt: no such file or directory

		// goroutine 1 [running]:
		// log.Panicln(0xc000107f10, 0x2, 0x2)
		// 		/usr/local/go/src/log/log.go:365 +0xae
		// main.dealingWithErrors_Panicln()
		// 		$GOPATH/github.com/camilacampos/learn-go/cap23/error_handling.go:126 +0x129
		// main.main()
		// 		$GOPATH/github.com/camilacampos/learn-go/cap23/error_handling.go:55 +0x8e
		// exit status 2
	}
	canRunCmd = canRun("panic")
	if canRunCmd {
		dealingWithErrors_Panic()
		printLine()
		// OUTPUT (só não tem a primeira linha, do log, antes da defer func):
		// Quando rola um fatal (os.Exit(2)), as defer funcs rodam
		// panic: open inexistent-file.txt: no such file or directory

		// goroutine 1 [running]:
		// main.dealingWithErrors_Panic()
		// 		/Users/camila.campos/go/src/github.com/camilacampos/learn-go/cap23/error_handling.go:153 +0x105
		// main.main()
		// 		/Users/camila.campos/go/src/github.com/camilacampos/learn-go/cap23/error_handling.go:73 +0xae
		// exit status 2
	}
}

func printLine() {
	fmt.Println("\n----------------------")
	fmt.Println()
}

func writeOnFile() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fmt.Println("arquivo names.txt gravado") // não é executado quando dá erro
	defer f.Close()                                // não é executado quando dá erro

	r := strings.NewReader("alala ô lalaô lalaô eô")

	_, err = io.Copy(f, r)
	if err != nil {
		fmt.Println(err)
		return // se desse erro aqui, o defer iria rodar
	}
}

func dealingWithErrors_Simple() {
	_, err := os.Open("inexistent-file.txt")
	if err != nil {
		fmt.Println("a wild error appeared (fmt.Println):", err) // comum, que a gente tá acostumada a rodar. Imprime a msg de erro na tela
		log.Println("a wild error appeared (log.Println):", err) // mais complexo. Imprime a msg de erro na tela, junto com um timestamp
	}
}

func dealingWithErrors_SetOutput() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f) // o SetOutput muda o output do STDOUT pro arquivo que for indicado
	defer log.SetOutput(os.Stdout)

	f2, err := os.Open("inexistent-file.txt")
	if err != nil {
		log.Println("a wild error appeared (log.Println):", err) // ao invés de escrever na tela, o log é escrito em `log.txt`
	}
	defer f2.Close()

	fmt.Println("Dá uma olhadela no arquivo log.txt")
}

func dealingWithErrors_Fatalln() {
	defer fmt.Println("Quando rola um fatal (os.Exit(1)), as defer funcs não rodam")

	_, err := os.Open("inexistent-file.txt")
	if err != nil {
		log.Fatalln("a wild error appeared (log.Fatalln):", err) // chama os.Exit(1) quando acontece isso aqui
	}
}

func dealingWithErrors_Panicln() {
	defer fmt.Println("Quando rola um fatal (os.Exit(2)), as defer funcs rodam")

	_, err := os.Open("inexistent-file.txt")
	if err != nil {
		log.Panicln("a wild error appeared (log.Panicln):", err) // faz um println, seguido de uma chamada pra panic()
		// o panic(), por sua vez, roda as defer funcs (dela e das funcs superiores), depois printa o erro, finaliza a execução (exit code 2) e, por fim, printa a stacktrace
		// pode ser recuperada com recover (veremos mais pra frente)
	}
}

func dealingWithErrors_Panic() {
	defer fmt.Println("Quando rola um fatal (os.Exit(2)), as defer funcs rodam")

	_, err := os.Open("inexistent-file.txt")
	if err != nil {
		panic(err) // funciona igual ao log.Panicln(), mas não faz o log (então não dá pra passar mensagemzinha), apenas printa o erro e a stack trace
	}
}

func canRun(cmd string) bool {
	fmt.Println("Devo rodar o", cmd, "? (S/N)")
	var resp string
	fmt.Scan(&resp)
	return strings.ToUpper(resp) == "S"
}
