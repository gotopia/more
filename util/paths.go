package util

import (
	mapset "github.com/deckarep/golang-set"
)

// Paths implements the some util functions of the update mask.
type Paths []string

// Has indicates whether the path is included in the update mask.
func (ps Paths) Has(path string) bool {
	for _, p := range ps {
		if p == path {
			return true
		}
	}
	return false
}

// Permit the fields in paths and returns a new paths.
func (ps Paths) Permit(fields ...string) Paths {
	pms := mapset.NewSetFromSlice(InterfaceSlice(ps))
	fms := mapset.NewSetFromSlice(InterfaceSlice(fields))
	ms := pms.Intersect(fms)
	s := ms.ToSlice()
	l := len(s)
	paths := make(Paths, l)
	for i := 0; i < l; i++ {
		paths[i] = s[i].(string)
	}
	return paths
}

// Append the fields to the paths.
func (ps Paths) Append(fields ...string) Paths {
	ps = append(ps, fields...)
	return ps
}
