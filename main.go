package main

import (
	"bufio"
	"fmt"
	"os"
	"seedline/cmd"
	"strconv"
	"strings"
	"time"
)

const (
	farmCols = 12
	farmRows = 6
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printUI(state *cmd.GameState, status string) {
	clearScreen()
	fmt.Printf("🧺 Basket: %d  🪙 Coins: %d\n\n", state.BasketSize(), state.Coins)
	cmd.RenderFarm(state.Farm, farmCols, farmRows)

	fmt.Println("\nFarm Game")
	fmt.Println("1. Plant Seed")
	fmt.Println("2. Harvest")
	fmt.Println("3. Sell All")
	fmt.Println("4. Exit")
	fmt.Printf("\n%s\n", status)
	fmt.Print("\n> Choose an action: ")
}

func main() {
	state := cmd.NewGameState()
	reader := bufio.NewReader(os.Stdin)

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	inputChan := make(chan string)
	go func() {
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				continue
			}
			inputChan <- strings.TrimSpace(input)
		}
	}()

	status := "Welcome to CLI Farm!"
	printUI(state, status)

	for {
		select {
		case <-ticker.C:
			state.AdvanceTurn()
			status = "Time passed. Crops grew!"
			printUI(state, status)

		case input := <-inputChan:
			choice, err := strconv.Atoi(input)
			if err != nil {
				status = "Invalid input. Please enter a number."
				printUI(state, status)
				continue
			}

			switch choice {
			case 1:
				status = cmd.HandlePlantSeeds(state, inputChan)
			case 2:
				status = cmd.HandleHarvest(state)
			case 3:
				status = cmd.HandleSellAll(state)
			case 4:
				fmt.Println("Exiting game...")
				os.Exit(0)
			default:
				status = "Invalid choice. Please try again."
			}
			printUI(state, status)
		}
	}
}
