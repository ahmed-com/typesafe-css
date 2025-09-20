package tailwind

import (
	"fmt"
	"sync"

	"github.com/ahmed-com/typesafe-css/css"
)

// UtilityManager manages utility classes and ensures deduplication.
type UtilityManager struct {
	theme    *Theme
	rules    map[string]css.Rule // Maps class name to CSS rule
	mu       sync.RWMutex         // Protects concurrent access
}

// NewUtilityManager creates a new utility manager with the given theme.
func NewUtilityManager(theme *Theme) *UtilityManager {
	if theme == nil {
		theme = DefaultTheme()
	}
	
	return &UtilityManager{
		theme: theme,
		rules: make(map[string]css.Rule),
	}
}

// GetOrCreateRule gets an existing rule or creates a new one if it doesn't exist.
// This ensures each utility class is only defined once.
func (um *UtilityManager) GetOrCreateRule(className string, createFn func() css.Rule) css.Rule {
	um.mu.Lock()
	defer um.mu.Unlock()
	
	if rule, exists := um.rules[className]; exists {
		return rule
	}
	
	rule := createFn()
	um.rules[className] = rule
	return rule
}

// GetRules returns all currently registered utility rules.
func (um *UtilityManager) GetRules() []css.Rule {
	um.mu.RLock()
	defer um.mu.RUnlock()
	
	rules := make([]css.Rule, 0, len(um.rules))
	for _, rule := range um.rules {
		rules = append(rules, rule)
	}
	return rules
}

// AddToStylesheet adds all utility rules to the given stylesheet.
func (um *UtilityManager) AddToStylesheet(stylesheet *css.Stylesheet) {
	rules := um.GetRules()
	for _, rule := range rules {
		stylesheet.Add(rule)
	}
}

// Theme returns the current theme.
func (um *UtilityManager) Theme() *Theme {
	return um.theme
}

// UpdateTheme updates the theme and clears the rule cache.
func (um *UtilityManager) UpdateTheme(theme *Theme) {
	um.mu.Lock()
	defer um.mu.Unlock()
	
	um.theme = theme
	um.rules = make(map[string]css.Rule) // Clear cache since theme changed
}

// ClassName generates a utility class name from a prefix and value key.
func ClassName(prefix, value string) string {
	if value == "" {
		return prefix
	}
	return fmt.Sprintf("%s-%s", prefix, value)
}

// Default global utility manager for convenience
var defaultManager = NewUtilityManager(nil)

// SetDefaultTheme sets the theme for the default utility manager.
func SetDefaultTheme(theme *Theme) {
	defaultManager.UpdateTheme(theme)
}

// GetDefaultManager returns the default utility manager.
func GetDefaultManager() *UtilityManager {
	return defaultManager
}