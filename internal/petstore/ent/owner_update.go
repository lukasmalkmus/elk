// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/masseelch/elk/internal/petstore/ent/owner"
	"github.com/masseelch/elk/internal/petstore/ent/pet"
	"github.com/masseelch/elk/internal/petstore/ent/predicate"
)

// OwnerUpdate is the builder for updating Owner entities.
type OwnerUpdate struct {
	config
	hooks    []Hook
	mutation *OwnerMutation
}

// Where adds a new predicate for the OwnerUpdate builder.
func (ou *OwnerUpdate) Where(ps ...predicate.Owner) *OwnerUpdate {
	ou.mutation.predicates = append(ou.mutation.predicates, ps...)
	return ou
}

// SetName sets the "name" field.
func (ou *OwnerUpdate) SetName(s string) *OwnerUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetAge sets the "age" field.
func (ou *OwnerUpdate) SetAge(s string) *OwnerUpdate {
	ou.mutation.SetAge(s)
	return ou
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (ou *OwnerUpdate) AddPetIDs(ids ...int) *OwnerUpdate {
	ou.mutation.AddPetIDs(ids...)
	return ou
}

// AddPets adds the "pets" edges to the Pet entity.
func (ou *OwnerUpdate) AddPets(p ...*Pet) *OwnerUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.AddPetIDs(ids...)
}

// Mutation returns the OwnerMutation object of the builder.
func (ou *OwnerUpdate) Mutation() *OwnerMutation {
	return ou.mutation
}

// ClearPets clears all "pets" edges to the Pet entity.
func (ou *OwnerUpdate) ClearPets() *OwnerUpdate {
	ou.mutation.ClearPets()
	return ou
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (ou *OwnerUpdate) RemovePetIDs(ids ...int) *OwnerUpdate {
	ou.mutation.RemovePetIDs(ids...)
	return ou
}

// RemovePets removes "pets" edges to Pet entities.
func (ou *OwnerUpdate) RemovePets(p ...*Pet) *OwnerUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.RemovePetIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OwnerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OwnerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OwnerUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OwnerUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OwnerUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OwnerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   owner.Table,
			Columns: owner.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: owner.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: owner.FieldName,
		})
	}
	if value, ok := ou.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: owner.FieldAge,
		})
	}
	if ou.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedPetsIDs(); len(nodes) > 0 && !ou.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{owner.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OwnerUpdateOne is the builder for updating a single Owner entity.
type OwnerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OwnerMutation
}

// SetName sets the "name" field.
func (ouo *OwnerUpdateOne) SetName(s string) *OwnerUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetAge sets the "age" field.
func (ouo *OwnerUpdateOne) SetAge(s string) *OwnerUpdateOne {
	ouo.mutation.SetAge(s)
	return ouo
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (ouo *OwnerUpdateOne) AddPetIDs(ids ...int) *OwnerUpdateOne {
	ouo.mutation.AddPetIDs(ids...)
	return ouo
}

// AddPets adds the "pets" edges to the Pet entity.
func (ouo *OwnerUpdateOne) AddPets(p ...*Pet) *OwnerUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.AddPetIDs(ids...)
}

// Mutation returns the OwnerMutation object of the builder.
func (ouo *OwnerUpdateOne) Mutation() *OwnerMutation {
	return ouo.mutation
}

// ClearPets clears all "pets" edges to the Pet entity.
func (ouo *OwnerUpdateOne) ClearPets() *OwnerUpdateOne {
	ouo.mutation.ClearPets()
	return ouo
}

// RemovePetIDs removes the "pets" edge to Pet entities by IDs.
func (ouo *OwnerUpdateOne) RemovePetIDs(ids ...int) *OwnerUpdateOne {
	ouo.mutation.RemovePetIDs(ids...)
	return ouo
}

// RemovePets removes "pets" edges to Pet entities.
func (ouo *OwnerUpdateOne) RemovePets(p ...*Pet) *OwnerUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.RemovePetIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OwnerUpdateOne) Select(field string, fields ...string) *OwnerUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Owner entity.
func (ouo *OwnerUpdateOne) Save(ctx context.Context) (*Owner, error) {
	var (
		err  error
		node *Owner
	)
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OwnerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OwnerUpdateOne) SaveX(ctx context.Context) *Owner {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OwnerUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OwnerUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OwnerUpdateOne) sqlSave(ctx context.Context) (_node *Owner, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   owner.Table,
			Columns: owner.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: owner.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Owner.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, owner.FieldID)
		for _, f := range fields {
			if !owner.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != owner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: owner.FieldName,
		})
	}
	if value, ok := ouo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: owner.FieldAge,
		})
	}
	if ouo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedPetsIDs(); len(nodes) > 0 && !ouo.mutation.PetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   owner.PetsTable,
			Columns: []string{owner.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Owner{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{owner.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
