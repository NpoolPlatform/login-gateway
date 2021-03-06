// Code generated by entc, DO NOT EDIT.

package loginhistory

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the loginhistory type in the database.
	Label = "login_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldClientIP holds the string denoting the client_ip field in the database.
	FieldClientIP = "client_ip"
	// FieldUserAgent holds the string denoting the user_agent field in the database.
	FieldUserAgent = "user_agent"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// Table holds the table name of the loginhistory in the database.
	Table = "login_histories"
)

// Columns holds all SQL columns for loginhistory fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldUserID,
	FieldClientIP,
	FieldUserAgent,
	FieldLocation,
	FieldCreateAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
