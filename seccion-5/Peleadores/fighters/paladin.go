package fighters

import "math/rand"


type Paladin struct {
	BaseFighter
}

func (p *Paladin) ThrowAttack() int {
	var intensidad float64

	numero:=rand.Intn(10)
	if p.Life==200{
		return numero
	}else {
		intensidad= float64(p.Life)/200.00
		valor:=float64(numero)*intensidad
		return  int(valor)
	}	
}

func (p *Paladin) RecieveAttack(intensity *int) {

	p.Life -= *intensity
	if p.Life <0{
		p.Life=0
	}
}

func (p *Paladin) GetName() string {
	return "Paladin"
}