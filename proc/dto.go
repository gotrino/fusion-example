package proc

// ProcInfo contains the full set of readable meta data for a proc.
type ProcInfo struct {
	// A longer and more descriptive text in the default language for an end-user about what the expression is about. Use the translations field for language specific values.
	Description string `json:"description"`

	// The unique ID of the expression.
	ID string `json:"id"`

	// A short but arbitrary debug name in the default language for an end-user. Use the translations field for language specific values.
	Name string `json:"name"`

	// An arbitrary set of strings used as tags, e.g. indicating specific topics or templates.
	Tags []string `json:"tags"`

	// Holds translations for the name and summary fields. Keys are arbitrary in the RFC 5646 format.
	Translations ProcInfoTranslation `json:"translations"`
}

// ProcInfoTranslation Holds translations for the name and summary fields. Keys are arbitrary in the RFC 5646 format.
type ProcInfoTranslation struct {
	AdditionalProperties map[string]struct {
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"-"`
}

// ParamInfo describes the input and output specification of a MiEL program. However, this is just a hint from sane programs.
type ParamInfo struct {
	Example struct {
		// An arbitrary response example.
		Request map[string]interface{} `json:"request"`

		// An arbitrary response example.
		Response map[string]interface{} `json:"response"`
	} `json:"example"`
}
