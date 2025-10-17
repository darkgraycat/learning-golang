package actors

import (
	"math/rand/v2"
	"time"
)

type Actor interface {
	DoTurn() string
}

func HandleActorTurn_Mutation(a Actor, history *[]string) {
	action := a.DoTurn()

	time.Sleep(time.Second * time.Duration(rand.Float32()))

	*history = append(*history, action)
}

func HandleActorTurn_Pure(a Actor, history []string) []string {
	action := a.DoTurn()

	time.Sleep(time.Second * time.Duration(rand.Float32()))

	return append(history, action)
}
