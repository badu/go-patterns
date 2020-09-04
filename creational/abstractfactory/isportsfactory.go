package abstractfactory

import (
	"errors"
)

const factoryNotFoundErr = "factory not found"

// iSportsFactory represents a sports factory
// Defines all methods need to make a sports kit
type iSportsFactory interface {
	GetShoe() iShoe
	GetShort() iShort
}

// GetSportsFactory returns an instance of the specified sports factory if found
// Otherwise returns an error
func GetSportsFactory(factory string) (iSportsFactory, error) {

	switch factory {
	case "adidas":
		return &adidas{}, nil
	case "nike":
		return &nike{}, nil
	}
	return nil, errors.New(factoryNotFoundErr)
}
