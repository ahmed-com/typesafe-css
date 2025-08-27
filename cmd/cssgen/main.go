// Package main implements the CSS property code generator.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Spec represents the CSS specification data.
type Spec struct {
	Properties []PropertySpec `json:"properties"`
	Version    string         `json:"version"`
}

// PropertySpec represents a single CSS property specification.
type PropertySpec struct {
	Name     string   `json:"name"`     // "background-repeat"
	Keywords []string `json:"keywords"` // ["repeat", "no-repeat", ...] (may be empty)
	Syntax   string   `json:"syntax"`   // e.g., "<color> | <image> | ...", used later for richer types
	Status   string   `json:"status"`   // "standard" | "experimental" | "deprecated"
}

func main() {
	var (
		inPath            = flag.String("in", "", "Path to spec.json or directory containing spec files")
		outPath           = flag.String("out", "./cssgen", "Output directory for generated files")
		pkg               = flag.String("pkg", "css", "Package name for generated files")
		strict            = flag.Bool("strict", false, "Fail if unknown/unsupported syntax segments")
		allowExperimental = flag.Bool("allow-experimental", false, "Include non-standard properties")
	)
	flag.Parse()

	if *inPath == "" {
		log.Fatal("Error: -in flag is required")
	}

	// Load spec
	spec, err := loadSpec(*inPath)
	if err != nil {
		log.Fatalf("Error loading spec: %v", err)
	}

	// Normalize spec
	normalized := normalize(spec, *allowExperimental, *strict)

	// Ensure output directory exists
	if err := os.MkdirAll(*outPath, 0755); err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	// Generate files
	if err := generateProperties(normalized, *outPath, *pkg); err != nil {
		log.Fatalf("Error generating properties: %v", err)
	}

	if err := generateKeywords(normalized, *outPath, *pkg); err != nil {
		log.Fatalf("Error generating keywords: %v", err)
	}

	if err := generateSetters(normalized, *outPath, *pkg); err != nil {
		log.Fatalf("Error generating setters: %v", err)
	}

	fmt.Printf("Generated CSS property definitions in %s\n", *outPath)
}

// loadSpec loads the CSS specification from a JSON file or directory.
func loadSpec(path string) (Spec, error) {
	var spec Spec

	// Check if path is a file or directory
	info, err := os.Stat(path)
	if err != nil {
		return spec, fmt.Errorf("cannot access path %s: %w", path, err)
	}

	if info.IsDir() {
		// Look for spec.json in directory
		specPath := filepath.Join(path, "spec.json")
		return loadSpecFromFile(specPath)
	}

	// Load from single file
	return loadSpecFromFile(path)
}

// loadSpecFromFile loads the CSS specification from a JSON file.
func loadSpecFromFile(filePath string) (Spec, error) {
	var spec Spec

	data, err := os.ReadFile(filePath)
	if err != nil {
		return spec, fmt.Errorf("cannot read file %s: %w", filePath, err)
	}

	if err := json.Unmarshal(data, &spec); err != nil {
		return spec, fmt.Errorf("cannot parse JSON from %s: %w", filePath, err)
	}

	return spec, nil
}

// normalize processes the spec to ensure consistency.
func normalize(spec Spec, allowExperimental, strict bool) Spec {
	var normalized Spec
	normalized.Version = spec.Version

	for _, prop := range spec.Properties {
		// Filter by status
		if prop.Status == "deprecated" {
			continue
		}
		if prop.Status == "experimental" && !allowExperimental {
			continue
		}

		// Normalize property name to lowercase kebab-case
		name := strings.ToLower(prop.Name)
		name = strings.ReplaceAll(name, "_", "-")

		// Normalize keywords: unique, sorted, lowercase
		keywordSet := make(map[string]bool)
		for _, kw := range prop.Keywords {
			normalized := strings.ToLower(strings.TrimSpace(kw))
			if normalized != "" {
				keywordSet[normalized] = true
			}
		}

		var keywords []string
		for kw := range keywordSet {
			keywords = append(keywords, kw)
		}
		sort.Strings(keywords)

		normalized.Properties = append(normalized.Properties, PropertySpec{
			Name:     name,
			Keywords: keywords,
			Syntax:   prop.Syntax,
			Status:   prop.Status,
		})
	}

	// Sort properties by name for consistent output
	sort.Slice(normalized.Properties, func(i, j int) bool {
		return normalized.Properties[i].Name < normalized.Properties[j].Name
	})

	return normalized
}

// generateProperties creates the properties_gen.go file with property constants.
func generateProperties(spec Spec, outPath, pkg string) error {
	var buf strings.Builder

	// Header
	buf.WriteString(generateHeader(spec.Version))
	buf.WriteString(fmt.Sprintf("package %s\n\n", pkg))
	
	// Import css package if we're generating a separate package
	if pkg != "css" {
		buf.WriteString("import \"github.com/ahmed-com/typesafe-css/css\"\n\n")
	}
	
	buf.WriteString("// Property constants for CSS properties.\n")
	buf.WriteString("const (\n")

	for _, prop := range spec.Properties {
		constName := propToConstName(prop.Name)
		propType := "Property"
		if pkg != "css" {
			propType = "css.Property"
		}
		buf.WriteString(fmt.Sprintf("\t%s %s = %q\n", constName, propType, prop.Name))
	}

	buf.WriteString(")\n")

	// Format and write
	formatted, err := format.Source([]byte(buf.String()))
	if err != nil {
		return fmt.Errorf("failed to format generated properties: %w", err)
	}

	outFile := filepath.Join(outPath, "properties_gen.go")
	return os.WriteFile(outFile, formatted, 0644)
}

// generateKeywords creates the keywords_gen.go file with keyword types and constants.
func generateKeywords(spec Spec, outPath, pkg string) error {
	var buf strings.Builder

	// Header
	buf.WriteString(generateHeader(spec.Version))
	buf.WriteString(fmt.Sprintf("package %s\n\n", pkg))
	buf.WriteString("// Keyword types and constants for CSS property values.\n\n")

	for _, prop := range spec.Properties {
		if len(prop.Keywords) == 0 {
			continue // Skip properties without finite keyword sets
		}

		typeName := propToTypeName(prop.Name)
		buf.WriteString(fmt.Sprintf("// %s represents values for the %s property.\n", typeName, prop.Name))
		buf.WriteString(fmt.Sprintf("type %s string\n\n", typeName))

		buf.WriteString(fmt.Sprintf("// %s constants.\n", typeName))
		buf.WriteString("const (\n")

		for _, keyword := range prop.Keywords {
			constName := typeName + keywordToConstSuffix(keyword)
			buf.WriteString(fmt.Sprintf("\t%s %s = %q\n", constName, typeName, keyword))
		}

		buf.WriteString(")\n\n")
	}

	// Format and write
	formatted, err := format.Source([]byte(buf.String()))
	if err != nil {
		return fmt.Errorf("failed to format generated keywords: %w", err)
	}

	outFile := filepath.Join(outPath, "keywords_gen.go")
	return os.WriteFile(outFile, formatted, 0644)
}

// generateSetters creates the setters_gen.go file with type-safe setter functions.
func generateSetters(spec Spec, outPath, pkg string) error {
	var buf strings.Builder

	// Header
	buf.WriteString(generateHeader(spec.Version))
	buf.WriteString(fmt.Sprintf("package %s\n\n", pkg))
	
	// Import css package if we're generating a separate package
	if pkg != "css" {
		buf.WriteString("import \"github.com/ahmed-com/typesafe-css/css\"\n\n")
	}
	
	buf.WriteString("// Type-safe setter functions for CSS properties.\n\n")

	for _, prop := range spec.Properties {
		if len(prop.Keywords) == 0 {
			continue // Skip properties without finite keyword sets for now
		}

		funcName := propToSetterName(prop.Name)
		typeName := propToTypeName(prop.Name)
		constName := propToConstName(prop.Name)
		
		// Determine types based on package
		declType := "Decl"
		keywordType := "Keyword"
		setFunc := "Set"
		if pkg != "css" {
			declType = "css.Decl"
			keywordType = "css.Keyword"
			setFunc = "css.Set"
		}

		buf.WriteString(fmt.Sprintf("// %s creates a declaration for the %s property.\n", funcName, prop.Name))
		buf.WriteString(fmt.Sprintf("func %s(v %s) %s {\n", funcName, typeName, declType))
		buf.WriteString(fmt.Sprintf("\treturn %s(%s, %s(string(v)))\n", setFunc, constName, keywordType))
		buf.WriteString("}\n\n")
	}

	// Format and write
	formatted, err := format.Source([]byte(buf.String()))
	if err != nil {
		return fmt.Errorf("failed to format generated setters: %w", err)
	}

	outFile := filepath.Join(outPath, "setters_gen.go")
	return os.WriteFile(outFile, formatted, 0644)
}

// generateHeader creates the common header for generated files.
func generateHeader(version string) string {
	timestamp := time.Now().Format(time.RFC3339)
	return fmt.Sprintf(`// Code generated by cssgen; DO NOT EDIT.
// Source: spec version %s at %s

`, version, timestamp)
}

// propToConstName converts a property name to a Go constant name.
// e.g., "background-color" -> "BackgroundColor"
func propToConstName(propName string) string {
	parts := strings.Split(propName, "-")
	for i, part := range parts {
		if part != "" {
			parts[i] = strings.Title(part)
		}
	}
	return strings.Join(parts, "")
}

// propToTypeName converts a property name to a Go type name.
// e.g., "background-color" -> "BackgroundColorVal"
func propToTypeName(propName string) string {
	return propToConstName(propName) + "Val"
}

// propToSetterName converts a property name to a Go setter function name.
// e.g., "background-color" -> "SetBackgroundColor"
func propToSetterName(propName string) string {
	return "Set" + propToConstName(propName)
}

// keywordToConstSuffix converts a keyword to a Go constant suffix.
// e.g., "no-repeat" -> "NoRepeat"
func keywordToConstSuffix(keyword string) string {
	parts := strings.Split(keyword, "-")
	for i, part := range parts {
		if part != "" {
			parts[i] = strings.Title(part)
		}
	}
	return strings.Join(parts, "")
}