package cmd

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simple todolist",
	Short: "Simple todolist",
	Long:  "functional todolist CLI. Add, delete, update tasks from your terminal",
	Run: func(cmd *cobra.Command, args []string) {
		usage()
	},
}

func usage() {
	pterm.Println()
	y, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Y", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("TD", pterm.FgLightMagenta.ToStyle())).
		Srender()
	pterm.DefaultCenter.Println(y)

	subtext := pterm.FgCyan.Sprint("Yet Another Todolist")
	pterm.DefaultCenter.Println(subtext)

	availableCommands()
	exampleUsage()

}

func availableCommands() {
	pterm.Println(pterm.FgLightMagenta.Sprint("Avaialble commands"))
	pterm.DefaultTable.WithData([][]string{
		{pterm.FgYellow.Sprint("Add"), pterm.FgCyan.Sprint("add a task entry to todo-list")},

		{pterm.FgYellow.Sprint("View"), pterm.FgCyan.Sprint("lists all entries in todolist")},

		{pterm.FgYellow.Sprint("Delete"), pterm.FgCyan.Sprint("remove a task from todolist")},

		{pterm.FgYellow.Sprint("complete"), pterm.FgCyan.Sprint("mark a task when completed")},
	}).Render()

	pterm.Println()
}

func exampleUsage() {
	pterm.Println(pterm.FgLightMagenta.Sprint("Usage example"))
	pterm.DefaultTable.WithData([][]string{
		{pterm.FgYellow.Sprint("Add"), pterm.FgCyan.Sprint("./ytd add \"task entry goes here\"")},

		{pterm.FgYellow.Sprint("View"), pterm.FgCyan.Sprint("./ytd view")},

		{pterm.FgYellow.Sprint("Delete 'taskid'"), pterm.FgCyan.Sprint("./ytd delete 69")},

		{pterm.FgYellow.Sprint("complete 'taskid' "), pterm.FgCyan.Sprint("./yd complete 69")},
	}).Render()

	pterm.Println()
	pterm.Println()

}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
