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
		Use:   "install",
		Short: "Installs Go and GoUp",
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
			return check()
		},
	}
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(uninstallCmd)
	rootCmd.AddCommand(upgradeCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.Execute()
}
