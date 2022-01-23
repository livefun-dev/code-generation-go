package generator

import (
	"strings"
	"testing"

	_ "embed"

	"github.com/matryer/is"
)

//go:embed example.go.txt
var example string

//go:embed example_result.go.txt
var example_result string

func TestGenerator(t *testing.T) {
	is := is.New(t)
	inputReader := strings.NewReader(example)
	res, err := RunGenerator(inputReader, "main")

	is.NoErr(err)
	is.Equal(res, example_result)
}
