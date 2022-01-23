package generator

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"io"
	"strings"
	"text/template"
)

//go:embed template.txt
var templateString string

const commentPrefix = "//+prom:metric"

type Variables struct {
	PackageName string
	Metrics     []Metric
}

type Metric struct {
	Name     string
	Variable string
	Type     string
}

func RunGenerator(content io.Reader, packageName string) (string, error) {
	tmp, err := template.New("template").Parse(templateString)
	if err != nil {
		return "", fmt.Errorf("cannot create template, %w", err)
	}

	var res string
	buf := bytes.NewBufferString(res)

	variabels := Variables{
		PackageName: packageName,
	}

	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, commentPrefix) {
			if ok := scanner.Scan(); !ok {
				return "", fmt.Errorf("not expected EOF")
			}
			nextLine := strings.TrimSpace(scanner.Text())
			varName := strings.Split(nextLine, ".")[0]
			metricName := strings.Split(strings.Split(line, "name:")[1], " ")[0]
			metricType := strings.Split(strings.Split(line, "+prom:metric:")[1], " ")[0]

			variabels.Metrics = append(variabels.Metrics, Metric{
				Name:     metricName,
				Variable: varName,
				Type:     metricType,
			})
		}
	}

	if err := tmp.Execute(buf, variabels); err != nil {
		return "", err
	}

	formattedOut, err := format.Source(buf.Bytes())
	if err != nil {
		return "", err
	}
	return string(formattedOut), nil
}
