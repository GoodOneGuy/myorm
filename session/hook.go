package session

type IBeforeQuery interface {
	BeforeQuery(s *Session) error
}

type IAfterQuery interface {
	AfterQuery(s *Session) error
}
