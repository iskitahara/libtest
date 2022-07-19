package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Workflow struct {
	Flow  int `yaml:"flow"`
	Retyr int `yaml:"retry"`
}

func main() {
	workflow := Workflow{}
	b, e := os.ReadFile("workflow/index.yaml")
	fmt.Println(e)
	yaml.Unmarshal(b, &workflow)
	fmt.Printf("WorkFlow: %v", workflow)
}
