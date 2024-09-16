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

var BidVotes = newBidVotesTable("public", "bid_votes", "")

type bidVotesTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	BidID     postgres.ColumnInteger
	Username  postgres.ColumnString
	Decision  postgres.ColumnBool
	CreatedAt postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type BidVotesTable struct {
	bidVotesTable

	EXCLUDED bidVotesTable
}

// AS creates new BidVotesTable with assigned alias
func (a BidVotesTable) AS(alias string) *BidVotesTable {
	return newBidVotesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new BidVotesTable with assigned schema name
func (a BidVotesTable) FromSchema(schemaName string) *BidVotesTable {
	return newBidVotesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new BidVotesTable with assigned table prefix
func (a BidVotesTable) WithPrefix(prefix string) *BidVotesTable {
	return newBidVotesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new BidVotesTable with assigned table suffix
func (a BidVotesTable) WithSuffix(suffix string) *BidVotesTable {
	return newBidVotesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newBidVotesTable(schemaName, tableName, alias string) *BidVotesTable {
	return &BidVotesTable{
		bidVotesTable: newBidVotesTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newBidVotesTableImpl("", "excluded", ""),
	}
}

func newBidVotesTableImpl(schemaName, tableName, alias string) bidVotesTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		BidIDColumn     = postgres.IntegerColumn("bid_id")
		UsernameColumn  = postgres.StringColumn("username")
		DecisionColumn  = postgres.BoolColumn("decision")
		CreatedAtColumn = postgres.TimestampColumn("created_at")
		allColumns      = postgres.ColumnList{IDColumn, BidIDColumn, UsernameColumn, DecisionColumn, CreatedAtColumn}
		mutableColumns  = postgres.ColumnList{BidIDColumn, UsernameColumn, DecisionColumn, CreatedAtColumn}
	)

	return bidVotesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		BidID:     BidIDColumn,
		Username:  UsernameColumn,
		Decision:  DecisionColumn,
		CreatedAt: CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
