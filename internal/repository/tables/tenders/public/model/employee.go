//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type Employee struct {
	ID        uuid.UUID `sql:"primary_key"`
	Username  string
	FirstName *string
	LastName  *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
