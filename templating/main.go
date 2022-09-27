package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)


func main() {
	fmt.Println("Sanity check")

	parseSingleFile()
	parseFolder()
}


func parseSingleFile() {
        fmt.Println("parsing single file....")

        tmpl, err := template.ParseFiles("./templates/template_01.tmpl")

        if err != nil {
                log.Fatalln(err)
        }

        tmpl.Execute(os.Stdout, nil)
}

func parseFolder() {
        fmt.Println("parsing folder....")

        tmpl, err := template.ParseGlob("templates/*.tmpl")

        if err != nil {
                log.Fatalln(err)
        }

        tmpl.ExecuteTemplate(os.Stdout, "template_01.tmpl", nil)
}
