package models

import (
	"github.com/revel/revel"
)


type Category struct {
	CategoryId	int
	Title		string
}

func (product *Category) Validate(v *revel.Validation) {
	v.Check(product.Name,
		revel.Required{},
		revel.MaxSize{50},
	)
}