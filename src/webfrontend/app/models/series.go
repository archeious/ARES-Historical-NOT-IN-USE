package models

import (
	"github.com/revel/revel"
	"log"
	"webfrontend/app"
)

type Series struct {
	Id          int64
	Seasons     []*Season
	Name        string
	Description string
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

func (s *Series) GetSeasons() []Season {
	return GetSeasonsBySeriesId(s.Id)
}

func (s *Series) SeasonCount() int64 {
	const query = "SELECT count(id) from season where series_id = ? "

	var count int64

	row := app.DB.QueryRow(query, s.Id)
	if err := row.Scan(&count); err != nil {
		log.Fatal(err)
	}
	return count
}

func (s *Series) EpisodeCount() int64 {
	const query = "select count(*) from episode e join season sea on e.season_id = sea.id join series ser on sea.series_id=ser.id and ser.id = ?;"
	var count int64

	row := app.DB.QueryRow(query, s.Id)
	if err := row.Scan(&count); err != nil {
		log.Fatal(err)
	}
	return count
}

func GetAllSeries() []*Series {
	const query = "SELECT id, name,description from series"
	//const query = "SELECT ser.id, ser.name, case when temp.count>= 1 then temp.count else 0 end from series ser left join (select series_id, count(id) as count from season group by series_id) temp on temp.series_id = ser.id;"
	var name string
	var id int64
	var desc string
	series := make([]*Series, 0)

	rows, err := app.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &desc)
		if err != nil {
			log.Fatal(err)
		}
		series = append(series, &Series{Name: name, Id: id, Description: desc})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return series
}

func GetSeriesByName(n string) *Series {
	const query = "SELECT id, name,desc from series where name = ? "

	var name string
	var id int64
	var desc string

	row := app.DB.QueryRow(query, n)
	if err := row.Scan(&id, &name, &desc); err != nil {
		log.Fatal(err)
	}
	return &Series{Name: name, Id: id, Description: desc}
}

func GetSeriesById(id int64) *Series {
	const query = "SELECT id, name, description from series where id = ? "

	var name string
	var desc string

	row := app.DB.QueryRow(query, id)
	if err := row.Scan(&id, &name, &desc); err != nil {
		log.Fatal(err)
	}
	return &Series{Name: name, Id: id, Description: desc}
}
