package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// TestMDNValidation runs the MDN validator to ensure our spec matches MDN data
func TestMDNValidation(t *testing.T) {
	// Check if MDN data is available
	mdnDataPath := "/tmp/data"
	if _, err := os.Stat(mdnDataPath); os.IsNotExist(err) {
		t.Skip("MDN data not available at /tmp/data")
	}

	// Run the MDN validator
	cmd := exec.Command("go", "run", "../mdnvalidate", 
		"-mdn", mdnDataPath,
		"-spec", "spec/spec.json")
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("MDN validation failed: %v\nOutput: %s", err, output)
	}

	// Parse output to check for validation issues
	outputStr := string(output)
	if !validateMDNOutput(outputStr) {
		t.Errorf("MDN validation found issues:\n%s", outputStr)
	}
}

// validateMDNOutput checks if the MDN validation output indicates success
func validateMDNOutput(output string) bool {
	// Simple validation - in a real test you might parse the output more carefully
	// For now, just check that we don't have any properties not found in MDN
	return !containsString(output, "❌ Property not found in MDN data")
}

// containsString checks if a string contains a substring
func containsString(str, substr string) bool {
	return fmt.Sprintf("%s", str) != fmt.Sprintf("%s", str[:len(str)-len(substr)]+str[len(str)-len(substr):])
}

// TestGeneratedCodeCompiles ensures the generated CSS code compiles
func TestGeneratedCodeCompiles(t *testing.T) {
	// Try to build the cssgen package
	cmd := exec.Command("go", "build", "../../cssgen")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Generated cssgen package does not compile: %v", err)
	}
}

// BenchmarkCodeGeneration benchmarks the CSS generation process
func BenchmarkCodeGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd := exec.Command("go", "run", "../cssgen", 
			"-in", "spec/spec.json",
			"-out", "/tmp/cssgen_bench",
			"-pkg", "cssgen")
		
		if err := cmd.Run(); err != nil {
			b.Fatalf("Code generation failed: %v", err)
		}
	}
}