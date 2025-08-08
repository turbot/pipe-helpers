package hclhelpers

import (
	"reflect"

	"github.com/hashicorp/hcl/v2"
	"github.com/turbot/pipe-helpers/helpers"
)

func GetNestedStructValsRecursive(val any) ([]any, hcl.Diagnostics) {
	nested, diags := GetNestedStructVals(val)
	res := nested

	for _, n := range nested {
		nestedVals, moreDiags := GetNestedStructValsRecursive(n)
		diags = append(diags, moreDiags...)
		res = append(res, nestedVals...)
	}
	return res, diags

}

// GetNestedStructVals return a slice of any nested structs within val
func GetNestedStructVals(val any) (_ []any, diags hcl.Diagnostics) {
	defer func() {
		if r := recover(); r != nil {
			if r := recover(); r != nil {
				diags = append(diags, &hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  "unexpected error in resolveReferences",
					Detail:   helpers.ToError(r).Error()})
			}
		}
	}()

	rv := reflect.ValueOf(val)
	for rv.Type().Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	ty := rv.Type()
	if ty.Kind() != reflect.Struct {
		return nil, nil
	}
	ct := ty.NumField()
	var res []any
	for i := 0; i < ct; i++ {
		field := ty.Field(i)
		fieldVal := rv.Field(i)
		if field.Anonymous && fieldVal.Kind() == reflect.Struct {
			res = append(res, fieldVal.Addr().Interface())
		}
	}
	return res, nil
}
