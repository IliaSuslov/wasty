package gene

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Field struct {
	Name       string
	Label      string
	Type       string
	Annotation string
}

type Configs map[string]Config

type Config struct {
	Jobs []string
	Name string
	List struct{
		Permission []string
		Filter []Field
	}
	Fields []Field
	Collection string
}

type Job struct {
	Template   string
	Filename   string
}


func ReadConfig(filename string) map[string]Config {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	c := map[string]Config{}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	return c
}

