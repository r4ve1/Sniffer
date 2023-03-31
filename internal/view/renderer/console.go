package renderer

import (
	"fmt"
)

type Console struct {
}

func (c *Console) RenderBrief(b *Brief) {
	fmt.Printf("%v\n", b)
}

func (c *Console) RenderDetail(d *Detail) {
	fmt.Printf("%v\n", d)
}

func (c *Console) Reset() {

}
