package cmd

import (
	"github.com/kingdom-tower/kingdom/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ServerCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "server",
		Short: "Runs the kingdom API server",
		RunE:  serverCmdF,
	}

	c.Flags().IntP("port", "p", 8080, "The port to run the server at")
	viper.BindPFlag("port", c.Flags().Lookup("port"))
	c.Flags().StringP("dir", "d", "./data", "The directory where the server reads and stores the files")
	viper.BindPFlag("dir", c.Flags().Lookup("dir"))

	return c
}

func init() {
	RootCmd.AddCommand(
		ServerCmd(),
	)
}

func serverCmdF(c *cobra.Command, args []string) error {
	port := viper.GetInt("port")
	dir := viper.GetString("dir")

	s := server.NewServer(port, dir)
	s.Start()
	return nil
}
