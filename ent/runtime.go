// Code generated by ent, DO NOT EDIT.

package ent

import (
	"atiinc/entgo-sqlite/ent/example"
	"atiinc/entgo-sqlite/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	exampleMixin := schema.Example{}.Mixin()
	exampleMixinFields0 := exampleMixin[0].Fields()
	_ = exampleMixinFields0
	exampleFields := schema.Example{}.Fields()
	_ = exampleFields
	// exampleDescCreatedAt is the schema descriptor for CreatedAt field.
	exampleDescCreatedAt := exampleMixinFields0[1].Descriptor()
	// example.DefaultCreatedAt holds the default value on creation for the CreatedAt field.
	example.DefaultCreatedAt = exampleDescCreatedAt.Default.(func() time.Time)
	// exampleDescID is the schema descriptor for id field.
	exampleDescID := exampleMixinFields0[0].Descriptor()
	// example.DefaultID holds the default value on creation for the id field.
	example.DefaultID = exampleDescID.Default.(func() uuid.UUID)
}
