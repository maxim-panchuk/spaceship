package carrier

import (
	"fmt"
	"log"
	"os"
	"spaceship/internal/util"
)

type CarrierGen struct {
	repo CarrierRepository
}

func NewCarrierGen(repo CarrierRepository) *CarrierGen {
	return &CarrierGen{
		repo: repo,
	}
}

func (g *CarrierGen) Generate() {
	amount := 20

	for i := 0; i < amount; i++ {
		speed := util.NewCryptoRand(10000, 100000)
		power := util.NewCryptoRand(1, 50)

		name := fmt.Sprintf("carrier-%v", speed+power)

		id, err := g.repo.Insert(name, int(power), int(speed))

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fmt.Printf("Carrier created!: id: %v, name: %s, speed: %v, power: %v\n", id, name, speed, power)
	}
}
