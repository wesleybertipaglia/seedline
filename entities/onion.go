package entities

import "seedline/interfaces"

type OnionSeed struct{}

func (o *OnionSeed) Name() string {
	return "Onion"
}

func (o *OnionSeed) Emoji() string {
	return "🧅"
}

func (o *OnionSeed) Grow() interfaces.Plant {
	return &OnionPlant{}
}

func (o *OnionSeed) Cost() int {
	return 8
}

type OnionPlant struct {
	growth int
}

func (o *OnionPlant) Name() string {
	return "Onion"
}

func (o *OnionPlant) Emoji() string {
	return "🧅"
}

func (o *OnionPlant) TurnsToGrow() int {
	return 4
}

func (o *OnionPlant) AdvanceGrowth() {
	o.growth++
}

func (o *OnionPlant) IsMature() bool {
	return o.growth >= o.TurnsToGrow()
}

func (o *OnionPlant) SellPrice() int {
	return 8
}
