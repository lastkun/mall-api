package config

type UserServiceConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type RedisConfig struct {
	Host   string `mapstructure:"host" json:"host"`
	Port   int    `mapstructure:"port" json:"port"`
	Expire int    `mapstructure:"expire" json:"expire"`
}

type ServerConfig struct {
	Name string            `mapstructure:"name" json:"name"`
	Port int               `mapstructure:"port" json:"name"`
	Usc  UserServiceConfig `mapstructure:"user_service" json:"name"`
	JWTc JWTConfig         `mapstructure:"jwt" json:"jwt"`
	RC   RedisConfig       `mapstructure:"redis" json:"redis"`
}
