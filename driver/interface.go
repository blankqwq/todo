package driver

type RepositoryDriver interface {
	Select() (map[int]string, error)
	Insert(data interface{}) error
	Update(id int, data interface{}) error
	Delete(id int) error
	Find(id int) (string, error)
	Init() error
	Free() error
	Logout() error
	Login(username, password string) error
}
