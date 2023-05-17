package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// var rootCmd = &cobra.Command{
// 	Use:   "gh-tpm",
// 	Short: "gh-tpm is a GitHub cli extension for TPMs",
// 	Long:  `gh-tpm is a GitHub cli extension for TPMs, focusing on program management tasks such as bulk issue creation, issue labeling, and issue assignment.`,
// Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("root called")
// 	},
// }

func NewCmdRoot() *cobra.Command {
	cmdRoot := &cobra.Command{
		Use:   "gh-tpm",
		Short: "gh-tpm is a GitHub cli extension for TPMs\n",
		Long:  `gh-tpm is a GitHub cli extension for TPMs, 
		focusing on program management tasks such as bulk issue creation, issue labeling, and issue assignment.`,
	}
	cmdRoot.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	fmt.Printf("cmdRoot: Inside Command Root\n")
	cmdRoot.AddCommand(LabelsCmd)
	return cmdRoot
}
