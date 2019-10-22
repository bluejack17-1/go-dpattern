package creational

import (
	"fmt"
	"strings"
)

type Query struct {
	columns    []string
	table      string
	conditions []Condition
}

type Condition struct {
	Connector string
	Operator  string
	Column    string
	Value     interface{}
}

func (q *Query) From(table string) *Query {
	q.table = table

	return q
}

func (q *Query) Select(columns ...string) *Query {
	if len(columns) < 1 {
		q.columns = append(q.columns, "*")
	} else {
		q.columns = append(q.columns, columns...)
	}

	return q
}

func (q *Query) Where(column string, operator string, value interface{}) *Query {
	q.conditions = append(q.conditions, Condition{
		Connector: "AND",
		Operator:  operator,
		Column:    column,
		Value:     value,
	})

	return q
}

func (q *Query) OrWhere(column string, operator string, value interface{}) *Query {
	q.conditions = append(q.conditions, Condition{
		Connector: "OR",
		Operator:  operator,
		Column:    column,
		Value:     value,
	})

	return q
}

func (q *Query) Build() string {
	var query string = ""

	query = "SELECT"

	query = fmt.Sprintf("%s %s", query, strings.Join(q.columns, ","))

	query = fmt.Sprintf("%s FROM %s", query, q.table)

	if len(q.conditions) > 0 {
		query = fmt.Sprintf("%s WHERE", query)

		for i, cond := range q.conditions {
			if i == 0 {
				query = fmt.Sprintf("%s %s %s '%v'", query, cond.Column, cond.Operator, cond.Value)
			} else {
				query = fmt.Sprintf("%s %s %s %s '%v'", query, cond.Connector, cond.Column, cond.Operator, cond.Value)
			}
		}
	}

	return query
}
