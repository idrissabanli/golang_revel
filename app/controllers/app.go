package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	// salam := "Hello"
	// return c.Render(salam)
	var res Response
	res.Message = "salam"
	return c.RenderJSON(res)
}
