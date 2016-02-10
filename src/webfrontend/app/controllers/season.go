package controllers

import (
	"github.com/revel/revel"
	"webfrontend/app/models"
)

type Season struct {
	App
}

func (c Season) Index() revel.Result {
	series := models.GetAllSeries()
	return c.Render(series)
}

func (c Season) Display(id int) revel.Result {
	series := models.GetSeriesById(int64(id))
	return c.Render(series)
}

func (c Season) Add(seriesId int) revel.Result {
	series := models.GetSeriesById(int64(seriesId))
	return c.Render(series)
}
