package controllers

import (
	"github.com/revel/revel"
	"time"
)

type Faq struct {
	*revel.Controller
}

func (c Faq) Index() revel.Result {
	var date string = time.Now().Format("2006")
	var title string = "Lösenord.xyz"
	// Show frequently asked questions
	var action string = c.Action
	return c.Render(action, title, date)
}
