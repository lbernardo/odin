package models

type Module struct {
	Create   Create    `yaml:"create"`
	Commands []Command `yaml:"commands"`
	Resource string    `yaml:"resource"`
}

type Create struct {
	Directories []string `yaml:"directories"`
	Files       []string `yaml:"files"`
}

type Command struct {
	Cmd         string   `yaml:"cmd"`
	Description string   `yaml:"description"`
	Args        []Args   `yaml:"args"`
	Directories []string `yaml:"directories"`
	Files       []string `yaml:"files"`
}

type Args struct {
	Name        string `yaml:"name"`
	Value       string `yaml:"value"`
	Description string `yaml:"description"`
}
