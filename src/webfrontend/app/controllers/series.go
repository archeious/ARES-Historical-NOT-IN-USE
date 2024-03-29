package controllers

import (
	"github.com/revel/revel"
	"math/rand"
	"time"
	"webfrontend/app/models"
)

type Series struct {
	App
}

func (c Series) Index() revel.Result {
	series := models.GetAllSeries()
	revel.INFO.Printf("%s\n", c.Session["Test"])
	return c.Render(series)
}

func (c Series) Display(id int) revel.Result {
	s1 := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(s1)
	series := models.GetSeriesById(int64(id))
	seasons := series.GetSeasons()
	return c.Render(series, seasons, rng)
}
