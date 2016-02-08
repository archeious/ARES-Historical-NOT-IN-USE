package models

import (
	"github.com/revel/revel"
	"log"
	"webfrontend/app"
)

type Series struct {
	Id      int64
	Seasons []*Season
	Name    string
}

func (s *Series) String() string {
	return s.Name
}

func (s *Series) Validate(v *revel.Validation) {
	v.Check(s.Name,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{255},
	)
}

func GetAllSeries() []Series {
	const query = "SELECT id, name from series"

	var name string
	var id int64
	series := make([]Series, 0)

	rows, err := app.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		series = append(series, Series{Name: name, Id: id})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return series
}

func GetSeriesByName(n string) *Series {
	const query = "SELECT id, name from series where name = ? "

	var name string
	var id int64

	row := app.DB.QueryRow(query, n)
	if err := row.Scan(&id, &name); err != nil {
		log.Fatal(err)
	}
	return &Series{Name: name, Id: id}
}
