package factory

// scarl represents the ScarL Gun type
type scarl struct {
	gun
}

// Scarl return a scarl gun type
func (s *scarl) Scarl() iGun {
	s.SetName("ScarL")
	s.SetBullets(30)

	return s
}
