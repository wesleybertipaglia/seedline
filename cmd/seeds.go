package cmd

import (
	"seedline/entities"
	"seedline/interfaces"
)

func GetAvailableSeeds() []interfaces.Seed {
	return []interfaces.Seed{
		&entities.CornSeed{},
		&entities.BeanSeed{},
		&entities.CarrotSeed{},
		&entities.OnionSeed{},
		&entities.PeanutSeed{},
		&entities.PotatoSeed{},
		&entities.TomatoSeed{},
	}
}

var SeedsByName = map[string]interfaces.Seed{
	"Corn":   &entities.CornSeed{},
	"Bean":   &entities.BeanSeed{},
	"Carrot": &entities.CarrotSeed{},
	"Onion":  &entities.OnionSeed{},
	"Peanut": &entities.PeanutSeed{},
	"Potato": &entities.PotatoSeed{},
	"Tomato": &entities.TomatoSeed{},
}
