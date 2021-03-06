package gherkin

// Language represents the language
type Language string

// Translation represents the Gherkin syntax keywords in a single language.
type Translation struct {
	// Language specific term representing a feature.
	Attack string

	// Language specific term representing the feature background.
	Background string

	// Language specific term representing a scenario.
	Scenario string

	// Language specific term representing a scenario outline.
	Outline string

	// Language specific term representing the "And" step.
	And string

	// Language specific term representing the "Given" step.
	Given string

	// Language specific term representing the "When" step.
	When string

	// Language specific term representing the "Then" step.
	Then string

	// Language specific term representing a scenario outline prefix.
	Examples string
}

const (
	// LangEN is the language code for English translations.
	LangEN = Language("en")
)

var (
	// Translations contains internationalized translations of the Gherkin
	// syntax keywords in a variety of supported languages.
	Translations = map[Language]Translation{
		LangEN: {
			Attack:     "Attack",
			Background: "Background",
			Scenario:   "Scenario",
			Outline:    "Scenario Outline",
			And:        "And",
			Given:      "Given",
			When:       "When",
			Then:       "Then",
			Examples:   "Examples",
		},
	}
)
