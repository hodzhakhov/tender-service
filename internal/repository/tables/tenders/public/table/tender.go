//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Tender = newTenderTable("public", "tender", "")

type tenderTable struct {
	postgres.Table

	// Columns
	ID              postgres.ColumnInteger
	Name            postgres.ColumnString
	Description     postgres.ColumnString
	ServiceType     postgres.ColumnString
	Status          postgres.ColumnString
	OrganizationID  postgres.ColumnString
	CreatorUsername postgres.ColumnString
	Version         postgres.ColumnInteger
	CreatedAt       postgres.ColumnTimestamp
	UpdatedAt       postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TenderTable struct {
	tenderTable

	EXCLUDED tenderTable
}

// AS creates new TenderTable with assigned alias
func (a TenderTable) AS(alias string) *TenderTable {
	return newTenderTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TenderTable with assigned schema name
func (a TenderTable) FromSchema(schemaName string) *TenderTable {
	return newTenderTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TenderTable with assigned table prefix
func (a TenderTable) WithPrefix(prefix string) *TenderTable {
	return newTenderTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TenderTable with assigned table suffix
func (a TenderTable) WithSuffix(suffix string) *TenderTable {
	return newTenderTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTenderTable(schemaName, tableName, alias string) *TenderTable {
	return &TenderTable{
		tenderTable: newTenderTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newTenderTableImpl("", "excluded", ""),
	}
}

func newTenderTableImpl(schemaName, tableName, alias string) tenderTable {
	var (
		IDColumn              = postgres.IntegerColumn("id")
		NameColumn            = postgres.StringColumn("name")
		DescriptionColumn     = postgres.StringColumn("description")
		ServiceTypeColumn     = postgres.StringColumn("service_type")
		StatusColumn          = postgres.StringColumn("status")
		OrganizationIDColumn  = postgres.StringColumn("organization_id")
		CreatorUsernameColumn = postgres.StringColumn("creator_username")
		VersionColumn         = postgres.IntegerColumn("version")
		CreatedAtColumn       = postgres.TimestampColumn("created_at")
		UpdatedAtColumn       = postgres.TimestampColumn("updated_at")
		allColumns            = postgres.ColumnList{IDColumn, NameColumn, DescriptionColumn, ServiceTypeColumn, StatusColumn, OrganizationIDColumn, CreatorUsernameColumn, VersionColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns        = postgres.ColumnList{NameColumn, DescriptionColumn, ServiceTypeColumn, StatusColumn, OrganizationIDColumn, CreatorUsernameColumn, VersionColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return tenderTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		Name:            NameColumn,
		Description:     DescriptionColumn,
		ServiceType:     ServiceTypeColumn,
		Status:          StatusColumn,
		OrganizationID:  OrganizationIDColumn,
		CreatorUsername: CreatorUsernameColumn,
		Version:         VersionColumn,
		CreatedAt:       CreatedAtColumn,
		UpdatedAt:       UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
