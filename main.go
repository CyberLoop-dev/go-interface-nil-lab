package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"reflect"
)

type Greeter interface { Greet() string }
type Person struct { Name string }

func (p *Person) Greet() string {
	if p == nil { return "method received a nil *Person" }
	return "Hello, " + p.Name
}

type Result struct {
	NilInterfaceIsNil bool   `json:"nil_interface_is_nil"`
	TypedNilIsNil     bool   `json:"typed_nil_is_nil"`
	DynamicType       string `json:"dynamic_type"`
	NilPointerValue   bool   `json:"nil_pointer_value"`
	MethodResult      string `json:"method_result"`
}

func inspect() Result {
	var empty Greeter
	var pointer *Person
	var typed Greeter = pointer
	value := reflect.ValueOf(typed)
	return Result{
		NilInterfaceIsNil: empty == nil,
		TypedNilIsNil: typed == nil,
		DynamicType: reflect.TypeOf(typed).String(),
		NilPointerValue: value.IsNil(),
		MethodResult: typed.Greet(),
	}
}

func main() {
	asJSON := flag.Bool("json", false, "print the result as JSON")
	flag.Parse()
	result := inspect()
	if *asJSON {
		_ = json.NewEncoder(flag.CommandLine.Output()).Encode(result)
		return
	}
	fmt.Println("Go interface nil lab")
	fmt.Println("--------------------")
	fmt.Printf("nil interface == nil: %t\n", result.NilInterfaceIsNil)
	fmt.Printf("interface holding (*Person)(nil) == nil: %t\n", result.TypedNilIsNil)
	fmt.Printf("dynamic type stored in interface: %s\n", result.DynamicType)
	fmt.Printf("stored pointer itself is nil: %t\n", result.NilPointerValue)
	fmt.Printf("calling the nil-safe method: %s\n", result.MethodResult)
	fmt.Println("\nAn interface is nil only when both its dynamic type and value are absent.")
}
