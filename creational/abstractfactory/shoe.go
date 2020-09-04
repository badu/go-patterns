package abstractfactory

// shoe represents a shoe type
type shoe struct {
	size  int
	color string
}

// SetSize sets the shoe size
func (s *shoe) SetSize(size int) {
	s.size = size
}

// Size returns the shoe size
func (s *shoe) Size() int {
	return s.size
}

// SetColor sets the shoe color
func (s *shoe) SetColor(color string) {
	s.color = color
}

// Color returns the shoe color
func (s *shoe) Color() string {
	return s.color
}
