package forms

type Errors map[string][]string

// Add an error message for a given form field
func (e Errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message
func (e Errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
