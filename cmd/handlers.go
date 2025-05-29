package cmd

import (
	"fmt"
	"strconv"
)

func HandlePlantSeeds(state *GameState, inputChan <-chan string) string {
	seeds := GetAvailableSeeds()

	fmt.Println("\nAvailable Seeds:")
	for i, seed := range seeds {
		fmt.Printf("%d. %s (Cost: %d 🪙) %s\n", i+1, seed.Name(), seed.Cost(), seed.Emoji())
	}

	seedInput := readInput(inputChan, "\nChoose a seed to plant: ")
	seedChoice, err := strconv.Atoi(seedInput)
	if err != nil || seedChoice < 1 || seedChoice > len(seeds) {
		return "Invalid seed choice."
	}
	selectedSeed := seeds[seedChoice-1]

	qtyInput := readInput(inputChan, fmt.Sprintf("How many %s do you want to plant? ", selectedSeed.Name()))
	quantity, err := strconv.Atoi(qtyInput)
	if err != nil || quantity <= 0 {
		return "Invalid quantity."
	}

	return state.PlantSeeds(selectedSeed, quantity, selectedSeed.Cost())
}

func HandleHarvest(state *GameState) string {
	harvested := state.Harvest()
	if harvested == 0 {
		return "No mature crops to harvest."
	}
	return fmt.Sprintf("Harvested %d crop(s)!", harvested)
}

func HandleSellAll(state *GameState) string {
	sold, coins := state.SellAll()
	if sold == 0 {
		return "Nothing to sell."
	}
	return fmt.Sprintf("Sold %d items for %d 🪙!", sold, coins)
}

func readInput(inputChan <-chan string, prompt string) string {
	fmt.Print(prompt)
	return <-inputChan
}
