package models

type Module struct {
	Create   Create                 `yaml:"create"`
	Commands map[string]interface{} `yaml:"commands"`
	Resource string                 `yaml:"resource"`
}

type Create struct {
	Directories []string `yaml:"directories"`
	Files       []string `yaml:"files"`
}
