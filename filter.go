package gocassa

import (
	"bytes"
	"fmt"
	"strconv"
)

type filter struct {
	t  t
	rs []Relation
}

func generateWhere(rs []Relation) (string, []interface{}) {
	var (
		vals []interface{}
		buf  = new(bytes.Buffer)
	)

	if len(rs) > 0 {
		buf.WriteString(" WHERE ")
		for i, r := range rs {
			if i > 0 {
				buf.WriteString(" AND ")
			}
			s, v := r.cql()
			buf.WriteString(s)
			vals = append(vals, v...)
		}
	}
	return buf.String(), vals
}

// UPDATE keyspace.Movies SET col1 = val1, col2 = val2
func updateStatement(kn, cfName string, fields map[string]interface{}, opts Options) (string, []interface{}) {
	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("UPDATE %s.%s ", kn, cfName))

	// Apply options
	if opts.TTL != 0 {
		buf.WriteString("USING TTL ")
		buf.WriteString(strconv.FormatFloat(opts.TTL.Seconds(), 'f', 0, 64))
		buf.WriteRune(' ')
	}

	buf.WriteString("SET ")
	i := 0
	ret := []interface{}{}
	for k, v := range fields {
		if i > 0 {
			buf.WriteString(", ")
		}
		if mod, ok := v.(Modifier); ok {
			stmt, vals := mod.cql(k)
			buf.WriteString(stmt)
			ret = append(ret, vals...)
		} else {
			buf.WriteString(k + " = ?")
			ret = append(ret, v)
		}
		i++
	}

	return buf.String(), ret
}

func (f filter) UpdateWithOptions(m map[string]interface{}, opts Options) Op {
	str, wvals := generateWhere(f.rs)
	stmt, uvals := updateStatement(f.t.keySpace.name, f.t.Name(), m, opts)
	vs := append(uvals, wvals...)
	if f.t.keySpace.debugMode {
		fmt.Println(stmt+" "+str, vs)
	}
	return newWriteOp(f.t.keySpace.qe, stmt+str, vs)
}

// Update does a partial update on the filter.
func (f filter) Update(m map[string]interface{}) Op {
	return f.UpdateWithOptions(m, f.t.options)
}

func (f filter) Delete() Op {
	str, vals := generateWhere(f.rs)
	stmt := fmt.Sprintf("DELETE FROM %s.%s%s", f.t.keySpace.name, f.t.Name(), str)
	if f.t.keySpace.debugMode {
		fmt.Println(stmt, vals)
	}
	return newWriteOp(f.t.keySpace.qe, stmt, vals)
}

// Query returns the query from the filter so you can read the rows matching the filter.
func (f filter) Query() Query {
	return &query{
		f:       f,
		options: f.t.options,
	}
}
