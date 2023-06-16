package session

import (
	"database/sql"
	"github.com/GoodOneGuy/myorm/clause"
	"github.com/GoodOneGuy/myorm/dialect"
	"github.com/GoodOneGuy/myorm/schema"
	"log"
	"strings"
)

type Session struct {
	db       *sql.DB
	sql      strings.Builder
	sqlVars  []interface{}
	dialect  dialect.Dialect
	refTable *schema.Schema
	clause   clause.Clause
	tx       *sql.Tx
}

// CommonDB is a minimal function set of db
type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

func New(db *sql.DB, _dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: _dialect,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = s.sqlVars[:0]
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

func (s *Session) Exec() (sql.Result, error) {
	defer s.Clear()
	log.Println(s.sql.String(), s.sqlVars)
	result, err := s.DB().Exec(s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Println(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session) QueryRows() (*sql.Rows, error) {
	defer s.Clear()
	log.Println(s.sql.String(), s.sqlVars)
	result, err := s.DB().Query(s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
