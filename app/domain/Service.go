package domain

//Service : ...
type Service struct {
	title string
	price int
}

//NewService : constructor for Service
func NewService() *Service {
	return &Service{}
}
