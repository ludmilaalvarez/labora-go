package main

import (
	"fmt"
	"math/rand"
	"time"
	"Peleadores/fighters"


)

var numero int

func main() {
	var police fighters.Police = fighters.Police{
		BaseFighter: fighters.BaseFighter{
			Life: 100,
		},
		Armour: 50,
	}
	var criminal fighters.Criminal = fighters.Criminal{
		BaseFighter: fighters.BaseFighter{
			Life: 100,
		},
	}
	var paladin fighters.Paladin= fighters.Paladin{
		BaseFighter: fighters.BaseFighter{
			Life: 200,
		},
	} 

	var contenders []fighters.Contender = make([]fighters.Contender, 3)

	//randomValueBetweenOneAndZero := rand.Intn(3)
	contenders[0] = &police
	contenders[1] = &criminal
	contenders[2] = &paladin


	var areBothAlive = police.IsAlive() && criminal.IsAlive() && paladin.IsAlive()
	for areBothAlive {
		if contenders[0].IsAlive(){
			intensity := contenders[0].ThrowAttack()
			numero=RandomGolpe(0)
			fmt.Println(contenders[0].GetName(), " tira golpe con intensidad = ", intensity, " a ", contenders[numero].GetName())
			contenders[numero].RecieveAttack(&intensity)

		}

		if contenders[1].IsAlive() {
			intensity := contenders[1].ThrowAttack()
			numero=RandomGolpe(1)
			fmt.Println(contenders[1].GetName(), " tira golpe con intensidad =", intensity, " a ", contenders[numero].GetName())
			contenders[numero].RecieveAttack(&intensity)
		}

		if contenders[2].IsAlive(){
			intensity:= contenders[2].ThrowAttack()
			numero=RandomGolpe(2)
			fmt.Println(contenders[2].GetName(), "tira golpe con intensidad=", intensity, " a ", contenders[numero].GetName())
			
			contenders[numero].RecieveAttack(&intensity)

		}


		fmt.Printf("PoliceLife=%d, CriminalLife=%d, PaladinLife=%d\n", police.Life, criminal.Life, paladin.Life)
		areBothAlive = police.IsAlive() && criminal.IsAlive() && paladin.IsAlive()
		time.Sleep(3 * time.Second)
	}
}


func RandomGolpe(numeros int) int{
	num:=numeros
	numero=rand.Intn(3)

	if numero == num{
		RandomGolpe(num)
	}

	return numero
}
