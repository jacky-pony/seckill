// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/peter-wow/seckill/app/goods/service/internal/data/ent/goods"
)

// GoodsCreate is the builder for creating a Goods entity.
type GoodsCreate struct {
	config
	mutation *GoodsMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (gc *GoodsCreate) SetTitle(s string) *GoodsCreate {
	gc.mutation.SetTitle(s)
	return gc
}

// SetIntro sets the "intro" field.
func (gc *GoodsCreate) SetIntro(s string) *GoodsCreate {
	gc.mutation.SetIntro(s)
	return gc
}

// SetCreatedAt sets the "created_at" field.
func (gc *GoodsCreate) SetCreatedAt(t time.Time) *GoodsCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GoodsCreate) SetNillableCreatedAt(t *time.Time) *GoodsCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GoodsCreate) SetUpdatedAt(t time.Time) *GoodsCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GoodsCreate) SetNillableUpdatedAt(t *time.Time) *GoodsCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GoodsCreate) SetID(i int64) *GoodsCreate {
	gc.mutation.SetID(i)
	return gc
}

// Mutation returns the GoodsMutation object of the builder.
func (gc *GoodsCreate) Mutation() *GoodsMutation {
	return gc.mutation
}

// Save creates the Goods in the database.
func (gc *GoodsCreate) Save(ctx context.Context) (*Goods, error) {
	var (
		err  error
		node *Goods
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GoodsCreate) SaveX(ctx context.Context) *Goods {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (gc *GoodsCreate) defaults() {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := goods.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		v := goods.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GoodsCreate) check() error {
	if _, ok := gc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if _, ok := gc.mutation.Intro(); !ok {
		return &ValidationError{Name: "intro", err: errors.New("ent: missing required field \"intro\"")}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	return nil
}

func (gc *GoodsCreate) sqlSave(ctx context.Context) (*Goods, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (gc *GoodsCreate) createSpec() (*Goods, *sqlgraph.CreateSpec) {
	var (
		_node = &Goods{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: goods.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: goods.FieldID,
			},
		}
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goods.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := gc.mutation.Intro(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goods.FieldIntro,
		})
		_node.Intro = value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goods.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goods.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// GoodsCreateBulk is the builder for creating many Goods entities in bulk.
type GoodsCreateBulk struct {
	config
	builders []*GoodsCreate
}

// Save creates the Goods entities in the database.
func (gcb *GoodsCreateBulk) Save(ctx context.Context) ([]*Goods, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Goods, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoodsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GoodsCreateBulk) SaveX(ctx context.Context) []*Goods {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
