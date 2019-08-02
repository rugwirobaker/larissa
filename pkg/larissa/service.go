package larissa

// Service descirbes the user acceible interface to end users
type Service interface {
	Put(path string, content []byte) error
	Get(path string) (*Object, error)
	Del(path string) error
	Exists(path string) bool
}

var _ (Service) = (*service)(nil)

type service struct {
	backend Backend
}

// New creates a new larissa service
func New(backend Backend) Service {
	return &service{backend}
}

func (svc *service) Put(path string, content []byte) error {
	return svc.backend.Put(path, content)
}

func (svc *service) Get(path string) (*Object, error) {
	return svc.backend.Get(path)
}

func (svc *service) Del(path string) error {
	return svc.backend.Del(path)
}

func (svc *service) Exists(path string) bool {
	return svc.backend.Exists(path)
}
