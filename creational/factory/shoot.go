package factory

import "fmt"

func Shoot(gun iGun) {
	fmt.Printf("Shooting %s\n", gun.Name())
	gun.SetBullets(gun.Bullets() - 1)
	fmt.Printf("Bullets left: %d\n", gun.Bullets())
}
