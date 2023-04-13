package bootstrap

import "os"

type Service struct {
	Name     string
	Version  string
	Id       string
	Metadata map[string]string
}

func NewServiceInfo(name, version, id string) *Service {
	if id == "" {
		id, _ = os.Hostname()
	}
	return &Service{
		Name:     name,
		Version:  version,
		Id:       id,
		Metadata: map[string]string{},
	}
}

func (s *Service) GetInstanceId() string {
	return s.Id + "." + s.Name
}

func (s *Service) SetMataData(k, v string) {
	s.Metadata[k] = v
}
