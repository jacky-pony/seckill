// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GoodsColumns holds the columns for the "goods" table.
	GoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "intro", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// GoodsTable holds the schema information for the "goods" table.
	GoodsTable = &schema.Table{
		Name:        "goods",
		Columns:     GoodsColumns,
		PrimaryKey:  []*schema.Column{GoodsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "sn", Type: field.TypeString},
		{Name: "gid", Type: field.TypeInt64},
		{Name: "uid", Type: field.TypeInt64},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:        "orders",
		Columns:     OrdersColumns,
		PrimaryKey:  []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// OrderGoodsColumns holds the columns for the "order_goods" table.
	OrderGoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "order_id", Type: field.TypeInt64},
		{Name: "goods_id", Type: field.TypeInt64},
		{Name: "goods_title", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// OrderGoodsTable holds the schema information for the "order_goods" table.
	OrderGoodsTable = &schema.Table{
		Name:        "order_goods",
		Columns:     OrderGoodsColumns,
		PrimaryKey:  []*schema.Column{OrderGoodsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GoodsTable,
		OrdersTable,
		OrderGoodsTable,
	}
)

func init() {
}
