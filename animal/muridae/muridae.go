package muridae

import "github/tiechui1994/memoryleak/animal"

type Muridae interface {
	animal.Animal
	Hole()
	Steal()
}
