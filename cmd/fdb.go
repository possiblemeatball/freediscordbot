package cmd

import (
	"log"
	"log/slog"
	"os"

	"github.com/kirsle/configdir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	fdbCmd = &cobra.Command{
		Use:   "fdb",
		Short: "free and open source general purpose Discord bot",
		Long: `FreeDiscordBot: A free and open source general-purpose Discord guild
management, community engagement, entertainment, and moderation bot with a
built-in HTTP server control panel that allows for Discord members to log
in frictionlessly for internal application configuration and management.`,
		Run: ServeBot,
	}
	userConfigPath string
)

func ServeBot(cmd *cobra.Command, args []string) {

}

func init() {
	log.SetFlags(0)

	fdbCmd.Flags().StringVarP(&userConfigPath, "config", "c", "", "custom path for application config")
	cobra.OnInitialize(InitConfig)
}

func InitConfig() {
	if userConfigPath != "" {
		err := configdir.MakePath(userConfigPath)
		cobra.CheckErr(err)
		viper.AddConfigPath(userConfigPath)
	}

	viper.AddConfigPath(configdir.LocalConfig("freediscordbot"))
	for _, path := range configdir.SystemConfig("freediscordbot") {
		viper.AddConfigPath(path)
	}

	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	cobra.CheckErr(viper.ReadInConfig())
	slog.Info("found config file " + viper.ConfigFileUsed())
}

func Execute() {
	err := fdbCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
