package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	//actor := Actor{
	//	Name:  name,
	//	HP:    hp,
	//	Speed: speed,
	//}

	//actor := Actor{}
	//actor.Name = name
	//actor.HP = hp
	//actor.Speed = speed
	var actor *Actor = new(Actor)
	(*actor).Name = name
	(*actor).HP = hp
	(*actor).Speed = speed

	return actor
}

func main() {
	var actor = NewActor("금토끼", 99, 100)
	fmt.Println(actor.Speed)
	fmt.Println(actor.Name)

}
