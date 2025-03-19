package main

import (
  "github.com/nikolis/MehungryGoParser/cmd/mehungry-parser/cli"
)

// If you need a main function for testing
func main() {
	// Example usage
  cli.Execute()

  /*var rootCmd = &cobra.Command{
        Use:   "greet",
        Short: "A simple greeting CLI",
        Long:  "A simple CLI that greets a user based on the provided name and count.",
        Run: func(cmd *cobra.Command, args []string) {
                fmt.Printf("Hello")
        },
    }

  var exitCmd = &cobra.Command{
		Use:   "exit",
		Short: "Exit the program",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Exiting program...")
			os.Exit(0)
		},
	}
  //var count int = 0
 
  //rootCmd.PersistentFlags().IntVarP(&count, "count", "c", 1, "Number of greetings")
    
  //  if err := rootCmd.Execute(); err != nil {
  //      fmt.Println(err)
  //      os.Exit(1)
   // }


//	parser := &FdcFoodParserSplitter{}
//	err := parser.GetIngredientsFromFoodDataCentralJSONFile("../FoodData_Central_sr_legacy_food_json_2021-10-28.json")
//	if err != nil {
//		fmt.Println("Error:", err)
//	} */
}
