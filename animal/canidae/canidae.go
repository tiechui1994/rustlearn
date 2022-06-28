package canidae

import "github/tiechui1994/memoryleak/animal"

type Canidae interface {
	animal.Animal
	Run()
	Howl()
}
