// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/peter-wow/seckill/app/goods/service/internal/data/ent/goods"
)

// Goods is the model entity for the Goods schema.
type Goods struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	// 标题
	Title string `json:"title,omitempty"`
	// Intro holds the value of the "intro" field.
	// 简介
	Intro string `json:"intro,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Goods) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goods.FieldID:
			values[i] = new(sql.NullInt64)
		case goods.FieldTitle, goods.FieldIntro:
			values[i] = new(sql.NullString)
		case goods.FieldCreatedAt, goods.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Goods", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Goods fields.
func (_go *Goods) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goods.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_go.ID = int64(value.Int64)
		case goods.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				_go.Title = value.String
			}
		case goods.FieldIntro:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field intro", values[i])
			} else if value.Valid {
				_go.Intro = value.String
			}
		case goods.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				_go.CreatedAt = value.Time
			}
		case goods.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				_go.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Goods.
// Note that you need to call Goods.Unwrap() before calling this method if this Goods
// was returned from a transaction, and the transaction was committed or rolled back.
func (_go *Goods) Update() *GoodsUpdateOne {
	return (&GoodsClient{config: _go.config}).UpdateOne(_go)
}

// Unwrap unwraps the Goods entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_go *Goods) Unwrap() *Goods {
	tx, ok := _go.config.driver.(*txDriver)
	if !ok {
		panic("ent: Goods is not a transactional entity")
	}
	_go.config.driver = tx.drv
	return _go
}

// String implements the fmt.Stringer.
func (_go *Goods) String() string {
	var builder strings.Builder
	builder.WriteString("Goods(")
	builder.WriteString(fmt.Sprintf("id=%v", _go.ID))
	builder.WriteString(", title=")
	builder.WriteString(_go.Title)
	builder.WriteString(", intro=")
	builder.WriteString(_go.Intro)
	builder.WriteString(", created_at=")
	builder.WriteString(_go.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(_go.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// GoodsSlice is a parsable slice of Goods.
type GoodsSlice []*Goods

func (_go GoodsSlice) config(cfg config) {
	for _i := range _go {
		_go[_i].config = cfg
	}
}
