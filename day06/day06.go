package main

import (
	"fmt"
)

type ZooAnimal interface {
	Feeder
	Rester
	Placer
	// Feed() int
	// Rest() int
	// Place() string
}

type Feeder interface {
	Feed() int
}

type Rester interface {
	Rest() int
}

type Placer interface {
	Place() string
}

type Lion struct {
	AnimalInfo Animal
}

type Animal struct {
	Amount    int
	FeedDaily int
	RestDaily int
	Place     string
}

func (lion *Lion) Feed() int {
	return lion.AnimalInfo.Amount * lion.AnimalInfo.FeedDaily
}

func (lion *Lion) Rest() int {
	return lion.AnimalInfo.Amount * lion.AnimalInfo.RestDaily
}

func (lion *Lion) Place() string {
	return lion.AnimalInfo.Place
}

func ShowFeedYear(animal ZooAnimal) int {
	feed_daily := animal.Feed()
	var ratio int
	switch animal.Place() {
	case "熱帶":
		ratio = 1
	case "寒帶":
		ratio = 2
	}
	return feed_daily * ratio * 365

}

func ShowRestYear(animal ZooAnimal) int {
	rest_daily := animal.Rest()
	return rest_daily * 365
}

func ShowTotalYear(animal ZooAnimal) {
	feed_year := ShowFeedYear(animal)
	rest_year := ShowRestYear(animal)
	fmt.Printf("每年花費為 %d 元 (此為毫無根據的數字) \n", feed_year+rest_year)
}

func main() {
	// lion := &Lion{AnimalInfo: Animal{Amount: 2, FeedDaily: 500, RestDaily: 50, Place: "熱帶"}}
	// ShowTotalYear(lion, "獅子")
	// penguin := &Penguin{AnimalInfo: Animal{Amount: 10, FeedDaily: 100, RestDaily: 100, Place: "寒帶"}}
	// ShowTotalYear(penguin, "企鵝")

	animals := [...]ZooAnimal{
		&Lion{AnimalInfo: Animal{Amount: 2, FeedDaily: 500, RestDaily: 50, Place: "熱帶"}},
		&Penguin{AnimalInfo: Animal{Amount: 10, FeedDaily: 100, RestDaily: 100, Place: "寒帶"}},
	}

	for _, animal := range animals {
		switch animal.(type) {
		case *Lion:
			fmt.Printf("獅子 ")
		case *Penguin:
			fmt.Printf("企鵝 ")
		default:
			fmt.Println("是誰偷偷混進來啊～!!")
		}
		ShowTotalYear(animal)
	}
}

func (penguin *Penguin) Feed() int {
	return penguin.AnimalInfo.Amount * penguin.AnimalInfo.FeedDaily
}

func (penguin *Penguin) Rest() int {
	return penguin.AnimalInfo.Amount * penguin.AnimalInfo.RestDaily
}

func (penguin *Penguin) Place() string {
	return penguin.AnimalInfo.Place
}

type Penguin struct {
	AnimalInfo Animal
}
