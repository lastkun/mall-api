package config

type GoodsServiceConfig struct {
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

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name string             `mapstructure:"name" json:"name"`
	Port int                `mapstructure:"port" json:"port"`
	Usc  GoodsServiceConfig `mapstructure:"Goods_service" json:"Goods_service"`
	JWTc JWTConfig          `mapstructure:"jwt" json:"jwt"`
	RC   RedisConfig        `mapstructure:"redis" json:"redis"`
	CC   ConsulConfig       `mapstructure:"consul" json:"consul"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      uint64 `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	Goods     string `mapstructure:"Goods" json:"Goods"`
	Password  string `mapstructure:"password" json:"password"`
	DataId    string `mapstructure:"dataid" json:"dataid"`
	Group     string `mapstructure:"group" json:"group"`
}
