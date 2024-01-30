package cmd

import (
	"log"
	"sort"
	"strconv"

	"github.com/iByteABit256/tri/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Completes a TODO item",
	Long:  `Marks a TODO item as completed`,
	Run:   markDone,
}

func markDone(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		log.Fatalln("%v", err)
	}

	if len(args) == 0 {
		log.Fatalln("You need to provide a label in order to mark a task as done")
		return
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("%v is not a valid label:\n%v\n", args[0], err)
	}

	if i > 0 && i <= len(items) {
		items[i-1].Done = true
		log.Printf("%q marked done\n", items[i-1].Text)
		sort.Sort(todo.ByPri(items))
		todo.SaveItems(dataFile, items)
	} else {
		log.Fatalf("%v doesn't match any items\n", i)
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
