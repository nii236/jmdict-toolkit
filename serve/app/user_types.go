//************************************************************************//
// API "jmdict": Application User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/nii236/jmdict-toolkit/serve
// --design=github.com/nii236/jmdict-toolkit/serve/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// Word type
type Word struct {
	// Word to be translated
	Word string `json:"word" xml:"word"`
}

// Validate validates the type instance.
func (ut *Word) Validate() (err error) {
	if ut.Word == "" {
		err = goa.MissingAttributeError(`response`, "word", err)
	}

	return
}
