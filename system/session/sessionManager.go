package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Option struct {
	Path     string
	Domain   string
	MaxAge   int
	Secure   bool
	HttpOnly bool
}

type SessionManager struct {
	Name    string
	manager map[string]*Session
	mutex   sync.Mutex
	option  *Option
	store   Store
}

func NewSessionManager(store Store) *SessionManager {
	s := &SessionManager{
		manager: make(map[string]*Session),
		Name:    "sessid",
		option: &Option{
			Path:     "/",
			MaxAge:   1200,
			HttpOnly: true,
		},
		store: store,
	}
	s.GC()

	return s
}

func (s *SessionManager) Start(rw http.ResponseWriter, req *http.Request) *Session {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var sess *Session

	cookie, err := req.Cookie(s.Name)

	if err != nil || cookie.Value == "" {
		sessId := s.SessionId()
		sess = NewSession(sessId, s.store)
		cookie := &http.Cookie{
			Name:     s.Name,
			Value:    sessId,
			Path:     s.option.Path,
			HttpOnly: s.option.HttpOnly,
			MaxAge:   s.option.MaxAge,
		}
		s.manager[sessId] = sess

		http.SetCookie(rw, cookie)
	} else {
		//坑啊
		var ok bool
		sess, ok = s.manager[cookie.Value]

		if !ok {
			sess = NewSession(cookie.Value, s.store)
		}
	}

	return sess
}

func (s *SessionManager) SessionId() string {
	buf := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		now := time.Now().UnixNano()
		strNow := strconv.FormatInt(now, 10)
		return base64.URLEncoding.EncodeToString([]byte(strNow))
	}

	return base64.URLEncoding.EncodeToString(buf)
}

func (s *SessionManager) GC() {
	go func() {
		for {
			//fmt.Println("tick...")
			s.store.GC()
			tc := time.Tick(2 * time.Second)
			<-tc
		}
	}()
}
