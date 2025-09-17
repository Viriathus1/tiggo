/*
Copyright © 2025 Viriathus1 <49337323+Viriathus1@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"github.com/Viriathus1/tiggo/internal/tui"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Shows the commit logs",
	Run: func(cmd *cobra.Command, args []string) {
		cm := exec.Command("git", "log", "--oneline", "-n", "10")
		output, err := cm.Output()
		if err != nil {
			fmt.Printf("an error has appeared: %v", err)
			os.Exit(1)
		}

		commitLines := strings.Split(strings.TrimSpace(string(output)), "\n")
		if _, err := tea.NewProgram(tui.NewLogList(commitLines)).Run(); err != nil {
			fmt.Printf("error running program: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
