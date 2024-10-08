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

type Bid struct {
	ID              int32 `sql:"primary_key"`
	Name            string
	Description     *string
	Status          *string
	TenderID        *int32
	OrganizationID  *uuid.UUID
	CreatorUsername string
	Version         *int32
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}
