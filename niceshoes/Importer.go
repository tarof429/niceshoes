package niceshoes

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Importer struct {
	cobblerSystems []CobblerSystem
	importMessages []string
	count int

}

func (importer * Importer) Load(file *string) error {
	data, err := os.ReadFile(*file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}


	if err := json.Unmarshal(data, &importer.cobblerSystems); err != nil {
		log.Fatalf("JSON unmarshalling failed: %s\n", err)
	}

	return nil
}

func (importer *Importer) Import() {

	for _, cs := range importer.cobblerSystems {
		if cs.SystemExists() {
			importer.AddImportMessage(
				fmt.Sprintf("Unable to import %s: already exists", cs.Name))
		} else {
			err := cs.Import()
			if err != nil {
				importer.AddImportMessage(fmt.Sprintf("Error importing %s", cs.Name))
			} else {
				importer.count++
			}
		}
	}
}

func (importer *Importer)AddImportMessage(m string) {
	importer.importMessages = append(importer.importMessages, m)
}

func (importer *Importer) GetImportMessages() [] string {
	return importer.importMessages
}

func (importer *Importer) GetCount() int {
	return importer.count
}