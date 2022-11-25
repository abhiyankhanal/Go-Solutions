package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name       string
	leftSpoon  int
	rightSpoon int
}

var hunger int = 3
var eatTime = 3 * time.Second
var sleepTime = 2 * time.Second
var thinkTime = 2 * time.Second
var philosophers = []Philosopher{
	{name: "abhiyan", leftSpoon: 4, rightSpoon: 0},
	{name: "anish", leftSpoon: 0, rightSpoon: 1},
	{name: "akash", leftSpoon: 1, rightSpoon: 2},
	{name: "bishal", leftSpoon: 2, rightSpoon: 3},
	{name: "bishwo", leftSpoon: 3, rightSpoon: 4},
}

func main() {
	fmt.Println("Dining Philosopher problem")
	fmt.Println("-------------------------------")
	time.Sleep(sleepTime)

	dine()
	fmt.Println("Dining Philosopher problem ending")
	time.Sleep(sleepTime)
	fmt.Println("Dining Philosopher problem ended")

}

func dine() {
	// wg is the WaitGroup that keeps track of how many philosophers are still at the table. When
	// it reaches zero, everyone is finished eating and has left. We add 5 (the number of philosophers) to this
	// wait group.
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// We want everyone to be seated before they start eating, so create a WaitGroup for that, and set it to 5.
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// Spoons is a map of all 5 Spoons. Spoons are assigned using the fields leftSpoon and rightSpoon in the Philosopher
	// type. Each Spoon, then, can be found using the index (an integer), and each Spoon has a unique mutex.
	var spoons = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		spoons[i] = &sync.Mutex{}
	}

	// Start the meal by iterating through our slice of Philosophers.
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, spoons, seated)
	}

	// Wait for the philosophers to finish. This blocks until the wait group is 0.
	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, spoons map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name)

	// Decrement the seated WaitGroup by one.
	seated.Done()

	// Wait until everyone is seated.
	seated.Wait()

	// Have this philosopher eatTime and thinkTime "hunger" times (3).
	for i := hunger; i > 0; i-- {
		// Get a lock on the left and right spoons. We have to choose the lower numbered Spoon first in order
		// to avoid a logical race condition, which is not detected by the -race flag in tests; if we don't do this,
		// we have the potential for a deadlock, since two philosophers will wait endlessly for the same Spoon.
		// Note that the goroutine will block (pause) until it gets a lock on both the right and left Spoons.
		if philosopher.leftSpoon > philosopher.rightSpoon {
			spoons[philosopher.rightSpoon].Lock()
			fmt.Printf("\t%s takes the right Spoon.\n", philosopher.name)
			spoons[philosopher.leftSpoon].Lock()
			fmt.Printf("\t%s takes the left Spoon.\n", philosopher.name)
		} else {
			spoons[philosopher.leftSpoon].Lock()
			fmt.Printf("\t%s takes the left Spoon.\n", philosopher.name)
			spoons[philosopher.rightSpoon].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		// By the time we get to this line, the philosopher has a lock (mutex) on both spoons.
		fmt.Printf("\t%s has both spoons and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		// The philosopher starts to think, but does not drop the spoons yet.
		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		// Unlock the mutexes for both spoons.
		spoons[philosopher.leftSpoon].Unlock()
		spoons[philosopher.rightSpoon].Unlock()

		fmt.Printf("\t%s put down the spoons.\n", philosopher.name)
	}

	// The philosopher has finished eating, so print out a message.
	fmt.Println(philosopher.name, "is satisified.")
	fmt.Println(philosopher.name, "left the table.")

}
