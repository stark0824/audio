package config

type userRpc struct {
	App App `yaml:"app" json:"app"`
}

type App struct {
	Host string `yaml:"host" json:"host"`
}
