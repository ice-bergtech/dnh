// Code generated by ent, DO NOT EDIT.

package model_ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/asninfo"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/ipaddress"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/predicate"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/registrar"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/scanjob"
	"github.com/ice-bergtech/dnh/src/internal/model_ent/whois"
)

// ASNInfoUpdate is the builder for updating ASNInfo entities.
type ASNInfoUpdate struct {
	config
	hooks    []Hook
	mutation *ASNInfoMutation
}

// Where appends a list predicates to the ASNInfoUpdate builder.
func (aiu *ASNInfoUpdate) Where(ps ...predicate.ASNInfo) *ASNInfoUpdate {
	aiu.mutation.Where(ps...)
	return aiu
}

// SetAsn sets the "asn" field.
func (aiu *ASNInfoUpdate) SetAsn(i int) *ASNInfoUpdate {
	aiu.mutation.ResetAsn()
	aiu.mutation.SetAsn(i)
	return aiu
}

// SetNillableAsn sets the "asn" field if the given value is not nil.
func (aiu *ASNInfoUpdate) SetNillableAsn(i *int) *ASNInfoUpdate {
	if i != nil {
		aiu.SetAsn(*i)
	}
	return aiu
}

// AddAsn adds i to the "asn" field.
func (aiu *ASNInfoUpdate) AddAsn(i int) *ASNInfoUpdate {
	aiu.mutation.AddAsn(i)
	return aiu
}

// SetCountry sets the "country" field.
func (aiu *ASNInfoUpdate) SetCountry(s string) *ASNInfoUpdate {
	aiu.mutation.SetCountry(s)
	return aiu
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (aiu *ASNInfoUpdate) SetNillableCountry(s *string) *ASNInfoUpdate {
	if s != nil {
		aiu.SetCountry(*s)
	}
	return aiu
}

// SetRegistry sets the "registry" field.
func (aiu *ASNInfoUpdate) SetRegistry(s string) *ASNInfoUpdate {
	aiu.mutation.SetRegistry(s)
	return aiu
}

// SetNillableRegistry sets the "registry" field if the given value is not nil.
func (aiu *ASNInfoUpdate) SetNillableRegistry(s *string) *ASNInfoUpdate {
	if s != nil {
		aiu.SetRegistry(*s)
	}
	return aiu
}

// AddScanIDs adds the "scan" edge to the ScanJob entity by IDs.
func (aiu *ASNInfoUpdate) AddScanIDs(ids ...int) *ASNInfoUpdate {
	aiu.mutation.AddScanIDs(ids...)
	return aiu
}

// AddScan adds the "scan" edges to the ScanJob entity.
func (aiu *ASNInfoUpdate) AddScan(s ...*ScanJob) *ASNInfoUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return aiu.AddScanIDs(ids...)
}

// AddIpaddresIDs adds the "ipaddress" edge to the IPAddress entity by IDs.
func (aiu *ASNInfoUpdate) AddIpaddresIDs(ids ...int) *ASNInfoUpdate {
	aiu.mutation.AddIpaddresIDs(ids...)
	return aiu
}

// AddIpaddress adds the "ipaddress" edges to the IPAddress entity.
func (aiu *ASNInfoUpdate) AddIpaddress(i ...*IPAddress) *ASNInfoUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return aiu.AddIpaddresIDs(ids...)
}

// AddRegistrarIDs adds the "registrar" edge to the Registrar entity by IDs.
func (aiu *ASNInfoUpdate) AddRegistrarIDs(ids ...int) *ASNInfoUpdate {
	aiu.mutation.AddRegistrarIDs(ids...)
	return aiu
}

// AddRegistrar adds the "registrar" edges to the Registrar entity.
func (aiu *ASNInfoUpdate) AddRegistrar(r ...*Registrar) *ASNInfoUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return aiu.AddRegistrarIDs(ids...)
}

// AddWhoiIDs adds the "whois" edge to the Whois entity by IDs.
func (aiu *ASNInfoUpdate) AddWhoiIDs(ids ...int) *ASNInfoUpdate {
	aiu.mutation.AddWhoiIDs(ids...)
	return aiu
}

// AddWhois adds the "whois" edges to the Whois entity.
func (aiu *ASNInfoUpdate) AddWhois(w ...*Whois) *ASNInfoUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return aiu.AddWhoiIDs(ids...)
}

// Mutation returns the ASNInfoMutation object of the builder.
func (aiu *ASNInfoUpdate) Mutation() *ASNInfoMutation {
	return aiu.mutation
}

// ClearScan clears all "scan" edges to the ScanJob entity.
func (aiu *ASNInfoUpdate) ClearScan() *ASNInfoUpdate {
	aiu.mutation.ClearScan()
	return aiu
}

// RemoveScanIDs removes the "scan" edge to ScanJob entities by IDs.
func (aiu *ASNInfoUpdate) RemoveScanIDs(ids ...int) *ASNInfoUpdate {
	aiu.mutation.RemoveScanIDs(ids...)
	return aiu
}

// RemoveScan removes "scan" edges to ScanJob entities.
func (aiu *ASNInfoUpdate) RemoveScan(s ...*ScanJob) *ASNInfoUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return aiu.RemoveScanIDs(ids...)
}

// ClearIpaddress clears all "ipaddress" edges to the IPAddress entity.
func (aiu *ASNInfoUpdate) ClearIpaddress() *ASNInfoUpdate {
	aiu.mutation.ClearIpaddress()
	return aiu
}

// RemoveIpaddresIDs removes the "ipaddress" edge to IPAddress entities by IDs.
func (aiu *ASNInfoUpdate) RemoveIpaddresIDs(ids ...int) *ASNInfoUpdate {
	aiu.mutation.RemoveIpaddresIDs(ids...)
	return aiu
}

// RemoveIpaddress removes "ipaddress" edges to IPAddress entities.
func (aiu *ASNInfoUpdate) RemoveIpaddress(i ...*IPAddress) *ASNInfoUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return aiu.RemoveIpaddresIDs(ids...)
}

// ClearRegistrar clears all "registrar" edges to the Registrar entity.
func (aiu *ASNInfoUpdate) ClearRegistrar() *ASNInfoUpdate {
	aiu.mutation.ClearRegistrar()
	return aiu
}

// RemoveRegistrarIDs removes the "registrar" edge to Registrar entities by IDs.
func (aiu *ASNInfoUpdate) RemoveRegistrarIDs(ids ...int) *ASNInfoUpdate {
	aiu.mutation.RemoveRegistrarIDs(ids...)
	return aiu
}

// RemoveRegistrar removes "registrar" edges to Registrar entities.
func (aiu *ASNInfoUpdate) RemoveRegistrar(r ...*Registrar) *ASNInfoUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return aiu.RemoveRegistrarIDs(ids...)
}

// ClearWhois clears all "whois" edges to the Whois entity.
func (aiu *ASNInfoUpdate) ClearWhois() *ASNInfoUpdate {
	aiu.mutation.ClearWhois()
	return aiu
}

// RemoveWhoiIDs removes the "whois" edge to Whois entities by IDs.
func (aiu *ASNInfoUpdate) RemoveWhoiIDs(ids ...int) *ASNInfoUpdate {
	aiu.mutation.RemoveWhoiIDs(ids...)
	return aiu
}

// RemoveWhois removes "whois" edges to Whois entities.
func (aiu *ASNInfoUpdate) RemoveWhois(w ...*Whois) *ASNInfoUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return aiu.RemoveWhoiIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aiu *ASNInfoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aiu.sqlSave, aiu.mutation, aiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aiu *ASNInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := aiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aiu *ASNInfoUpdate) Exec(ctx context.Context) error {
	_, err := aiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiu *ASNInfoUpdate) ExecX(ctx context.Context) {
	if err := aiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aiu *ASNInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(asninfo.Table, asninfo.Columns, sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt))
	if ps := aiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiu.mutation.Asn(); ok {
		_spec.SetField(asninfo.FieldAsn, field.TypeInt, value)
	}
	if value, ok := aiu.mutation.AddedAsn(); ok {
		_spec.AddField(asninfo.FieldAsn, field.TypeInt, value)
	}
	if value, ok := aiu.mutation.Country(); ok {
		_spec.SetField(asninfo.FieldCountry, field.TypeString, value)
	}
	if value, ok := aiu.mutation.Registry(); ok {
		_spec.SetField(asninfo.FieldRegistry, field.TypeString, value)
	}
	if aiu.mutation.ScanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.ScanTable,
			Columns: asninfo.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.RemovedScanIDs(); len(nodes) > 0 && !aiu.mutation.ScanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.ScanTable,
			Columns: asninfo.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.ScanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.ScanTable,
			Columns: asninfo.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aiu.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.IpaddressTable,
			Columns: asninfo.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.RemovedIpaddressIDs(); len(nodes) > 0 && !aiu.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.IpaddressTable,
			Columns: asninfo.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.IpaddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.IpaddressTable,
			Columns: asninfo.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aiu.mutation.RegistrarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.RegistrarTable,
			Columns: asninfo.RegistrarPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(registrar.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.RemovedRegistrarIDs(); len(nodes) > 0 && !aiu.mutation.RegistrarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.RegistrarTable,
			Columns: asninfo.RegistrarPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(registrar.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.RegistrarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.RegistrarTable,
			Columns: asninfo.RegistrarPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(registrar.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aiu.mutation.WhoisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.WhoisTable,
			Columns: asninfo.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.RemovedWhoisIDs(); len(nodes) > 0 && !aiu.mutation.WhoisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.WhoisTable,
			Columns: asninfo.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiu.mutation.WhoisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.WhoisTable,
			Columns: asninfo.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asninfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aiu.mutation.done = true
	return n, nil
}

// ASNInfoUpdateOne is the builder for updating a single ASNInfo entity.
type ASNInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ASNInfoMutation
}

// SetAsn sets the "asn" field.
func (aiuo *ASNInfoUpdateOne) SetAsn(i int) *ASNInfoUpdateOne {
	aiuo.mutation.ResetAsn()
	aiuo.mutation.SetAsn(i)
	return aiuo
}

// SetNillableAsn sets the "asn" field if the given value is not nil.
func (aiuo *ASNInfoUpdateOne) SetNillableAsn(i *int) *ASNInfoUpdateOne {
	if i != nil {
		aiuo.SetAsn(*i)
	}
	return aiuo
}

// AddAsn adds i to the "asn" field.
func (aiuo *ASNInfoUpdateOne) AddAsn(i int) *ASNInfoUpdateOne {
	aiuo.mutation.AddAsn(i)
	return aiuo
}

// SetCountry sets the "country" field.
func (aiuo *ASNInfoUpdateOne) SetCountry(s string) *ASNInfoUpdateOne {
	aiuo.mutation.SetCountry(s)
	return aiuo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (aiuo *ASNInfoUpdateOne) SetNillableCountry(s *string) *ASNInfoUpdateOne {
	if s != nil {
		aiuo.SetCountry(*s)
	}
	return aiuo
}

// SetRegistry sets the "registry" field.
func (aiuo *ASNInfoUpdateOne) SetRegistry(s string) *ASNInfoUpdateOne {
	aiuo.mutation.SetRegistry(s)
	return aiuo
}

// SetNillableRegistry sets the "registry" field if the given value is not nil.
func (aiuo *ASNInfoUpdateOne) SetNillableRegistry(s *string) *ASNInfoUpdateOne {
	if s != nil {
		aiuo.SetRegistry(*s)
	}
	return aiuo
}

// AddScanIDs adds the "scan" edge to the ScanJob entity by IDs.
func (aiuo *ASNInfoUpdateOne) AddScanIDs(ids ...int) *ASNInfoUpdateOne {
	aiuo.mutation.AddScanIDs(ids...)
	return aiuo
}

// AddScan adds the "scan" edges to the ScanJob entity.
func (aiuo *ASNInfoUpdateOne) AddScan(s ...*ScanJob) *ASNInfoUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return aiuo.AddScanIDs(ids...)
}

// AddIpaddresIDs adds the "ipaddress" edge to the IPAddress entity by IDs.
func (aiuo *ASNInfoUpdateOne) AddIpaddresIDs(ids ...int) *ASNInfoUpdateOne {
	aiuo.mutation.AddIpaddresIDs(ids...)
	return aiuo
}

// AddIpaddress adds the "ipaddress" edges to the IPAddress entity.
func (aiuo *ASNInfoUpdateOne) AddIpaddress(i ...*IPAddress) *ASNInfoUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return aiuo.AddIpaddresIDs(ids...)
}

// AddRegistrarIDs adds the "registrar" edge to the Registrar entity by IDs.
func (aiuo *ASNInfoUpdateOne) AddRegistrarIDs(ids ...int) *ASNInfoUpdateOne {
	aiuo.mutation.AddRegistrarIDs(ids...)
	return aiuo
}

// AddRegistrar adds the "registrar" edges to the Registrar entity.
func (aiuo *ASNInfoUpdateOne) AddRegistrar(r ...*Registrar) *ASNInfoUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return aiuo.AddRegistrarIDs(ids...)
}

// AddWhoiIDs adds the "whois" edge to the Whois entity by IDs.
func (aiuo *ASNInfoUpdateOne) AddWhoiIDs(ids ...int) *ASNInfoUpdateOne {
	aiuo.mutation.AddWhoiIDs(ids...)
	return aiuo
}

// AddWhois adds the "whois" edges to the Whois entity.
func (aiuo *ASNInfoUpdateOne) AddWhois(w ...*Whois) *ASNInfoUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return aiuo.AddWhoiIDs(ids...)
}

// Mutation returns the ASNInfoMutation object of the builder.
func (aiuo *ASNInfoUpdateOne) Mutation() *ASNInfoMutation {
	return aiuo.mutation
}

// ClearScan clears all "scan" edges to the ScanJob entity.
func (aiuo *ASNInfoUpdateOne) ClearScan() *ASNInfoUpdateOne {
	aiuo.mutation.ClearScan()
	return aiuo
}

// RemoveScanIDs removes the "scan" edge to ScanJob entities by IDs.
func (aiuo *ASNInfoUpdateOne) RemoveScanIDs(ids ...int) *ASNInfoUpdateOne {
	aiuo.mutation.RemoveScanIDs(ids...)
	return aiuo
}

// RemoveScan removes "scan" edges to ScanJob entities.
func (aiuo *ASNInfoUpdateOne) RemoveScan(s ...*ScanJob) *ASNInfoUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return aiuo.RemoveScanIDs(ids...)
}

// ClearIpaddress clears all "ipaddress" edges to the IPAddress entity.
func (aiuo *ASNInfoUpdateOne) ClearIpaddress() *ASNInfoUpdateOne {
	aiuo.mutation.ClearIpaddress()
	return aiuo
}

// RemoveIpaddresIDs removes the "ipaddress" edge to IPAddress entities by IDs.
func (aiuo *ASNInfoUpdateOne) RemoveIpaddresIDs(ids ...int) *ASNInfoUpdateOne {
	aiuo.mutation.RemoveIpaddresIDs(ids...)
	return aiuo
}

// RemoveIpaddress removes "ipaddress" edges to IPAddress entities.
func (aiuo *ASNInfoUpdateOne) RemoveIpaddress(i ...*IPAddress) *ASNInfoUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return aiuo.RemoveIpaddresIDs(ids...)
}

// ClearRegistrar clears all "registrar" edges to the Registrar entity.
func (aiuo *ASNInfoUpdateOne) ClearRegistrar() *ASNInfoUpdateOne {
	aiuo.mutation.ClearRegistrar()
	return aiuo
}

// RemoveRegistrarIDs removes the "registrar" edge to Registrar entities by IDs.
func (aiuo *ASNInfoUpdateOne) RemoveRegistrarIDs(ids ...int) *ASNInfoUpdateOne {
	aiuo.mutation.RemoveRegistrarIDs(ids...)
	return aiuo
}

// RemoveRegistrar removes "registrar" edges to Registrar entities.
func (aiuo *ASNInfoUpdateOne) RemoveRegistrar(r ...*Registrar) *ASNInfoUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return aiuo.RemoveRegistrarIDs(ids...)
}

// ClearWhois clears all "whois" edges to the Whois entity.
func (aiuo *ASNInfoUpdateOne) ClearWhois() *ASNInfoUpdateOne {
	aiuo.mutation.ClearWhois()
	return aiuo
}

// RemoveWhoiIDs removes the "whois" edge to Whois entities by IDs.
func (aiuo *ASNInfoUpdateOne) RemoveWhoiIDs(ids ...int) *ASNInfoUpdateOne {
	aiuo.mutation.RemoveWhoiIDs(ids...)
	return aiuo
}

// RemoveWhois removes "whois" edges to Whois entities.
func (aiuo *ASNInfoUpdateOne) RemoveWhois(w ...*Whois) *ASNInfoUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return aiuo.RemoveWhoiIDs(ids...)
}

// Where appends a list predicates to the ASNInfoUpdate builder.
func (aiuo *ASNInfoUpdateOne) Where(ps ...predicate.ASNInfo) *ASNInfoUpdateOne {
	aiuo.mutation.Where(ps...)
	return aiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aiuo *ASNInfoUpdateOne) Select(field string, fields ...string) *ASNInfoUpdateOne {
	aiuo.fields = append([]string{field}, fields...)
	return aiuo
}

// Save executes the query and returns the updated ASNInfo entity.
func (aiuo *ASNInfoUpdateOne) Save(ctx context.Context) (*ASNInfo, error) {
	return withHooks(ctx, aiuo.sqlSave, aiuo.mutation, aiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aiuo *ASNInfoUpdateOne) SaveX(ctx context.Context) *ASNInfo {
	node, err := aiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aiuo *ASNInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := aiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aiuo *ASNInfoUpdateOne) ExecX(ctx context.Context) {
	if err := aiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aiuo *ASNInfoUpdateOne) sqlSave(ctx context.Context) (_node *ASNInfo, err error) {
	_spec := sqlgraph.NewUpdateSpec(asninfo.Table, asninfo.Columns, sqlgraph.NewFieldSpec(asninfo.FieldID, field.TypeInt))
	id, ok := aiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model_ent: missing "ASNInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asninfo.FieldID)
		for _, f := range fields {
			if !asninfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model_ent: invalid field %q for query", f)}
			}
			if f != asninfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aiuo.mutation.Asn(); ok {
		_spec.SetField(asninfo.FieldAsn, field.TypeInt, value)
	}
	if value, ok := aiuo.mutation.AddedAsn(); ok {
		_spec.AddField(asninfo.FieldAsn, field.TypeInt, value)
	}
	if value, ok := aiuo.mutation.Country(); ok {
		_spec.SetField(asninfo.FieldCountry, field.TypeString, value)
	}
	if value, ok := aiuo.mutation.Registry(); ok {
		_spec.SetField(asninfo.FieldRegistry, field.TypeString, value)
	}
	if aiuo.mutation.ScanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.ScanTable,
			Columns: asninfo.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.RemovedScanIDs(); len(nodes) > 0 && !aiuo.mutation.ScanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.ScanTable,
			Columns: asninfo.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.ScanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.ScanTable,
			Columns: asninfo.ScanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(scanjob.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aiuo.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.IpaddressTable,
			Columns: asninfo.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.RemovedIpaddressIDs(); len(nodes) > 0 && !aiuo.mutation.IpaddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.IpaddressTable,
			Columns: asninfo.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.IpaddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.IpaddressTable,
			Columns: asninfo.IpaddressPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aiuo.mutation.RegistrarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.RegistrarTable,
			Columns: asninfo.RegistrarPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(registrar.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.RemovedRegistrarIDs(); len(nodes) > 0 && !aiuo.mutation.RegistrarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.RegistrarTable,
			Columns: asninfo.RegistrarPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(registrar.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.RegistrarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.RegistrarTable,
			Columns: asninfo.RegistrarPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(registrar.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aiuo.mutation.WhoisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.WhoisTable,
			Columns: asninfo.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.RemovedWhoisIDs(); len(nodes) > 0 && !aiuo.mutation.WhoisCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.WhoisTable,
			Columns: asninfo.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aiuo.mutation.WhoisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   asninfo.WhoisTable,
			Columns: asninfo.WhoisPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(whois.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ASNInfo{config: aiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asninfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aiuo.mutation.done = true
	return _node, nil
}
