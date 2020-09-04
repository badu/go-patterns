package factory

// iGun interface represents the gun
// Has all the methods needed to create a gun type
// This is implemented by the gun type
type iGun interface {
	SetName(string)
	Name() string
	SetBullets(int)
	Bullets() int
}
