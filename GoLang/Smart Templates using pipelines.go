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

func (p Product) PriceWithTax() float32 {
	return p.Price * (1 + tax)
}

// - sign escapes the white space before the string

const templateString = `
{{- "Item Information"}}  
Name: {{ .Name }}
Price: {{ printf "$%.2f" .Price }}
Price with Tax: {{ .PriceWithTax | printf "$%.2f" }}   
`

//this | operation will pass the o/p of func "PriceWithTax" as the last argument to the printf
func main() {
	p := Product{
		Name:  "Lemonade",
		Price: 2.16,
	}
	t := template.Must(template.New("").Parse(templateString))
	t.Execute(os.Stdout, p) // p as a data context instead of nil
}
