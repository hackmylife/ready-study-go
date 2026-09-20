package koan

type FieldError struct{ Field, Reason string }

func (e *FieldError) Error() string { return "" }
func Validate(email string) error { // TODO: フィールドのエラーを返す
	return nil
}
