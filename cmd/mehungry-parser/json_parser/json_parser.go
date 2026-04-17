package json_parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type FdcFoodParserSplitter struct{}

func streamArrayFromKey(filePath string, key string, process func(item interface{}) error) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read opening object "{"
	t, err := decoder.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("expected object start")
	}

	// Find target key
	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}

		k, ok := t.(string)
		if !ok {
			continue
		}

		// Found our key
		if k == key {

			// Expect array start
			t, err := decoder.Token()
			if err != nil {
				return err
			}

			if t != json.Delim('[') {
				return fmt.Errorf("expected array for key %s", key)
			}

			// Stream array items
			for decoder.More() {
				var item interface{}
				if err := decoder.Decode(&item); err != nil {
					return err
				}

				if err := process(item); err != nil {
					return err
				}
			}

			return nil
		}
	}

	return fmt.Errorf("key %s not found", key)
}

func (f *FdcFoodParserSplitter) GetIngredientsFromFoodDataCentralJSONFile(
	filePath string,
	outputDir string,
	key string,
) error {

	os.MkdirAll(outputDir, 0755)

	var (
		chunk []interface{}
		index int
	)

	err := streamArrayFromKey(filePath, key, func(item interface{}) error {
		chunk = append(chunk, item)

		if len(chunk) >= 10 {
			if err := writeSliceToFile(chunk, index, outputDir); err != nil {
				return err
			}

			chunk = chunk[:0]
			index++
		}

		return nil
	})

	if err != nil {
		return err
	}

	// write remaining items
	if len(chunk) > 0 {
		if err := writeSliceToFile(chunk, index, outputDir); err != nil {
			return err
		}
	}

	return nil
}

func writeSliceToFile(slice []interface{}, index int, outputDir string) error {
	fileName := fmt.Sprintf("%s/ing_slice%d.json", outputDir, index)

	data, err := json.Marshal(slice)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

func extractArray(jsonBody map[string]interface{}, key string) ([]interface{}, error) {
	raw, ok := jsonBody[key]
	if !ok {
		return nil, fmt.Errorf("key %s not found", key)
	}

	arr, ok := raw.([]interface{})
	if !ok {
		return nil, fmt.Errorf("key %s is not an array", key)
	}

	return arr, nil
}




func getJSONOr(filename string) ([]interface{}, error) {
	absPath, _ := filepath.Abs(filename)

	body, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	var jsonData []interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (f *FdcFoodParserSplitter) GetIngredientsFromFilesDirectory(filePath string) error {
	allFiles, err := filepath.Glob(filePath + "*.json")

	if err != nil {
		return err
	}

	for _, file := range allFiles {
		jsonBody, err := getJSONOr(file)
		if err != nil {
			return err
		}

		for _, ingredient := range jsonBody {
			// This is equivalent to the Elixir call to Mehungry.FdcFoodParserLeg.create_ingredient
			// You would need to implement this function or call the appropriate Go equivalent
			err := CreateIngredient(ingredient)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Placeholder for the CreateIngredient function that would be implemented elsewhere
func CreateIngredient(ingredient interface{}) error {
	// Implementation would go here
	// This is a placeholder for the Mehungry.FdcFoodParserLeg.create_ingredient function
	return nil
}

func ParseTheFile(path string, outputDir string, key string) {
	fmt.Println("DEBUG path:", path) // 👈 add this

	parser := &FdcFoodParserSplitter{}
  err := parser.GetIngredientsFromFoodDataCentralJSONFile(path, outputDir, key)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
type Record struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}


func StreamParse(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)

    // Read the opening '['
    token, err := decoder.Token()
    if err != nil || token != json.Delim('[') {
        return fmt.Errorf("expected opening [ but got %v", token)
    }

    for decoder.More() {
        var record Record
        if err := decoder.Decode(&record); err != nil {
            return err
        }

        // Process the record
        fmt.Println("Got record:", record)
    }

    return nil
 }

