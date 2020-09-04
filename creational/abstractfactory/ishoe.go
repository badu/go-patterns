package abstractfactory

// iShoe represents a shoe interface
// defines all the methods needed for shoe
type iShoe interface {
	SetSize(int)
	Size() int
	SetColor(string)
	Color() string
}
