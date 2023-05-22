package fighters

import "math/rand"

type Criminal struct {
	BaseFighter
}

func (p *Criminal) ThrowAttack() int {
	return rand.Intn(10)
}

func (p *Criminal) RecieveAttack(intensity *int) {
	p.Life -= *intensity
	if p.Life <0{
		p.Life=0
	}
	
}

func (p *Criminal) GetName() string {
	return "Criminal"
}