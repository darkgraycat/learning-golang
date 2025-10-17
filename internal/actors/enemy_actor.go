package actors

import "fmt"

type EnemyActor struct{}

func (a *EnemyActor) DoTurn() string {
	fmt.Println("Enemy Turn")

	return "E"
}
