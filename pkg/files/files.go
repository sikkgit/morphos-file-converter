package files

// File interface is the main interface of the package,
// that defines what a file is in this context.
// It's moslty responsible to say other entitites what formats it can be converted to
// and provides a method to convert the current file given a target format, if supported.
type File interface {
	SupportedFormats() map[string][]string
	ConvertTo(string, []byte) ([]byte, error)
}
