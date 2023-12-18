package forms

type errors map[string][]string

// Add() method to add error messages for a given field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// a Get() method to get the first message for a given field
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
