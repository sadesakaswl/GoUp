package main

import (
	"github.com/spf13/cobra"
)

func initCLI() {
	rootCmd := &cobra.Command{
		Use:     "goup",
		Short:   "GoUp is a install and upgrade tool",
		Long:    "GoUp provides installing and upgrading basically!\nhttps://github.com/sadesakaswl/GoUp",
		Version: "0.1",
	}
	installCmd := &cobra.Command{
		Use:   "install [version]",
		Short: "Installs Go and GoUp.",
		Long:  "Installs Go and GoUp.\nDefault version: latest",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return install("latest")
			}
			return install(args[0])
		},
	}
	uninstallCmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstalls Go and GoUp",
		RunE: func(cmd *cobra.Command, args []string) error {
			return uninstall()
		},
	}
	upgradeCmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrades Go and GoUp",
		RunE: func(cmd *cobra.Command, args []string) error {
			return install("latest")
		},
	}
	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "Checks for latest version of Go",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := check()
			return err
		},
	}
	getCmd := &cobra.Command{
		Use:   "get [version]",
		Short: `Installs a version with "go get"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return get("latest")
			}
			return get(args[0])
		},
	}
	deleteCmd := &cobra.Command{
		Use:   "delete [version]",
		Short: `Deletes a version which installed with "goup get" or "go get"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return delete("latest")
			}
			return delete(args[0])
		},
	}
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(upgradeCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.Execute()
}
