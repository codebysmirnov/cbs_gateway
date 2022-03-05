package config

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

// Endpoint configures on service endpoint.
type Endpoint struct {
	From   string
	To     string
	Method string
}

// Service presents service settings
type Service struct {
	Name      string
	Host      string
	Endpoints []Endpoint
}

func LoadServices() ([]Service, error) {
	dir := "./configs"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	services := make([]Service, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		content, err := ioutil.ReadFile(path.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		var service Service
		err = json.Unmarshal(content, &service)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return services, err
}
