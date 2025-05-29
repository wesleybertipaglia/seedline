package entities

import "seedline/interfaces"

type PotatoSeed struct{}

func (p *PotatoSeed) Name() string {
	return "Potato"
}

func (p *PotatoSeed) Emoji() string {
	return "🥔"
}

func (p *PotatoSeed) Grow() interfaces.Plant {
	return &PotatoPlant{}
}

func (p *PotatoSeed) Cost() int {
	return 6
}

type PotatoPlant struct {
	growth int
}

func (p *PotatoPlant) Name() string {
	return "Potato"
}

func (p *PotatoPlant) Emoji() string {
	return "🥔"
}

func (p *PotatoPlant) TurnsToGrow() int {
	return 4
}

func (p *PotatoPlant) AdvanceGrowth() {
	p.growth++
}

func (p *PotatoPlant) IsMature() bool {
	return p.growth >= p.TurnsToGrow()
}

func (p *PotatoPlant) SellPrice() int {
	return 7
}
