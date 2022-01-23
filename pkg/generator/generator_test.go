package generator_test

import (
	"strings"
	"testing"

	_ "embed"

	"github.com/livefun/codege-test/pkg/generator"
	"github.com/matryer/is"
)

//go:embed tests/example.go.txt
var example string

//go:embed tests/example_result.go.txt
var example_result string

func TestGenerator(t *testing.T) {
	is := is.New(t)
	inputReader := strings.NewReader(example)
	res, err := generator.RunGenerator(inputReader, "main")

	is.NoErr(err)
	is.Equal(res, example_result)
}
