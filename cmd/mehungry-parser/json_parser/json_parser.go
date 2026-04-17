package json_parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type FdcFoodParserSplitter struct{}

func getJSON(filename string) (map[string]interface{}, error) {
	//os.Mkdir("fdc_legacy_splited_files", 0755)

	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (f *FdcFoodParserSplitter) GetIngredientsFromFoodDataCentralJSONFile(filePath string, 	outputDir string, key string,) error {
	jsonBody, err := getJSON(filePath)
	if err != nil {
		return err
	}

	// FoundationFoods // SRLegacyFoods //
  ingredients, err := extractArray(jsonBody, key)
if err != nil {
	return err
}

	// Split ingredients into chunks of 10
	var chunks [][]interface{}
	for i := 0; i < len(ingredients); i += 10 {
		end := i + 10
		if end > len(ingredients) {
			end = len(ingredients)
		}
		chunks = append(chunks, ingredients[i:end])
	}

	// Write each chunk to a file
	for i, chunk := range chunks {
		err := writeSliceToFile(chunk, i, outputDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeSliceToFile(slice []interface{}, index int, outputDir string) error {
	os.MkdirAll(outputDir, 0755)

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

