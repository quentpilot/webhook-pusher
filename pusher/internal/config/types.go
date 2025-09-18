package config

type AmqpPubConfig struct {
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
	Queue string `yaml:"queue"`
}
