package session

import "log"

func (s *Session) Begin() (err error) {
	log.Println("transaction begin")

	if s.tx, err = s.db.Begin(); err != nil {
		log.Println("err=", err)
		return
	}
	return
}

func (s *Session) Commit() (err error) {
	log.Println("transaction commit")

	if err = s.tx.Commit(); err != nil {
		log.Println("err=", err)
		return
	}
	return
}

func (s *Session) Rollback() (err error) {
	log.Println("transaction rollback")

	if err = s.tx.Rollback(); err != nil {
		log.Println("err=", err)
		return
	}
	return
}
