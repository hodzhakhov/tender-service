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

var Employee = newEmployeeTable("public", "employee", "")

type employeeTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnString
	Username  postgres.ColumnString
	FirstName postgres.ColumnString
	LastName  postgres.ColumnString
	CreatedAt postgres.ColumnTimestamp
	UpdatedAt postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type EmployeeTable struct {
	employeeTable

	EXCLUDED employeeTable
}

// AS creates new EmployeeTable with assigned alias
func (a EmployeeTable) AS(alias string) *EmployeeTable {
	return newEmployeeTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new EmployeeTable with assigned schema name
func (a EmployeeTable) FromSchema(schemaName string) *EmployeeTable {
	return newEmployeeTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new EmployeeTable with assigned table prefix
func (a EmployeeTable) WithPrefix(prefix string) *EmployeeTable {
	return newEmployeeTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new EmployeeTable with assigned table suffix
func (a EmployeeTable) WithSuffix(suffix string) *EmployeeTable {
	return newEmployeeTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newEmployeeTable(schemaName, tableName, alias string) *EmployeeTable {
	return &EmployeeTable{
		employeeTable: newEmployeeTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newEmployeeTableImpl("", "excluded", ""),
	}
}

func newEmployeeTableImpl(schemaName, tableName, alias string) employeeTable {
	var (
		IDColumn        = postgres.StringColumn("id")
		UsernameColumn  = postgres.StringColumn("username")
		FirstNameColumn = postgres.StringColumn("first_name")
		LastNameColumn  = postgres.StringColumn("last_name")
		CreatedAtColumn = postgres.TimestampColumn("created_at")
		UpdatedAtColumn = postgres.TimestampColumn("updated_at")
		allColumns      = postgres.ColumnList{IDColumn, UsernameColumn, FirstNameColumn, LastNameColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns  = postgres.ColumnList{UsernameColumn, FirstNameColumn, LastNameColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return employeeTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Username:  UsernameColumn,
		FirstName: FirstNameColumn,
		LastName:  LastNameColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
