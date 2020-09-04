package factory

import "errors"

const invalidGunErr = "invalid gun type"

// NewGun is factory for guns
// Returns the gun of requested type if found
// Otherwise returns an error
func NewGun(gunType string) (iGun, error) {

	switch gunType {
	case "akm":
		gun := &akm{}

		return gun.Akm(), nil
	case "scarl":
		gun := &scarl{}

		return gun.Scarl(), nil
	}
	return nil, errors.New(invalidGunErr)
}
