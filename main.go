package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func main() {
	profile := MakeProfile("Sasuke")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp := PowerUp(profile, 2)
	fmt.Println("Name : ", powerUp.Name)
	fmt.Println("Health : ", powerUp.Health)
	fmt.Println("Power : ", powerUp.Power)
	fmt.Println("Exp : ", powerUp.Exp)
}

func MakeProfile(character string) Profile {
	var profile = Profile{}
	if character == "Sasuke" {
		profile.Name = "Sasuke"
		profile.Health = 200
		profile.Power = 100
		profile.Exp = 0
		return profile
	}
	if character == "Goku" {
		profile.Name = "Goku"
		profile.Health = 400
		profile.Power = 300
		profile.Exp = 100
		return profile
	}
	if character == "Naruto" {
		profile.Name = "Naruto"
		profile.Health = 150
		profile.Power = 200
		profile.Exp = 50
		return profile
	}
	return profile
}

func PowerUp(prosedur Profile, multiplier int) Profile {
	prosedur.Health = prosedur.Health + (prosedur.Health * multiplier)
	prosedur.Power = prosedur.Power + (prosedur.Power * multiplier)
	prosedur.Exp = prosedur.Exp + (prosedur.Exp * multiplier)
	return prosedur
}
