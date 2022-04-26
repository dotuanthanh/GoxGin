package config

type RootUser struct {
	Email     string `default:"test@gmail.com" env:"ROOT_USER_EMAIL"`
	Phone     string `default:"test@gmail.com" env:"ROOT_USER_PHONE"`
	AccessKey string `default:"abc" env:"ROOT_USER_ACCESS_KEY"`
}
