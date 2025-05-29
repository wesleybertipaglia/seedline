package cmd

import (
	"fmt"
	"seedline/interfaces"
)

func RenderFarm(farm []interfaces.Plant, cols, rows int) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			idx := r*cols + c
			if idx < len(farm) {
				plant := farm[idx]
				if plant.IsMature() {
					fmt.Print(plant.Emoji())
				} else {
					fmt.Print("🌱")
				}
			} else {
				fmt.Print("🟫")
			}
		}
		fmt.Println()
	}
}
