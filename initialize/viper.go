package initialize

import (
	"fmt"

	"github.com/Ricky-fight/car-admin-server/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// func Viper() {
// 	viper.AddConfigPath(".")
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// }
func Viper() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error while reading config file: %s", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
}
