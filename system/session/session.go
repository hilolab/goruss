package session

type Store interface {
	Init(id string) SessionInterface
	GC()
	//Option() *Option
}

type SessionInterface interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Remove(key string)
	Flush()
	Save() error
}

type Session struct {
	id       string
	provider SessionInterface
}

func NewSession(id string, store Store) *Session {
	return &Session{
		id:       id,
		provider: store.Init(id),
	}
}

func (s *Session) Id() string {
	return s.id
}

func (s *Session) Get(key string) interface{} {
	return s.provider.Get(key)
}

func (s *Session) Set(key string, value interface{}) {
	s.provider.Set(key, value)
}

func (s *Session) Remove(key string) {
	s.provider.Remove(key)
}

func (s *Session) Flush() {
	s.provider.Flush()
}

func (s *Session) Save() error {
	return s.provider.Save()
}
