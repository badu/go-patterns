package abstractfactory

// iShort represents a short interface
// defines all the methods needed for short
type iShort interface {
	SetSize(int)
	Size() int
	SetColor(string)
	Color() string
}
