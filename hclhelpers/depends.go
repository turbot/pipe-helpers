package hclhelpers

import (
	"slices"

	"github.com/hashicorp/hcl/v2"
)

func ExpressionToDepends(expr hcl.Expression, validDependsOnTypes []string) ([]string, hcl.Diagnostics) {
	var dependsOn []string
	traversals := expr.Variables()
	for _, t := range traversals {
		parts := TraversalAsStringSlice(t)
		if len(parts) >= 3 {
			if slices.Contains(validDependsOnTypes, parts[0]) {
				if len(parts) >= 3 {
					dependsOn = append(dependsOn, parts[1]+"."+parts[2])
				}
			}
		}
	}

	// We have no diags to return here for now, but add this for future proofing
	return dependsOn, nil
}
