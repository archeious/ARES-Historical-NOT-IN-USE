package models

import (
	"github.com/revel/revel"
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
