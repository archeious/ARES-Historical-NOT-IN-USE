package models

import (
	"github.com/revel/revel"
)

type Episode struct {
	uuid   int64
	Series *Series
	Season *Season
	Number string
	Name   string
}

func (e *Episode) String() string {
	return e.Name
}

func (e *Episode) Validate(v *revel.Validation) {
	v.Check(e.Name,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{255},
	)
}
