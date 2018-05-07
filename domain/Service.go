package domain

//Service : ...
type Service struct {
	price money
}

//NewService : constructor for Service
func NewService() *Service {
	return &Service{}
}
