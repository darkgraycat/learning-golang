package actors

import "fmt"

type HeroActor struct{}

func (a *HeroActor) DoTurn() string {
	fmt.Println("Hero Turn")

	return "H"
}
