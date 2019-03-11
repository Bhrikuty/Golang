package main

import (
	"html/template"
	"os"
)

const tax = 6.75 / 100

type Product struct {
	Name  string
	Price float32
}

const templateString = `
{{- "Item Information" }}
Name: {{ .Name }}
Price: {{ printf "$%.2f" .Price }}
Price with Tax: {{ calctax .Price | printf "$%.2f" }}   
`

//calling the function here using the key "calctax" and passing the argument as ".Price"

func main() {
	p := Product{
		Name:  "Lemonade",
		Price: 2.16,
	}
	fm := template.FuncMap{}                      //FuncMap is just an alias for string of map
	fm["calctax"] = func(price float32) float32 { //assigning value to key
		return price * (1 + tax)
	}
	t := template.Must(template.New("").Funcs(fm).Parse(templateString))
	t.Execute(os.Stdout, p)
}
