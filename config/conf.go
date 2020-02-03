package config

type Conf struct {
	App   App
	Mysql Mysql
	Redis Redis
}

type App struct {
	Name           string
	URl            string
	Port           string
	LoggerLevel    string
	DataBaseDirver string
}

type Mysql struct {
}

type Redis struct {
}
