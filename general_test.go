package reflect_db

import (
	"fmt"
	"github.com/ElPeque/reflect-db/leaf_avg"
	"testing"
)

func TestGeneral(t *testing.T) {
	type cosa struct {
		uno      int
		Dos      int
		Tres     int
		Flotinga float64
	}

	type rarita struct {
		Reloco map[cosa]string
	}

	type cosita struct {
		Nombre            string
		Cosas             []*cosa
		Numeritos         []int
		NumeritosPointers []*int
		Weird             *rarita
	}

	type cosillas struct {
		Mapa *map[string]*cosita
		Pope *map[string]cosita
		algo string
		Otra int
	}

	var datos = cosillas{
		algo: "algoval",
		Otra: 5,
		Mapa: &map[string]*cosita{
			"ta": &cosita{
				Numeritos:         []int{100, 101, 102},
				NumeritosPointers: []*int{new(int), new(int), new(int)},
				Weird: &rarita{
					Reloco: map[cosa]string{
						cosa{
							uno:  1,
							Dos:  2,
							Tres: 3,
						}: "omfg",
					},
				},
				Nombre: "unacosita",
				Cosas: []*cosa{
					&cosa{
						uno:  1,
						Dos:  2,
						Tres: 3,
					},
				},
			},
			"te": &cosita{

				Nombre: "otracosita",
				Cosas: []*cosa{
					&cosa{
						uno:      1,
						Dos:      2,
						Tres:     3,
						Flotinga: 3.14,
					},
					&cosa{
						uno:  4,
						Dos:  5,
						Tres: 6,
					},
					&cosa{
						uno:  5,
						Dos:  6,
						Tres: 7,
					},
				},
			},
			"ti": &cosita{
				Nombre: "yotramas",
				Cosas: []*cosa{
					&cosa{
						uno:  1,
						Dos:  2,
						Tres: 3,
					},
					&cosa{
						uno:  0,
						Dos:  9,
						Tres: 8,
					},
				},
			},
		},
		Pope: &map[string]cosita{
			"ta": cosita{
				Weird: &rarita{
					Reloco: map[cosa]string{
						cosa{
							uno:  1,
							Dos:  2,
							Tres: 3,
						}: "omfg",
					},
				},
				Nombre: "unacosita",
				Cosas: []*cosa{
					&cosa{
						uno:  1,
						Dos:  2,
						Tres: 3,
					},
				},
			},
			"te": cosita{

				Nombre: "otracosita",
				Cosas: []*cosa{
					&cosa{
						uno:  1,
						Dos:  2,
						Tres: 3,
					},
					&cosa{
						uno:  4,
						Dos:  5,
						Tres: 6,
					},
					&cosa{
						uno:  5,
						Dos:  6,
						Tres: 7,
					},
				},
			},
			"ti": cosita{
				Nombre: "yotramas",
				Cosas: []*cosa{
					&cosa{
						uno:  1,
						Dos:  2,
						Tres: 3,
					},
					&cosa{
						uno:  0,
						Dos:  9,
						Tres: 8,
					},
				},
			},
		},
	}

	sum := 0
	avg := 0
	avgF := float64(0.0)
	avgUint64 := uint64(0)

	fmt.Println("avgF is at:", &avgF)

	Walk(
		&datos,
		All(leafPrintPath(), leafSum(&sum), leaf_avg.LeafAvg(&avg), leaf_avg.LeafAvg(&avgF), leaf_avg.LeafAvg(&avgUint64)),
		whilePattern([]string{"Mapa", "*", "Cosas", "0", "*"}),
	)

	fmt.Println("Sum:", sum)
	fmt.Println("Avg:", avg)
	fmt.Println("AvgF:", avgF)
	fmt.Println("avgUint64:", avgUint64)
}
