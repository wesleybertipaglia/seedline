package entities

import "seedline/interfaces"

type CornSeed struct{}

func (c *CornSeed) Name() string {
	return "Corn"
}

func (c *CornSeed) Emoji() string {
	return "🌽"
}

func (c *CornSeed) Grow() interfaces.Plant {
	return &CornPlant{}
}

func (c *CornSeed) Cost() int {
	return 5
}

type CornPlant struct {
	growth int
}

func (c *CornPlant) Name() string {
	return "Corn"
}

func (c *CornPlant) Emoji() string {
	return "🌽"
}

func (c *CornPlant) TurnsToGrow() int {
	return 3
}

func (c *CornPlant) AdvanceGrowth() {
	c.growth++
}

func (c *CornPlant) IsMature() bool {
	return c.growth >= c.TurnsToGrow()
}

func (c *CornPlant) SellPrice() int {
	return 10
}
