package entities

import "seedline/interfaces"

type BeanSeed struct{}

func (b *BeanSeed) Name() string {
	return "Bean"
}

func (b *BeanSeed) Emoji() string {
	return "🫘"
}

func (b *BeanSeed) Grow() interfaces.Plant {
	return &BeanPlant{}
}

func (b *BeanSeed) Cost() int {
	return 10
}

type BeanPlant struct {
	growth int
}

func (b *BeanPlant) Name() string {
	return "Bean"
}

func (b *BeanPlant) Emoji() string {
	return "🫘"
}

func (b *BeanPlant) TurnsToGrow() int {
	return 2
}

func (b *BeanPlant) AdvanceGrowth() {
	b.growth++
}

func (b *BeanPlant) IsMature() bool {
	return b.growth >= b.TurnsToGrow()
}

func (b *BeanPlant) SellPrice() int {
	return 30
}
