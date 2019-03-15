package yeay

import (
	"fmt"
)

type Yeay struct {
	name string
}

func DoMagic() {
	var x *Yeay = &Yeay{
		name: "Hi there",
	}

	var y *Yeay = nil

	fmt.Printf("wat %s", x.name)
	fmt.Printf("wat %s", y.name)
}
