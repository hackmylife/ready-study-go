//go:build ignore

package koan

type FieldError struct{ Field, Reason string }

func (e *FieldError) Error() string { return e.Field + ": " + e.Reason }
func Validate(email string) error {
	if email == "" {
		return &FieldError{Field: "email", Reason: "required"}
	}
	return nil
}
