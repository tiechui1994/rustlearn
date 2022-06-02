package animal

import (
	"github/tiechui1994/memoryleak/animal/canidae/dog"
	"github/tiechui1994/memoryleak/animal/canidae/wolf"
	"github/tiechui1994/memoryleak/animal/felidae/cat"
	"github/tiechui1994/memoryleak/animal/felidae/tiger"
	"github/tiechui1994/memoryleak/animal/muridae/mouse"
)

var (
	AllAnimals = []Animal{
		&dog.Dog{},
		&wolf.Wolf{},

		&cat.Cat{},
		&tiger.Tiger{},

		&mouse.Mouse{},
	}
)

type Animal interface {
	Name() string
	Live()

	Eat()
	Drink()
	Shit()
	Pee()
}
