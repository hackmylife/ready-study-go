package koan

type Formatter struct{ Prefix string }
type Option func(*Formatter)

func WithPrefix(prefix string) Option { return func(f *Formatter) { f.Prefix = prefix } }
func NewFormatter(options ...Option) *Formatter { // TODO: 単一引数へ簡素化する
	f := &Formatter{}
	for _, o := range options {
		o(f)
	}
	return f
}
func (f *Formatter) Format(s string) string { return f.Prefix + s }
