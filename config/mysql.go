package config

type MySql struct {
	Host        string `default:"127.0.0.1" env:"APP_MYSQL_DB_HOST"`
	Port        uint   `default:"3306" env:"APP_MYSQL_DB_PORT"`
	DBName      string `default:"test" env:"APP_MYSQL_DB_NAME"`
	User        string `default:"admin" env:"APP_MYSQL_DB_USER"`
	Password    string `default:"123456" env:"APP_MYSQL_DB_PASSWORD"`
	MaxOpenConn int    `default:"100" env:"APP_MYSQL_DB_MAX_OPEN_CONN"`
	MaxIdleConn int    `default:"100" env:"APP_MYSQL_DB_MAX_IDLE_CONN"`
	DriverName  string `default:"mysql" env:"APP_MYSQL_DB_DRIVER_NAME"`
}
