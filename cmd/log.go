/*
Copyright © 2025 Viriathus1 <49337323+Viriathus1@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"github.com/Viriathus1/tiggo/internal/gitclient"
	"github.com/Viriathus1/tiggo/internal/tui"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Shows the commit logs",
	Run: func(cmd *cobra.Command, args []string) {
		gitClient, err := gitclient.NewGitClient()
		if err != nil {
			fmt.Printf("an error has appeared: %v", err)
			os.Exit(1)
		}

		if _, err := tea.NewProgram(tui.NewLogList(gitClient)).Run(); err != nil {
			fmt.Printf("error running program: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
