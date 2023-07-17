package healthchek

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo Repository) Application() string {
	return "pong"
}
