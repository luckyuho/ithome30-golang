package main

import (
	"it/day18/people"
	"it/day18/say"
)

func main() {
	say.SayHi()

	people := people.People
	for i := range people {
		say.SayWords(people[i], "hello")
	}
}
