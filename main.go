package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"tarof429/niceshoes/niceshoes"
)



func load(file *string) ([]niceshoes.Csystem, error) {
	data, err := ioutil.ReadFile(*file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil, err
	}

	var cs []niceshoes.Csystem

	if err := json.Unmarshal(data, &cs); err != nil {
		log.Fatalf("JSON unmarshalling failed: %s\n", err)
	}

	return cs, nil
}

func main() {

	log.SetOutput(os.Stdout)

	file := flag.String("file", "", "JSON file containg systems to import")

	flag.Parse()
	
	css, err := load(file)

	if err != nil {
		log.Fatal(err)
	}

	if err == nil {
		for _, cs := range css {
			fmt.Printf("Importing: %s", cs)
			err := cs.Import()
			//err := cs.ImportSimulator()
			if err != nil {
				fmt.Println("Unable to import ", cs.Name)
			}
		}
	}

}
