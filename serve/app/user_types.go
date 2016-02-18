//************************************************************************//
// API "jmdict": Application User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/nii236/jmdict/serve
// --design=github.com/nii236/jmdict/serve/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// Word type
type Word struct {
	// Operand name
	Name string `json:"name" xml:"name"`
}

// Validate validates the type instance.
func (ut *Word) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}

	return
}
