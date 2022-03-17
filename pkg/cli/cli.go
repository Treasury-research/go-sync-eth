package cli

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
)

type CommandFunc func(c *Cli) *cobra.Command

type Cli struct {
	rootCmd *cobra.Command
	viper   *viper.Viper

	readConfigFunc func(viper *viper.Viper)
}

func NewCli(use, version, welcome string) *Cli {
	rootCmd := &cobra.Command{
		Use:     use,
		Version: version,
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println(welcome)
		},
	}
	return &Cli{
		rootCmd: rootCmd,
	}
}

func (c *Cli) RunE(runE func(cmd *cobra.Command, args []string) error) {
	c.rootCmd.RunE = runE
}

func (c *Cli) RootCmd() *cobra.Command {
	return c.rootCmd
}

func (c *Cli) Init(flagSetFunc func(viper *viper.Viper, flags *pflag.FlagSet), readConfigFunc func(viper *viper.Viper)) error {
	err := c.initFlags(flagSetFunc, readConfigFunc)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cli) initFlags(flagSetFunc func(viper *viper.Viper, flags *pflag.FlagSet), readConfigFunc func(viper *viper.Viper)) error {
	var configFile string
	rootFlags := c.rootCmd.PersistentFlags()

	rootFlags.StringVar(&configFile, "config", "./conf/config.yaml", "config file")
	cobra.OnInitialize(func() {
		if configFile != "" {
			viper.SetConfigFile(configFile)
		} else {
			viper.AddConfigPath(".")
			viper.AddConfigPath("./conf")
			viper.SetConfigName("config")
		}
		viper.AutomaticEnv()
		if configFile != "" {
			if err := viper.ReadInConfig(); err != nil {
				log.Print("viper.ReadInConfig err", err)
			}
		}
		if readConfigFunc != nil {
			c.readConfigFunc = readConfigFunc
			c.readConfigFunc(viper.GetViper())
		}
		c.viper = viper.GetViper()
		c.watchConfig()
	})
	viper.BindPFlags(rootFlags)

	if flagSetFunc != nil {
		flagSetFunc(viper.GetViper(), rootFlags)
	}

	return nil
}

func (c *Cli) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})
}

func (c *Cli) GetViper() *viper.Viper {
	return c.viper
}

func (c *Cli) AddCommands(cmds []CommandFunc) {
	for _, cmd := range cmds {
		c.rootCmd.AddCommand(cmd(c))
	}
}

func (c *Cli) AddCommand(cmd *cobra.Command) {
	c.rootCmd.AddCommand(cmd)
}

func AddCommand(rootCmd *cobra.Command, cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func (c *Cli) Execute() {
	err := c.rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (c *Cli) Get(key string) interface{} {
	return c.viper.Get(key)
}

func (c *Cli) GetString(key string) string {
	return c.viper.GetString(key)
}
func (c *Cli) GetInt(key string) int {
	return c.viper.GetInt(key)
}
func (c *Cli) GetInt64(key string) int64 {
	return c.viper.GetInt64(key)
}
func (c *Cli) GetBool(key string) bool {
	return c.viper.GetBool(key)
}
