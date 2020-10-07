package main

import (
	"errors"
	"fmt"
	"log"
)

var deuRuim = errors.New("[var errors.New] número negativo não pode")

type customError struct {
	msg string
}

func (e customError) Error() string {
	return fmt.Sprintf("[custom struct error] %s", e.msg)
}

func main() {
	resp, err := errorsNew(10)
	handleErrAndResponse("errors.New", resp, err)
	printLine()

	resp, err = errorsNew(-220)
	handleErrAndResponse("errors.New", resp, err)
	printLine()

	resp, err = varErrorsNew(-124)
	handleErrAndResponse("var errors.New", resp, err)
	printLine()

	resp, err = fmtErrorf(-543)
	handleErrAndResponse("fmt.Errorf", resp, err)
	printLine()

	resp, err = varFmtErrorf(-5435)
	handleErrAndResponse("var fmt.Errorf", resp, err)
	printLine()

	resp, err = errorFromStruct(-12890)
	handleErrAndResponse("error from struct", resp, err)
	printLine()
}

func printLine() {
	fmt.Println("\n----------------------")
	fmt.Println()
}

func errorsNew(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("[errors.New] numero negativo não pode")
	}
	return 42, nil
}

func varErrorsNew(n float64) (float64, error) {
	fmt.Printf("Tipo da variável deuRuim: %T\n", deuRuim)

	if n < 0 {
		return 0, deuRuim
	}
	return 42, nil
}

func fmtErrorf(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("[fmt.Errorf] numero negativo não pode: %v", n)
	}
	return 42, nil
}

func varFmtErrorf(n float64) (float64, error) {
	if n < 0 {
		myError := fmt.Errorf("[var fmt.Errorf] numero negativo não pode: %v", n)
		return 0, myError
	}
	return 42, nil
}

func errorFromStruct(n float64) (float64, error) {
	fmt.Printf("Tipo do struct customError: %T\n", customError{})

	if n < 0 {
		return 0, customError{fmt.Sprintf("numero negativo não pode: %v", n)}
	}
	return 42, nil
}

func handleErrAndResponse(desc string, resp interface{}, err error) {
	log.Println("Testando:", desc)
	if err != nil {
		log.Println("Deu ruim aqui:", err)
		return
	}
	log.Println("deu tudo certo e a resposta é:", resp)
}
