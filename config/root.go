package config

type Email struct {
	Email     string `default:"test@gmail.com" env:"ROOT_USER_EMAIL"`
	Phone     string `default:"test@gmail.com" env:"ROOT_USER_PHONE"`
	AccessKey string `default:"hihihi" env:"ROOT_USER_ACCESS_KEY"`
}
