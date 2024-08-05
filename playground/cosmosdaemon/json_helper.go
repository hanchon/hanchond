package cosmosdaemon

import (
	"encoding/json"

	"github.com/hanchon/hanchond/playground/filesmanager"
)

func readJSONFile(path string) (map[string]interface{}, error) {
	bytes, err := filesmanager.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func saveJSONFile(data map[string]interface{}, path string) error {
	values, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return filesmanager.SaveFile(values, path)
}
