package main

import (
	"%PKG%/types/apis/some.api.group/v1"
	"github.com/rancher/norman/generator"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := generator.DefaultGenerate(v1.Schemas, "%PKG%/types", false, nil); err != nil {
		logrus.Fatal(err)
	}
}
