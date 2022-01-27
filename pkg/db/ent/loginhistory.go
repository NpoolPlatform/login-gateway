// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/login-gateway/pkg/db/ent/loginhistory"
	"github.com/google/uuid"
)

// LoginHistory is the model entity for the LoginHistory schema.
type LoginHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// ClientIP holds the value of the "client_ip" field.
	ClientIP string `json:"client_ip,omitempty"`
	// UserAgent holds the value of the "user_agent" field.
	UserAgent string `json:"user_agent,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt int64 `json:"create_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LoginHistory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case loginhistory.FieldCreateAt:
			values[i] = new(sql.NullInt64)
		case loginhistory.FieldClientIP, loginhistory.FieldUserAgent:
			values[i] = new(sql.NullString)
		case loginhistory.FieldID, loginhistory.FieldAppID, loginhistory.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type LoginHistory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LoginHistory fields.
func (lh *LoginHistory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case loginhistory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				lh.ID = *value
			}
		case loginhistory.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				lh.AppID = *value
			}
		case loginhistory.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				lh.UserID = *value
			}
		case loginhistory.FieldClientIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_ip", values[i])
			} else if value.Valid {
				lh.ClientIP = value.String
			}
		case loginhistory.FieldUserAgent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_agent", values[i])
			} else if value.Valid {
				lh.UserAgent = value.String
			}
		case loginhistory.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				lh.CreateAt = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this LoginHistory.
// Note that you need to call LoginHistory.Unwrap() before calling this method if this LoginHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (lh *LoginHistory) Update() *LoginHistoryUpdateOne {
	return (&LoginHistoryClient{config: lh.config}).UpdateOne(lh)
}

// Unwrap unwraps the LoginHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (lh *LoginHistory) Unwrap() *LoginHistory {
	tx, ok := lh.config.driver.(*txDriver)
	if !ok {
		panic("ent: LoginHistory is not a transactional entity")
	}
	lh.config.driver = tx.drv
	return lh
}

// String implements the fmt.Stringer.
func (lh *LoginHistory) String() string {
	var builder strings.Builder
	builder.WriteString("LoginHistory(")
	builder.WriteString(fmt.Sprintf("id=%v", lh.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", lh.AppID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", lh.UserID))
	builder.WriteString(", client_ip=")
	builder.WriteString(lh.ClientIP)
	builder.WriteString(", user_agent=")
	builder.WriteString(lh.UserAgent)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", lh.CreateAt))
	builder.WriteByte(')')
	return builder.String()
}

// LoginHistories is a parsable slice of LoginHistory.
type LoginHistories []*LoginHistory

func (lh LoginHistories) config(cfg config) {
	for _i := range lh {
		lh[_i].config = cfg
	}
}
