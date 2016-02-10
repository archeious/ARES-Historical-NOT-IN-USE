package models

import (
	"github.com/revel/revel"
	"log"
	"webfrontend/app"
)

type Season struct {
	uuid     int64
	Name     string
	Series   *Series
	Episodes *[]Episode
	Number   string
}

func (s *Season) String() string {
	return s.Name
}

func (s *Season) Validate(v *revel.Validation) {
	v.Check(s.Name,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{255},
	)
}

func (s *Season) Add() {
	const query = "INSERT INTO season (name, series_id) VALUES (?,?)"

	stmt, err := app.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(s.Name, s.Series.Id)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	s.uuid = int64(id)
}
