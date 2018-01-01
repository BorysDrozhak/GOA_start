// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "cellar": Application Media Types
//
// Command:
// $ goagen
// --design=BorysDrozhak/GOA_start/design
// --out=$(GOPATH)/src/BorysDrozhak/GOA_start
// --version=v1.3.0

package app

import (
	"github.com/goadesign/goa"
)

// A bottle of wine (default view)
//
// Identifier: application/vnd.goa.example.bottle+json; view=default
type GoaExampleBottle struct {
	// API href for making requests on the bottle
	Href string `form:"href" json:"href" xml:"href"`
	// Unique bottle ID
	ID int `form:"id" json:"id" xml:"id"`
	// Name of wine
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GoaExampleBottle media type instance.
func (mt *GoaExampleBottle) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}
