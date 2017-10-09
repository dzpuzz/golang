package session

type Store interface {
	Get(id string) Session
	Remove(sess Session)
	Add(sess Session)
	Close()
}
