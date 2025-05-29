package entities

import "seedline/interfaces"

type CarrotSeed struct{}

func (c *CarrotSeed) Name() string {
	return "Carrot"
}

func (c *CarrotSeed) Emoji() string {
	return "🥕"
}

func (c *CarrotSeed) Grow() interfaces.Plant {
	return &CarrotPlant{}
}

func (c *CarrotSeed) Cost() int {
	return 7
}

type CarrotPlant struct {
	growth int
}

func (c *CarrotPlant) Name() string {
	return "Carrot"
}

func (c *CarrotPlant) Emoji() string {
	return "🥕"
}

func (c *CarrotPlant) TurnsToGrow() int {
	return 2
}

func (c *CarrotPlant) AdvanceGrowth() {
	c.growth++
}

func (c *CarrotPlant) IsMature() bool {
	return c.growth >= c.TurnsToGrow()
}

func (c *CarrotPlant) SellPrice() int {
	return 6
}
