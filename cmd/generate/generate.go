package main

import (
	"fmt"
	"github.com/alexsuslov/messages"
	"github.com/alexsuslov/wasty/api/gene"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"Title": strings.Title,
}

func main() {
	configFile := os.Args[1]
	name := os.Args[2]

	dirname := fmt.Sprintf("api/%v", name)
	if _, err:=os.Stat(dirname); err!= nil{
		log.Println(err)
		err=os.Mkdir(dirname,  0700)
		if err!= nil{
			log.Println(err)
		}
	}


	configs := gene.ReadConfig(configFile)
	config, ok := configs[name]
	if !ok{
		panic("error config name")
	}
	M := messages.New(config.Name)
	M.SetFuncMap(funcMap)
	err := M.ReLoad("./gene_templates", ".tmpl")
	if err != nil {
		panic(err)
	}
	for _, v := range config.Jobs{
		filename := fmt.Sprintf("api/%v/%v", name, v)
		//if _, err:=os.Stat(filename); err== nil{
		//	err=nil
		//	continue
		//}

		result := M.Execute(v, config)
		err := ioutil.WriteFile( filename,
				[]byte(result), 0644)
		if err!= nil{
			panic(err)
		}
	}
}




