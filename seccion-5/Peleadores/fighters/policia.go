package fighters

import "math/rand"

type Police struct {
	BaseFighter
	Armour int // 0..50
}

func (p Police) ThrowAttack() int {
	return rand.Intn(10)
}

func (p *Police) RecieveAttack(intensity *int) {
	if p.Armour > 0 {
		diff := (p.Armour - *intensity)
		hasArmour := diff > 0
		if hasArmour {
			p.Armour -= *intensity
			*intensity = 0
		} else {
			p.Armour = 0
			*intensity = -diff // intensity -= p.Armour
		}
	}
	p.Life -= *intensity
	if p.Life <0{
		p.Life=0
	}
	
}

func (p Police) GetName() string {
	return "Policia"
}