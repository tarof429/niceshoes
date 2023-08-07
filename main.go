package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"tarof429/niceshoes/niceshoes"
	"time"

	"github.com/briandowns/spinner"
)

func load(file *string) ([]niceshoes.SystemCobblerSystem, error) {
	data, err := os.ReadFile(*file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil, err
	}

	var cs []niceshoes.SystemCobblerSystem

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
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) 
		s.Color("Black")
		s.Start()

		for _, cs := range css {
			//log.Printf("Importing: %s", cs)
			err := cs.Import()
			//err := cs.ImportSimulator()
			if err != nil {
				log.Printf("Unable to import %s", cs.Name)
			}
		}
		s.Stop()
	}

}
