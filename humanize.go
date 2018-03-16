// Package humanize provides methods for formatting and parsing values in human readable form.
package humanize

import (
	"errors"
	"fmt"
	"regexp"
)

// List all the language providers here.
var languages = map[string]languageProvider{
	"pl": lang_pl,
	"en": lang_en,
}

// languageProvider is a struct defining all the needed language elements.
type languageProvider struct {
	times times
}

// Humanizer is the main struct that provides the public methods.
type Humanizer struct {
	provider    languageProvider
	timeInputRe *regexp.Regexp
	metricInputRe *regexp.Regexp
}

// New creates a new humanizer for a given language.
func New(language string) (*Humanizer, error) {
	if provider, exists := languages[language]; exists {
		humanizer := &Humanizer{
			provider: provider,
		}
		humanizer.buildTimeInputRe()
		humanizer.buildMetricInputRe()
		return humanizer, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Language not supported: %s", language))
	}
}
