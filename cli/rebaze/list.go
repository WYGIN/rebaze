package main

import "github.com/spf13/cobra"

var list = &cobra.Command{
	Use: "list",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Short: "",
	Long: ``,
	Example: ``,
	ValidArgsFunction: inspect.ValidArgsFunction,
	Annotations: map[string]string{},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
