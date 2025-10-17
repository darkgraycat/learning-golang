package cmd

import (
	"fmt"
	"my_project/internal/actors"
	"my_project/internal/utils"
	"sync"
)

func Execute() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Enter filename: ")
	// line, _ := reader.ReadString('\n')

	// fmt.Printf("Entered line %s\n", line)

	var bytes []byte

	bytes = append(bytes, "world"...)
	fmt.Printf(string(bytes))
	return

	var hero actors.Actor
	var enemy actors.Actor

	hero = &actors.HeroActor{}
	enemy = &actors.EnemyActor{}

	var history []string = make([]string, 0)
	var wg sync.WaitGroup

	totalRuns := 8
	for range totalRuns {
		wg.Add(1)
		// utils.DoStuff(&wg, func() { hero.DoTurn() })
		// utils.DoStuff(&wg, func() { enemy.DoTurn() })
		utils.DoStuff(&wg, func() {
			actors.HandleActorTurn_Mutation(hero, &history)
			actors.HandleActorTurn_Mutation(enemy, &history)

			history = actors.HandleActorTurn_Pure(hero, history)
			history = actors.HandleActorTurn_Pure(enemy, history)
		})
	}
	wg.Wait()

	fmt.Println(history)
}
