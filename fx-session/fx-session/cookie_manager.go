package session

import (
	"net/http"
	"time"
)

type CookieManager struct {
	store            Store // Backing Store
	sessIDCookieName string
	cookieSecure     bool
	cookieMaxAgeSec  int
	cookiePath       string
}

type CookieMangerOption struct {
	SessIDCookieName string
	AllowHTTP        bool
	CookieMaxAge     time.Duration
	CookiePath       string
}

func NewCookieManagerOptions(store Store, o *CookieMangerOption) Manager {
	m := &CookieManager{
		store:            store,
		cookieSecure:     !o.AllowHTTP, //secure true http不发送cookie
		sessIDCookieName: o.SessIDCookieName,
		cookiePath:       o.CookiePath,
	}
	return m
}
func (m *CookieManager) Get(r *http.Request) Session {
	c, err := r.Cookie(m.sessIDCookieName)
	if err != nil {
		return nil
	}
	return m.store.Get(c.Value)
}

func (m *CookieManager) Add(sess Session, w http.ResponseWriter) {
	c := http.Cookie{
		Name:     m.sessIDCookieName,
		Value:    sess.ID(),
		Path:     m.cookiePath,
		HttpOnly: true,
		Secure:   m.cookieSecure,
		MaxAge:   m.cookieMaxAgeSec,
	}
	http.SetCookie(w, &c)

	m.store.Add(sess)
}

func (m *CookieManager) Remove(sess Session, w http.ResponseWriter) {
	c := http.Cookie{
		Name:     m.sessIDCookieName,
		Value:    sess.ID(),
		Path:     m.cookiePath,
		HttpOnly: true,
		Secure:   m.cookieSecure,
		MaxAge:   -1,
	}
	http.SetCookie(w, &c)

	m.store.Remove(sess)
}
func (m *CookieManager) CookieMaxAgeSec() int {
	return m.cookieMaxAgeSec
}

func (m *CookieManager) Close() {
	m.store.Close()
}
