package qm

import (
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// QueryMod modifies a query object.
type QueryMod qm.QueryMod

// Where allows you to specify a where clause for your statement
func Where(clause string, args ...interface{}) qm.QueryMod {
	if clause == "" {
		clause = "true"
	}
	return qm.Where(clause, args...)
}

// And allows you to specify a where clause separated by an AND for your statement
// And is a duplicate of the Where function, but allows for more natural looking
// query mod chains, for example: (Where("a=?"), And("b=?"), Or("c=?")))
func And(clause string, args ...interface{}) qm.QueryMod {
	if clause == "" {
		clause = "true"
	}
	return qm.And(clause, args...)
}

// Or allows you to specify a where clause separated by an OR for your statement
func Or(clause string, args ...interface{}) qm.QueryMod {
	if clause == "" {
		clause = "false"
	}
	return qm.Or(clause, args...)
}

// WhereIn allows you to specify a "x IN (set)" clause for your where statement
// Example clauses: "column in ?", "(column1,column2) in ?"
func WhereIn(clause string, args ...interface{}) qm.QueryMod {
	if len(args) == 0 {
		clause = "false"
	}
	return qm.WhereIn(clause, args...)
}

// AndIn allows you to specify a "x IN (set)" clause separated by an AndIn
// for your where statement. AndIn is a duplicate of the WhereIn function, but
// allows for more natural looking query mod chains, for example:
// (WhereIn("column1 in ?"), AndIn("column2 in ?"), OrIn("column3 in ?"))
func AndIn(clause string, args ...interface{}) qm.QueryMod {
	if len(args) == 0 {
		clause = "false"
	}
	return qm.AndIn(clause, args...)
}

// OrIn allows you to specify an IN clause separated by
// an OR for your where statement
func OrIn(clause string, args ...interface{}) qm.QueryMod {
	if len(args) == 0 {
		clause = "false"
	}
	return qm.OrIn(clause, args...)
}

// GroupBy allows you to specify a group by clause for your statement
func GroupBy(clause string) qm.QueryMod {
	if clause == "" {
		clause = "id"
	}
	return qm.GroupBy(clause)
}

// OrderBy allows you to specify a order by clause for your statement
func OrderBy(clause string) qm.QueryMod {
	if clause == "" {
		clause = "id"
	}
	return qm.OrderBy(clause)
}

// Having allows you to specify a having clause for your statement
func Having(clause string, args ...interface{}) qm.QueryMod {
	if len(args) == 0 {
		clause = "true"
	}
	return qm.Having(clause, args...)
}

// From allows to specify the table for your statement
func From(from string) qm.QueryMod {
	return qm.From(from)
}

// Limit the number of returned rows
func Limit(limit int) qm.QueryMod {
	return qm.Limit(limit)
}

// Offset into the results
func Offset(offset int) qm.QueryMod {
	return qm.Offset(offset)
}

// For inserts a concurrency locking clause at the end of your statement
func For(clause string) qm.QueryMod {
	return qm.For(clause)
}

// InnerJoin on another table
func InnerJoin(clause string, args ...interface{}) qm.QueryMod {
	return qm.InnerJoin(clause, args...)
}
