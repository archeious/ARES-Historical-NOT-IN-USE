package controllers

import (
	"github.com/revel/revel"
	"webfrontend/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) About() revel.Result {
	return c.Render()
}

func (c App) Contact() revel.Result {
	return c.Render()
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) List() revel.Result {
	series := models.GetAllSeries()
	return c.Render(series)
}
