package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "embed"

	"github.com/livefun/codege-test/pkg/generator"
)

func main() {
	packageName := os.Getenv("GOPACKAGE")
	fileName := os.Getenv("GOFILE")
	fmt.Printf("file: %v\n", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	generatedTextContent, err := generator.RunGenerator(file, packageName)
	if err != nil {
		log.Fatalf("panic: %v\n", err)
	}

	outFile := strings.Split(fileName, ".")[0] + ".metrics.go"
	f, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(generatedTextContent)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Printf("done: out on %v\n", outFile)
}
