package config

type Redis struct {
	Host     string `default:"127.0.0.1" env:"APP_REDIS_HOST"`
	Port     uint   `default:"6380" env:"APP_REDIS_PORT"`
	Password string `default:"thanhdt" env:"APP_REDIS_Password"`
	DB       int    `default:"thanhdt" env:"APP_REDIS_DB"`
}
