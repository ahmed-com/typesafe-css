// Package main includes go:generate directives for code generation.
package main

//go:generate go run ./cmd/cssgen -in ./cmd/cssgen/spec/spec.json -out ./cssgen -pkg cssgen

func main() {
	// This file is only used for go:generate directives.
}