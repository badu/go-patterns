package abstractfactory

// short represents a short type
type short struct {
	color string
	size  int
}

// SetSize sets the short size
func (s *short) SetSize(size int) {
	s.size = size
}

// Size returns the short size
func (s *short) Size() int {
	return s.size
}

// SetColor sets the short color
func (s *short) SetColor(color string) {
	s.color = color
}

// Color returns the short color
func (s *short) Color() string {
	return s.color
}
