package interfaces

type Seed interface {
	Name() string
	Emoji() string
	Grow() Plant
	Cost() int
}
