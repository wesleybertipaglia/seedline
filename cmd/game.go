package cmd

import (
	"fmt"
	"seedline/interfaces"
)

type GameState struct {
	Coins     int
	Inventory map[string]int
	Farm      []interfaces.Plant
}

func NewGameState() *GameState {
	return &GameState{
		Coins:     20,
		Inventory: make(map[string]int),
		Farm:      []interfaces.Plant{},
	}
}

func (g *GameState) PlantSeed(seed interfaces.Seed) {
	plant := seed.Grow()
	g.Farm = append(g.Farm, plant)
	fmt.Printf("Planted %s %s\n", plant.Name(), plant.Emoji())
}

func (g *GameState) PlantSeeds(seed interfaces.Seed, count int, costPerSeed int) string {
	maxCapacity := 6 * 12
	if len(g.Farm)+count > maxCapacity {
		return "Not enough space on the farm."
	}
	totalCost := costPerSeed * count
	if g.Coins < totalCost {
		return fmt.Sprintf("Not enough coins. You need %d 🪙 to plant %d seeds.", totalCost, count)
	}

	for i := 0; i < count; i++ {
		g.Farm = append(g.Farm, seed.Grow())
	}
	g.Coins -= totalCost
	return fmt.Sprintf("Planted %d %s 🌱", count, seed.Name())
}

func (g *GameState) Harvest() int {
	var remaining []interfaces.Plant
	count := 0

	for _, plant := range g.Farm {
		if plant.IsMature() {
			g.Inventory[plant.Name()]++
			count++
		} else {
			remaining = append(remaining, plant)
		}
	}
	g.Farm = remaining
	return count
}

func (g *GameState) AdvanceTurn() {
	for _, plant := range g.Farm {
		plant.AdvanceGrowth()
	}
}

func (g *GameState) SellAll() (itemsSold int, totalEarned int) {
	if len(g.Inventory) == 0 {
		return 0, 0
	}

	for cropName, count := range g.Inventory {
		seed, ok := SeedsByName[cropName]
		var price int
		if ok {
			price = seed.Grow().SellPrice()
		} else {
			price = 5
		}

		earned := price * count
		totalEarned += earned
		itemsSold += count
	}

	g.Coins += totalEarned
	g.Inventory = make(map[string]int)
	return itemsSold, totalEarned
}

func (g *GameState) BasketSize() int {
	total := 0
	for _, count := range g.Inventory {
		total += count
	}
	return total
}
