package config

var Cfg = &cfg{}

type cfg struct {
	App    app    `json:"app"`
	Logger logger `json:"logger"`
	Server server `json:"server"`
}

type app struct {
	Name string `json:"name"`
}

type logger struct {
	File       string `json:"file"`
	Format     string `json:"format"`
	Level      string `json:"level"`
	MaxSize    int    `json:"maxSize"`
	MaxAge     int    `json:"maxAge"`
	MaxBackups int    `json:"maxBackups"`
}

type server struct {
	Port string `json:"port"`
}
