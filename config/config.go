package config

type Config struct {
	Enviroments []string `json:"environments"`
	ActiveEnv   string   `json:"active_env"`
}
