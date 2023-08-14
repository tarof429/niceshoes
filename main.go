package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"tarof429/niceshoes/niceshoes"
	"time"

	"github.com/briandowns/spinner"
)

var (
	importer niceshoes.Importer
)

func main() {

	log.SetOutput(os.Stdout)

	file := flag.String("file", "", "JSON file containg systems to import")

	flag.Parse()

	err := importer.Load(file)

	if err != nil {
		log.Fatal(err)
	}

	if err == nil {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) 
		s.Color("Black")
		s.Start()
		importer.Import()
		s.Stop()
	}

	actual := importer.GetActual()
	expected := importer.GetExpected()

	for _, message := range importer.GetImportMessages() {
		fmt.Println(message)
	}
	
	if actual == 0 {
		fmt.Printf("%s\n", "Nothing was imported")
	} else {
		fmt.Printf("%d/%d system(s) imported successfully\n", actual, expected)
		fmt.Println("Happy cobblering!")
	}


}
