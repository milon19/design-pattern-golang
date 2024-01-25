package query

import (
	"strings"
)

type SQLQueryBuilder struct {
	selectFields string
	fromTable    string
	whereClause  string
}

func NewSQLQueryBuilder() *SQLQueryBuilder {
	return &SQLQueryBuilder{}
}

func (qb *SQLQueryBuilder) Select(fields string) *SQLQueryBuilder {
	qb.selectFields = fields
	return qb
}

func (qb *SQLQueryBuilder) From(table string) *SQLQueryBuilder {
	qb.fromTable = table
	return qb
}

func (qb *SQLQueryBuilder) Where(condition string) *SQLQueryBuilder {
	qb.whereClause = condition
	return qb
}

func (qb *SQLQueryBuilder) Build() string {
	var query strings.Builder
	query.WriteString("SELECT ")
	query.WriteString(qb.selectFields)
	query.WriteString(" FROM ")
	query.WriteString(qb.fromTable)
	if qb.whereClause != "" {
		query.WriteString(" WHERE ")
		query.WriteString(qb.whereClause)
	}
	return query.String()
}
