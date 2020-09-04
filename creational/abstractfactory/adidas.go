package abstractfactory

// adidas represents adidas brand
type adidas struct{}

// GetShoe returns an adidas shoe instance
func (a *adidas) GetShoe() iShoe {

	return &adidasShoe{}
}

// GetShort return an adidas short instance
func (a *adidas) GetShort() iShort {

	return &adidasshort{}
}
