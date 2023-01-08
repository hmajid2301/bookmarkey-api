//go:build tools
// +build tools

package main

import (
	_ "github.com/boumenot/gocover-cobertura"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/jstemmer/go-junit-report/v2"
	_ "golang.org/x/tools/cmd/goimports"
)
