package models

type Migration struct {
  ProjectPath   string  `yaml:"projectPath"`
  From          string  `yaml:"from"`
  To            string  `yaml:"to"`
}

type Config struct {
  Migrations  []Migration `yaml:"migrations"`
}

