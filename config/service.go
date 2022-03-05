package config

// Auth presents auth type
type Auth map[string]string

// Service presents service settings
type Service struct {
	Name string
	Host string
	Auth Auth
}
