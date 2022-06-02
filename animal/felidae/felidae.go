package felidae

import "github/tiechui1994/memoryleak/animal"

type Felidae interface {
	animal.Animal
	Climb()
	Sneak()
}
