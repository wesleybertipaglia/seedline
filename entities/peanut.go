package entities

import "seedline/interfaces"

type PeanutSeed struct{}

func (p *PeanutSeed) Name() string {
	return "Peanut"
}

func (p *PeanutSeed) Emoji() string {
	return "🥜"
}

func (p *PeanutSeed) Grow() interfaces.Plant {
	return &PeanutPlant{}
}

func (p *PeanutSeed) Cost() int {
	return 7
}

type PeanutPlant struct {
	growth int
}

func (p *PeanutPlant) Name() string {
	return "Peanut"
}

func (p *PeanutPlant) Emoji() string {
	return "🥜"
}

func (p *PeanutPlant) TurnsToGrow() int {
	return 3
}

func (p *PeanutPlant) AdvanceGrowth() {
	p.growth++
}

func (p *PeanutPlant) IsMature() bool {
	return p.growth >= p.TurnsToGrow()
}

func (p *PeanutPlant) SellPrice() int {
	return 6
}
