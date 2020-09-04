package abstractfactory

// nike represents nike brand
type nike struct{}

// GetShoe returns an nike shoe instance
func (n *nike) GetShoe() iShoe {

	return &nikeShoe{}
}

// GetShort returns an nike short instance
func (n *nike) GetShort() iShort {

	return &nikeshort{}
}
