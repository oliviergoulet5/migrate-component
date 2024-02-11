package models

type Migration struct {
  ProjectPath   string          `yaml:"projectPath"`
  From          string          `yaml:"from"`
  To            string          `yaml:"to"`
  Components    map[string]bool `yaml:"components"`
}

type Config struct {
  Migrations  []Migration `yaml:"migrations"`
}

