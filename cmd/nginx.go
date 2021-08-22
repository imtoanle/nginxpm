/*
Copyright © 2021 Toan Le <info@imtoanle.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"fmt"

	rest_api_nginx "github.com/imtoanle/nginxpm/rest-api/nginx"
	common_utils "github.com/imtoanle/nginxpm/rest-api/utils"
	"github.com/spf13/cobra"
)

var (
	nginxCmd = &cobra.Command{
		Use:   "nginx",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("OK Cho nay de run command")
		},
	}

	proxyCmd = &cobra.Command{
		Use:   "proxy",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Proxy command")
		},
	}

	createProxyCmd = &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			rest_api_nginx.CreateProxyHost(map[string]string{
				"domain_name":  domain_name,
				"forward_host": forward_host,
				"forward_port": forward_port,
			})
		},
	}
	domain_name  string
	forward_host string
	forward_port string
)

func init() {
	rootCmd.AddCommand(nginxCmd)
	nginxCmd.AddCommand(proxyCmd)
	proxyCmd.AddCommand(createProxyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	createProxyCmd.Flags().StringVarP(&domain_name, "domain_name", "D", "", "Proxy Host domain")
	createProxyCmd.Flags().StringVarP(&forward_host, "forward_host", "H", common_utils.GetLocalIP(), "Proxy Host ip")
	createProxyCmd.Flags().StringVarP(&forward_port, "forward_port", "P", "", "Proxy Host port")
	createProxyCmd.MarkFlagRequired("domain_name")
	createProxyCmd.MarkFlagRequired("forward_port")

	// viper.BindPFlag(	"foo", nginxCmd.PersistentFlags().Lookup("foo"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createProxyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
