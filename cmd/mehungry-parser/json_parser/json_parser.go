package json_parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)


type FdcFoodParserSplitter struct{}

func getJSON(filename string) (map[string]interface{}, error) {
	os.Mkdir("fdc_legacy_splited_files", 0755)
	os.Chdir("fdc_legacy_splited_files")
	
	body, err := ioutil.ReadFile(filename)
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

func (f *FdcFoodParserSplitter) GetIngredientsFromFoodDataCentralJSONFile(filePath string) error {
	jsonBody, err := getJSON(filePath)
	if err != nil {
		return err
	}
  // FoundationFoods // SRLegacyFoods //	
	ingredients, ok := jsonBody["SRLegacyFoods"].([]interface{})
	if !ok {
		return fmt.Errorf("failed to parse SRLegacyFoods")
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
		err := writeSliceToFile(chunk, i)
		if err != nil {
			return err
		}
	}
	
	return nil
}

func writeSliceToFile(slice []interface{}, index int) error {
	fileName := fmt.Sprintf("ing_slice%d.json", index)
	
	data, err := json.Marshal(slice)
	if err != nil {
		return err
	}
	
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	
	return nil
}

func getJSONOr(filename string) ([]interface{}, error) {
  absPath, _ := filepath.Abs(filename)

  body, err := ioutil.ReadFile(absPath)
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



func ParseTheFile(path string){
  parser := &FdcFoodParserSplitter{}
  err := parser.GetIngredientsFromFoodDataCentralJSONFile(path)
 	if err != nil {
		fmt.Println("Error:", err)
	} 
}

