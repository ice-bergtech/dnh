// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/ice-bergtech/dnh/src/internal/model_ent"
)

// The ASNInfoFunc type is an adapter to allow the use of ordinary
// function as ASNInfo mutator.
type ASNInfoFunc func(context.Context, *model_ent.ASNInfoMutation) (model_ent.Value, error)

// Mutate calls f(ctx, m).
func (f ASNInfoFunc) Mutate(ctx context.Context, m model_ent.Mutation) (model_ent.Value, error) {
	if mv, ok := m.(*model_ent.ASNInfoMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model_ent.ASNInfoMutation", m)
}

// The DNSEntryFunc type is an adapter to allow the use of ordinary
// function as DNSEntry mutator.
type DNSEntryFunc func(context.Context, *model_ent.DNSEntryMutation) (model_ent.Value, error)

// Mutate calls f(ctx, m).
func (f DNSEntryFunc) Mutate(ctx context.Context, m model_ent.Mutation) (model_ent.Value, error) {
	if mv, ok := m.(*model_ent.DNSEntryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model_ent.DNSEntryMutation", m)
}

// The DomainFunc type is an adapter to allow the use of ordinary
// function as Domain mutator.
type DomainFunc func(context.Context, *model_ent.DomainMutation) (model_ent.Value, error)

// Mutate calls f(ctx, m).
func (f DomainFunc) Mutate(ctx context.Context, m model_ent.Mutation) (model_ent.Value, error) {
	if mv, ok := m.(*model_ent.DomainMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model_ent.DomainMutation", m)
}

// The ExampleFunc type is an adapter to allow the use of ordinary
// function as Example mutator.
type ExampleFunc func(context.Context, *model_ent.ExampleMutation) (model_ent.Value, error)

// Mutate calls f(ctx, m).
func (f ExampleFunc) Mutate(ctx context.Context, m model_ent.Mutation) (model_ent.Value, error) {
	if mv, ok := m.(*model_ent.ExampleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model_ent.ExampleMutation", m)
}

// The IPAddressFunc type is an adapter to allow the use of ordinary
// function as IPAddress mutator.
type IPAddressFunc func(context.Context, *model_ent.IPAddressMutation) (model_ent.Value, error)

// Mutate calls f(ctx, m).
func (f IPAddressFunc) Mutate(ctx context.Context, m model_ent.Mutation) (model_ent.Value, error) {
	if mv, ok := m.(*model_ent.IPAddressMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model_ent.IPAddressMutation", m)
}

// The NameserverFunc type is an adapter to allow the use of ordinary
// function as Nameserver mutator.
type NameserverFunc func(context.Context, *model_ent.NameserverMutation) (model_ent.Value, error)

// Mutate calls f(ctx, m).
func (f NameserverFunc) Mutate(ctx context.Context, m model_ent.Mutation) (model_ent.Value, error) {
	if mv, ok := m.(*model_ent.NameserverMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model_ent.NameserverMutation", m)
}

// The PathFunc type is an adapter to allow the use of ordinary
// function as Path mutator.
type PathFunc func(context.Context, *model_ent.PathMutation) (model_ent.Value, error)

// Mutate calls f(ctx, m).
func (f PathFunc) Mutate(ctx context.Context, m model_ent.Mutation) (model_ent.Value, error) {
	if mv, ok := m.(*model_ent.PathMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model_ent.PathMutation", m)
}

// The RegistrarFunc type is an adapter to allow the use of ordinary
// function as Registrar mutator.
type RegistrarFunc func(context.Context, *model_ent.RegistrarMutation) (model_ent.Value, error)

// Mutate calls f(ctx, m).
func (f RegistrarFunc) Mutate(ctx context.Context, m model_ent.Mutation) (model_ent.Value, error) {
	if mv, ok := m.(*model_ent.RegistrarMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model_ent.RegistrarMutation", m)
}

// The ScanFunc type is an adapter to allow the use of ordinary
// function as Scan mutator.
type ScanFunc func(context.Context, *model_ent.ScanMutation) (model_ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScanFunc) Mutate(ctx context.Context, m model_ent.Mutation) (model_ent.Value, error) {
	if mv, ok := m.(*model_ent.ScanMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model_ent.ScanMutation", m)
}

// The ScanJobFunc type is an adapter to allow the use of ordinary
// function as ScanJob mutator.
type ScanJobFunc func(context.Context, *model_ent.ScanJobMutation) (model_ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScanJobFunc) Mutate(ctx context.Context, m model_ent.Mutation) (model_ent.Value, error) {
	if mv, ok := m.(*model_ent.ScanJobMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model_ent.ScanJobMutation", m)
}

// The WhoisFunc type is an adapter to allow the use of ordinary
// function as Whois mutator.
type WhoisFunc func(context.Context, *model_ent.WhoisMutation) (model_ent.Value, error)

// Mutate calls f(ctx, m).
func (f WhoisFunc) Mutate(ctx context.Context, m model_ent.Mutation) (model_ent.Value, error) {
	if mv, ok := m.(*model_ent.WhoisMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *model_ent.WhoisMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, model_ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m model_ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m model_ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m model_ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op model_ent.Op) Condition {
	return func(_ context.Context, m model_ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model_ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model_ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m model_ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk model_ent.Hook, cond Condition) model_ent.Hook {
	return func(next model_ent.Mutator) model_ent.Mutator {
		return model_ent.MutateFunc(func(ctx context.Context, m model_ent.Mutation) (model_ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, model_ent.Delete|model_ent.Create)
func On(hk model_ent.Hook, op model_ent.Op) model_ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, model_ent.Update|model_ent.UpdateOne)
func Unless(hk model_ent.Hook, op model_ent.Op) model_ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) model_ent.Hook {
	return func(model_ent.Mutator) model_ent.Mutator {
		return model_ent.MutateFunc(func(context.Context, model_ent.Mutation) (model_ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []model_ent.Hook {
//		return []model_ent.Hook{
//			Reject(model_ent.Delete|model_ent.Update),
//		}
//	}
func Reject(op model_ent.Op) model_ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []model_ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...model_ent.Hook) Chain {
	return Chain{append([]model_ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() model_ent.Hook {
	return func(mutator model_ent.Mutator) model_ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...model_ent.Hook) Chain {
	newHooks := make([]model_ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
