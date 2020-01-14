package main

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	/*Setear valores default
	VALORES INICIALES
	name_example = lore
	color_example = green
	number_example = 1
	*/
	viper.SetDefault("name_example", "lore")
	viper.SetDefault("color_example", "green")
	viper.SetDefault("number_example", 5)
	/*Configuracion para leer el archivo yaml
	VALORES LUEGO DE LEER EL YAML
	name_example = owo
	color_example = red
	number_example = 1
	*/
	viper.SetConfigName("def")
	viper.AddConfigPath("./files")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	/*Leer y/o sobreescribir variables de entorno
	VALORES LUEGO DE LEER VARIABLE DE ENTORNO (export COLOR_EXAMPLE=blue)
	name_example = owo
	color_example = blue
	number_example = 1
	*/
	viper.AutomaticEnv()
	/*Leer y/o sobre escribir flags
	VALORES LUEGO DE LEER FLAGS (go run main.go --number_example=4)
	name_example = owo
	color_example = blue
	number_example = 4
	*/
	flag.Int("number_example", 3, "number")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlag("number_example", pflag.CommandLine.Lookup("number_example"))
	viper.BindPFlags(pflag.CommandLine)
	fmt.Println("el nombre es", viper.Get("name_example"))
	fmt.Println("el color es", viper.Get("color_example"))
	fmt.Println("el número es", viper.Get("number_example"))

}
