package main

import "testing"

func TestTypedNilDiffersFromNilInterface(t *testing.T) {
	result := inspect()
	if !result.NilInterfaceIsNil { t.Fatal("zero-value interface should be nil") }
	if result.TypedNilIsNil { t.Fatal("interface with a typed nil pointer must not be nil") }
	if !result.NilPointerValue { t.Fatal("stored pointer should remain nil") }
	if result.DynamicType != "*main.Person" { t.Fatalf("unexpected dynamic type: %s", result.DynamicType) }
}
