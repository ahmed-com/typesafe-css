// Package main implements a validator for CSS specifications against MDN data.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// MDNProperty represents a CSS property from MDN data.
type MDNProperty struct {
	Syntax        string      `json:"syntax"`
	Status        string      `json:"status"`
	Inherited     interface{} `json:"inherited"`     // Can be bool or string
	Groups        []string    `json:"groups"`
	Initial       interface{} `json:"initial"`       // Can be string or array
	MDNURL        string      `json:"mdn_url"`
	Media         interface{} `json:"media"`         // Can be string or array
	AnimationType interface{} `json:"animationType"` // Can be string or array
	Percentages   interface{} `json:"percentages"`   // Can be string or array
	AppliesToo    interface{} `json:"appliesto"`     // Can be string or array
	Computed      interface{} `json:"computed"`      // Can be string or array
	Order         interface{} `json:"order"`         // Can be string or array
}

// MDNSyntax represents a syntax definition from MDN data.
type MDNSyntax struct {
	Syntax string `json:"syntax"`
}

// PropertySpec represents our property specification format.
type PropertySpec struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
	Syntax   string   `json:"syntax"`
	Status   string   `json:"status"`
}

// Spec represents our CSS specification data.
type Spec struct {
	Properties []PropertySpec `json:"properties"`
	Version    string         `json:"version"`
}

func main() {
	var (
		mdnPath     = flag.String("mdn", "", "Path to MDN data directory")
		specPath    = flag.String("spec", "", "Path to our spec.json file")
		outPath     = flag.String("out", "", "Output path for updated spec (optional)")
		verbose     = flag.Bool("v", false, "Verbose output")
		updateSpec  = flag.Bool("update", false, "Update spec with MDN data")
	)
	flag.Parse()

	if *mdnPath == "" || *specPath == "" {
		log.Fatal("Error: Both -mdn and -spec flags are required")
	}

	validator := &Validator{
		verbose: *verbose,
	}

	// Load MDN data
	if err := validator.loadMDNData(*mdnPath); err != nil {
		log.Fatalf("Error loading MDN data: %v", err)
	}

	// Load our spec
	if err := validator.loadSpec(*specPath); err != nil {
		log.Fatalf("Error loading spec: %v", err)
	}

	// Validate against MDN
	results := validator.validate()

	// Print results
	validator.printResults(results)

	// Update spec if requested
	if *updateSpec {
		updated := validator.updateSpecWithMDN()
		outFile := *outPath
		if outFile == "" {
			outFile = *specPath
		}
		if err := validator.writeSpec(updated, outFile); err != nil {
			log.Fatalf("Error writing updated spec: %v", err)
		}
		fmt.Printf("Updated spec written to %s\n", outFile)
	}
}

// Validator handles the validation logic.
type Validator struct {
	mdnProperties map[string]MDNProperty
	mdnSyntaxes   map[string]MDNSyntax
	spec          Spec
	verbose       bool
}

// ValidationResult represents the result of validating a property.
type ValidationResult struct {
	Property        string
	Status          string
	MDNKeywords     []string
	SpecKeywords    []string
	MissingKeywords []string
	ExtraKeywords   []string
	SyntaxMatch     bool
	Exists          bool
}

// loadMDNData loads properties and syntax data from MDN.
func (v *Validator) loadMDNData(mdnPath string) error {
	// Load properties
	propertiesFile := filepath.Join(mdnPath, "css", "properties.json")
	propertiesData, err := os.ReadFile(propertiesFile)
	if err != nil {
		return fmt.Errorf("cannot read properties.json: %w", err)
	}

	if err := json.Unmarshal(propertiesData, &v.mdnProperties); err != nil {
		return fmt.Errorf("cannot parse properties.json: %w", err)
	}

	// Load syntaxes
	syntaxesFile := filepath.Join(mdnPath, "css", "syntaxes.json")
	syntaxesData, err := os.ReadFile(syntaxesFile)
	if err != nil {
		return fmt.Errorf("cannot read syntaxes.json: %w", err)
	}

	if err := json.Unmarshal(syntaxesData, &v.mdnSyntaxes); err != nil {
		return fmt.Errorf("cannot parse syntaxes.json: %w", err)
	}

	if v.verbose {
		fmt.Printf("Loaded %d properties and %d syntax definitions from MDN\n", 
			len(v.mdnProperties), len(v.mdnSyntaxes))
	}

	return nil
}

// loadSpec loads our CSS specification.
func (v *Validator) loadSpec(specPath string) error {
	data, err := os.ReadFile(specPath)
	if err != nil {
		return fmt.Errorf("cannot read spec file: %w", err)
	}

	if err := json.Unmarshal(data, &v.spec); err != nil {
		return fmt.Errorf("cannot parse spec file: %w", err)
	}

	if v.verbose {
		fmt.Printf("Loaded spec with %d properties\n", len(v.spec.Properties))
	}

	return nil
}

// validate compares our spec against MDN data.
func (v *Validator) validate() []ValidationResult {
	var results []ValidationResult

	for _, prop := range v.spec.Properties {
		result := ValidationResult{
			Property:     prop.Name,
			SpecKeywords: prop.Keywords,
		}

		mdnProp, exists := v.mdnProperties[prop.Name]
		result.Exists = exists
		if !exists {
			result.Status = "NOT_FOUND_IN_MDN"
			results = append(results, result)
			continue
		}

		result.Status = mdnProp.Status

		// Extract keywords from MDN syntax
		mdnKeywords := v.extractKeywords(mdnProp.Syntax)
		result.MDNKeywords = mdnKeywords

		// Compare keywords
		result.MissingKeywords = difference(mdnKeywords, prop.Keywords)
		result.ExtraKeywords = difference(prop.Keywords, mdnKeywords)

		// Check syntax similarity
		result.SyntaxMatch = strings.Contains(mdnProp.Syntax, strings.Join(prop.Keywords, " | "))

		results = append(results, result)
	}

	return results
}

// extractKeywords extracts concrete keyword values from MDN syntax.
func (v *Validator) extractKeywords(syntax string) []string {
	var keywords []string
	seen := make(map[string]bool)

	// First, expand any type references (e.g., <display-box> -> "contents | none")
	expanded := v.expandSyntaxTypes(syntax)

	// Extract keywords using regex for simple alternatives (word | word | word)
	// Include numbers for font-weight and similar properties
	re := regexp.MustCompile(`\b([a-z-]+|\d+)\b`)
	matches := re.FindAllString(expanded, -1)

	for _, match := range matches {
		// Filter out obvious non-keywords
		if v.isKeyword(match) && !seen[match] {
			keywords = append(keywords, match)
			seen[match] = true
		}
	}

	sort.Strings(keywords)
	return keywords
}

// expandSyntaxTypes replaces type references with their syntax definitions.
func (v *Validator) expandSyntaxTypes(syntax string) string {
	// Replace <type-name> with the actual syntax from syntaxes.json
	re := regexp.MustCompile(`<([^>]+)>`)
	
	expanded := re.ReplaceAllStringFunc(syntax, func(match string) string {
		typeName := strings.Trim(match, "<>")
		if syntaxDef, exists := v.mdnSyntaxes[typeName]; exists {
			return syntaxDef.Syntax
		}
		return match
	})

	return expanded
}

// isKeyword determines if a string looks like a CSS keyword value.
func (v *Validator) isKeyword(s string) bool {
	// Skip common syntax elements that aren't keywords
	skip := []string{
		"and", "or", "not", "where", "is", "has", "any", "all",
		"inherit", "initial", "unset", "revert",
		"length", "percentage", "number", "integer", "string",
		"color", "image", "url", "calc", "var", "env", "attr",
		"global", "values", "keyword", "value", "property",
	}

	s = strings.ToLower(s)
	
	// Skip if too short (but allow single digits for font-weight)
	if len(s) < 2 && !regexp.MustCompile(`^\d+$`).MatchString(s) {
		return false
	}

	// Skip if contains characters that suggest it's not a simple keyword
	if strings.ContainsAny(s, "()[]{}#%<>") {
		return false
	}

	// Skip common non-keyword terms
	for _, term := range skip {
		if s == term {
			return false
		}
	}

	// Accept numeric values (for font-weight, z-index, etc.)
	if regexp.MustCompile(`^\d+$`).MatchString(s) {
		return true
	}

	return true
}

// difference returns elements in a that are not in b.
func difference(a, b []string) []string {
	mb := make(map[string]bool, len(b))
	for _, x := range b {
		mb[x] = true
	}
	
	var diff []string
	for _, x := range a {
		if !mb[x] {
			diff = append(diff, x)
		}
	}
	return diff
}

// printResults prints validation results to stdout.
func (v *Validator) printResults(results []ValidationResult) {
	fmt.Println("=== MDN Validation Results ===")
	fmt.Println()

	for _, result := range results {
		fmt.Printf("Property: %s\n", result.Property)
		fmt.Printf("  Exists in MDN: %t\n", result.Exists)
		
		if result.Exists {
			fmt.Printf("  Status: %s\n", result.Status)
			fmt.Printf("  Spec keywords (%d): %v\n", len(result.SpecKeywords), result.SpecKeywords)
			fmt.Printf("  MDN keywords (%d): %v\n", len(result.MDNKeywords), result.MDNKeywords)
			
			if len(result.MissingKeywords) > 0 {
				fmt.Printf("  ⚠️  Missing keywords: %v\n", result.MissingKeywords)
			}
			
			if len(result.ExtraKeywords) > 0 {
				fmt.Printf("  ⚠️  Extra keywords: %v\n", result.ExtraKeywords)
			}
			
			if len(result.MissingKeywords) == 0 && len(result.ExtraKeywords) == 0 {
				fmt.Printf("  ✅ Keywords match MDN\n")
			}
		} else {
			fmt.Printf("  ❌ Property not found in MDN data\n")
		}
		
		fmt.Println()
	}
}

// updateSpecWithMDN creates an updated spec incorporating MDN data.
func (v *Validator) updateSpecWithMDN() Spec {
	updated := Spec{
		Version:    v.spec.Version,
		Properties: make([]PropertySpec, 0, len(v.spec.Properties)),
	}

	for _, prop := range v.spec.Properties {
		updatedProp := prop

		if mdnProp, exists := v.mdnProperties[prop.Name]; exists {
			// Update keywords with MDN data
			mdnKeywords := v.extractKeywords(mdnProp.Syntax)
			if len(mdnKeywords) > 0 {
				updatedProp.Keywords = mdnKeywords
			}
			
			// Update syntax
			updatedProp.Syntax = mdnProp.Syntax
			
			// Update status
			updatedProp.Status = mdnProp.Status
		}

		updated.Properties = append(updated.Properties, updatedProp)
	}

	return updated
}

// writeSpec writes a spec to a JSON file.
func (v *Validator) writeSpec(spec Spec, path string) error {
	data, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal spec: %w", err)
	}

	return os.WriteFile(path, data, 0644)
}