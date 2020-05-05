package models

type OdinConfig struct {
	Config Config `yaml:"config"`
}

type Config struct {
	Default string `yaml:"default"`
}

type ProjectConfig struct {
	Pkg string
}
