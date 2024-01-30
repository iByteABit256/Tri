package cmd

import (
	"log"

	"github.com/iByteABit256/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new Todo task",
	Long:  `Add will create a new Todo item in the Todo list`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	for _, x := range args {
		item := todo.Item{Text: x, Priority: priority}
		item.SetPriority(priority)
		items = append(items, item)
		log.Printf("%v added to list.\n", x)
	}

	err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "set a priority (1,2,3) on a new task")
}
