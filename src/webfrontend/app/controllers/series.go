package controllers

import (
	"github.com/revel/revel"
	"webfrontend/app/models"
)

type Series struct {
	App
}

func (c Series) Index() revel.Result {
	series := models.GetAllSeries()
	return c.Render(series)
}

func (c Series) Display(id int) revel.Result {
	series := models.GetSeriesById(int64(id))
	return c.Render(series)
}
