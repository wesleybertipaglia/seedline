package interfaces

type Plant interface {
	Name() string
	Emoji() string
	TurnsToGrow() int
	AdvanceGrowth()
	IsMature() bool
	SellPrice() int
}
