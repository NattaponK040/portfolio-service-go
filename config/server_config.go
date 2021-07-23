package config

type ServerConfig struct {
	Server Server `mapstructure:"server"`
	MongoDb MongoDb `mapstructure:"mongoDb"`
}

type Server struct {
	Name    string  `mapstructure:"name"`
	Port    int     `mapstructure:"port"`
}
type MongoDb struct {
	Uri      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
	MaxPool  uint64 `mapstructure:"maxpool"`
	MinPool  uint64 `mapstructure:"minpool"`
}
