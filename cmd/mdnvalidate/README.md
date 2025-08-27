# MDN CSS Validation

This tool validates our CSS property specifications against the official Mozilla Developer Network (MDN) CSS data.

## Overview

The MDN validation ensures that our generated CSS types and constants match the authoritative CSS specifications maintained by Mozilla. This helps catch missing properties, incorrect keyword values, and ensures comprehensive coverage of standard CSS features.

## Usage

### Basic Validation

```bash
go run ./cmd/mdnvalidate -mdn /path/to/mdn/data -spec ./cmd/cssgen/spec/spec.json
```

### Update Spec with MDN Data

```bash
go run ./cmd/mdnvalidate -mdn /path/to/mdn/data -spec ./cmd/cssgen/spec/spec.json -update -out ./cmd/cssgen/spec/spec_updated.json
```

### Verbose Output

```bash
go run ./cmd/mdnvalidate -mdn /path/to/mdn/data -spec ./cmd/cssgen/spec/spec.json -v
```

## Getting MDN Data

Clone the official MDN data repository:

```bash
git clone https://github.com/mdn/data.git /tmp/mdn-data
```

Then use `/tmp/mdn-data` as the `-mdn` path.

## Validation Process

1. **Load MDN Data**: Parses `properties.json` and `syntaxes.json` from the MDN data repository
2. **Extract Keywords**: Processes MDN syntax definitions to extract concrete keyword values
3. **Compare Specifications**: Matches our spec against MDN data to identify:
   - Missing keywords that should be added
   - Extra keywords that might be incorrect
   - Properties not found in MDN
4. **Generate Report**: Provides detailed comparison results

## Understanding Results

### ✅ Keywords match MDN
The property's keywords exactly match the MDN specification.

### ⚠️ Missing keywords
Keywords that exist in MDN but are missing from our spec. These should typically be added.

### ⚠️ Extra keywords  
Keywords in our spec that aren't found in MDN. These might be:
- Deprecated values that should be removed
- Alternative syntax forms not captured by the parser
- Custom extensions (review carefully)

### ❌ Property not found in MDN
The property doesn't exist in MDN data. This might indicate:
- Typo in property name
- Very new or experimental property
- Non-standard property

## Limitations

The keyword extraction is based on parsing CSS syntax grammar and has some limitations:

1. **Complex Syntax**: Very complex syntax patterns may not extract all keywords correctly
2. **Numeric Values**: Font-weight numeric values (100-900) are manually maintained in our spec
3. **Context-Sensitive**: Some keywords are only valid in specific contexts not captured by simple parsing

## Integration

The validation is integrated into our test suite:

```bash
# Run all tests including MDN validation
go test ./cmd/cssgen/...

# Run just MDN validation
go test ./cmd/cssgen/ -run TestMDNValidation
```

## Continuous Validation

To ensure ongoing accuracy:

1. Run validation when updating CSS specifications
2. Include in CI/CD pipeline to catch regressions
3. Periodically validate against latest MDN data to catch new CSS features

## Files

- `cmd/mdnvalidate/main.go` - Main validation tool
- `cmd/cssgen/mdn_test.go` - Test integration
- `cssgen/mdn_integration_test.go` - Tests for generated MDN-validated code