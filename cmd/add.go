/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
        "fmt"

        "github.com/spf13/cobra"
        "github.com/iByteABit256/tri/todo"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
        Use:   "add",
        Short: "Add a new Todo task",
        Long: `Add will create a new Todo item in the Todo list`,
        Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
    items := []todo.Item{}

    for _, x := range args {
        items = append(items, todo.Item{Text:x})
        fmt.Printf("%v added to list.\n", x)
    }

    fmt.Printf("List: %#v\n", items)

    filename := "my-todo-list.json"

    todo.SaveItems(filename, items)

    fmt.Printf("Saved to %v.", filename)
}

func init() {
        rootCmd.AddCommand(addCmd)

        // Here you will define your flags and configuration settings.

        // Cobra supports Persistent Flags which will work for this command
        // and all subcommands, e.g.:
        // addCmd.PersistentFlags().String("foo", "", "A help for foo")

        // Cobra supports local flags which will only run when this command
        // is called directly, e.g.:
        // addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
