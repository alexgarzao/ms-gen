package main

type Service struct {
	DB
}

func NewService() *Service {
	service := new(Service)
	service.DB = DB{}

	return service
}

func (s *Service) StartDB() {
	s.DB.InitDB()
	s.DB.InitSchema()
}
