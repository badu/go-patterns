package factory

// akm represents the AKM Gun type
type akm struct {
	gun
}

// Akm returns a AKM gun type
func (a *akm) Akm() iGun {
	a.SetName("AKM")
	a.SetBullets(30)

	return a
}
