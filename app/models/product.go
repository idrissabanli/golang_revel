package models

import (
	"github.com/revel/revel"
	"gopkg.in/gorp.v2"
)


type Product struct {
	HotelId		int
	CategoryId	int
	Name		string
	Image		string
	Price		int

	// Transient
	Category	*Category
}

func (product *Product) Validate(v *revel.Validation) {
	v.Check(product.Name,
		revel.Required{},
		revel.MaxSize{50},
	)

	v.MaxSize(product.Image, 255)
}

func (product *Product) PostGet(exe gorp.SqlExecutor) error {
	var (
		obj interface{}
		err error
	)
	obj, err = exe.Get(Category{}, b.CategoryId)
	if err != nil {
		return fmt.Errorf("Error loading a category's category (%d): %s", b.CategoryId, err)
	}

	b.Category = obj.(*Category)

	return nil
}