package config

type UserServiceConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type ServerConfig struct {
	Name string            `mapstructure:"name" json:"name"`
	Port int               `mapstructure:"port" json:"name"`
	Usc  UserServiceConfig `mapstructure:"user_service" json:"name"`
}
