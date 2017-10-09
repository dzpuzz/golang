package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"sync"
	"time"
)

type Session interface {
	ID() string

	New() bool

	CAttr(name string) interface{}

	Attr(name string) interface{}

	SetAttr(name string, value interface{})

	Attrs() map[string]interface{}

	Created() time.Time
	Accessed() time.Time

	Timeout() time.Duration
	Mutex() *sync.RWMutex
	Access()
}
type sessionImpl struct {
	IDF       string
	CAttrsF   map[string]interface{}
	AttrsF    map[string]interface{}
	CreatedF  time.Time
	AccessedF time.Time
	TimeoutF  time.Duration
	mux       *sync.RWMutex
}

type SessionOptions struct {
	CAttrs   map[string]interface{}
	Attrs    map[string]interface{}
	Timeout  time.Duration
	IDLength int
}

func NewSessionOption(s *SessionOptions) Session {
	now := time.Now()
	iDLength := s.IDLength
	if iDLength <= 0 {
		iDLength = 18
	}
	timeout := s.Timeout
	if timeout == 0 {
		timeout = time.Minute * 30
	}
	sess := sessionImpl{
		IDF:       genID(iDLength),
		CreatedF:  now,
		AccessedF: now,
		AttrsF:    make(map[string]interface{}),
		TimeoutF:  timeout,
		mux:       &sync.RWMutex{},
	}
	if len(s.CAttrs) > 0 {
		sess.CAttrsF = make(map[string]interface{}, len(s.CAttrs))
		for k, v := range s.CAttrs {
			sess.CAttrsF[k] = v
		}
	}

	for k, v := range s.Attrs {
		sess.AttrsF[k] = v
	}
	return &sess
}

// 生成随机sessionID
func genID(length int) string {
	r := make([]byte, length)
	io.Reader.Read(rand.Reader, r)
	return base64.URLEncoding.EncodeToString(r)
}

func (s *sessionImpl) Timeout() time.Duration {
	return s.TimeoutF
}
func (s *sessionImpl) ID() string {
	return s.IDF
}
func (s *sessionImpl) CAttr(name string) interface{} {
	return s.CAttrsF[name]
}
func (s *sessionImpl) Attr(name string) interface{} {
	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.AttrsF[name]
}
func (s *sessionImpl) SetAttr(name string, value interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if value == nil {
		delete(s.AttrsF, name)
	} else {
		s.AttrsF[name] = value
	}
}
func (s *sessionImpl) Access() {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.AccessedF = time.Now()
}
func (s *sessionImpl) Accessed() time.Time {
	s.mux.Lock()
	defer s.mux.Unlock()

	return s.AccessedF
}

func (s *sessionImpl) Mutex() *sync.RWMutex {
	return s.mux
}
func (s *sessionImpl) Created() time.Time {
	return s.CreatedF
}
func (s *sessionImpl) Attrs() map[string]interface{} {
	s.mux.RLock()
	defer s.mux.RUnlock()

	m := make(map[string]interface{}, len(s.AttrsF))
	for k, v := range s.AttrsF {
		m[k] = v
	}
	return m
}
func (s *sessionImpl) New() bool {
	return s.CreatedF == s.AccessedF
}
