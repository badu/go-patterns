package factory

// gun represents the gun type
// Implements the iGun interface
type gun struct {
	name    string
	bullets int
}

// SetName sets the gun name
func (g *gun) SetName(name string) {
	g.name = name
}

// Name returns the gun name
func (g *gun) Name() string {
	return g.name
}

// SetBullets sets the bullets in the gun
func (g *gun) SetBullets(bullets int) {
	g.bullets = bullets
}

// Bullets return the no. of bullets in the gun
func (g *gun) Bullets() int {
	return g.bullets
}
