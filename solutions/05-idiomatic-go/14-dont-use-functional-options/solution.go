//go:build ignore

package koan

type Formatter struct{ Prefix string }

func NewFormatter(prefix string) *Formatter {
	return &Formatter{Prefix: prefix}
}
func (f *Formatter) Format(s string) string { return f.Prefix + s }
