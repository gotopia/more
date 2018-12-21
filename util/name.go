package util

import (
	"fmt"
	"regexp"
	"strconv"
)

// ExtractID from resource name.
func ExtractID(name string, collection string) uint {
	str := fmt.Sprintf("%v/([^/]+)", collection)
	var reg = regexp.MustCompile(str)
	subs := reg.FindStringSubmatch(name)
	if len(subs) != 2 {
		return 0
	}
	idstr := subs[1]
	id, _ := strconv.Atoi(idstr)
	return uint(id)
}

// ExtractIDs from resource names.
func ExtractIDs(names []string, collection string) []interface{} {
	var ids []interface{}
	for _, name := range names {
		ids = append(ids, ExtractID(name, collection))
	}
	return ids
}
