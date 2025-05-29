package entities

import "seedline/interfaces"

type TomatoSeed struct{}

func (t *TomatoSeed) Name() string {
	return "Tomato"
}

func (t *TomatoSeed) Emoji() string {
	return "🍅"
}

func (t *TomatoSeed) Grow() interfaces.Plant {
	return &TomatoPlant{}
}

func (t *TomatoSeed) Cost() int {
	return 4
}

type TomatoPlant struct {
	growth int
}

func (t *TomatoPlant) Name() string {
	return "Tomato"
}

func (t *TomatoPlant) Emoji() string {
	return "🍅"
}

func (t *TomatoPlant) TurnsToGrow() int {
	return 2
}

func (t *TomatoPlant) AdvanceGrowth() {
	t.growth++
}

func (t *TomatoPlant) IsMature() bool {
	return t.growth >= t.TurnsToGrow()
}

func (t *TomatoPlant) SellPrice() int {
	return 8
}
