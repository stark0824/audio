package globalkey

type Config struct {
	App App `yaml:"app" json:"app"`
}

type App struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
	Ssl  bool   `yaml:"ssl" json:"ssl"`
}
