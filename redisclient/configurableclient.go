package redisclient

import (
	"gopkg.in/redis.v4"
	"github.com/spf13/viper"
	"fmt"
)

var viperInstance *viper.Viper

func init(){
	viperInstance = viper.New()
	//add search path for the config file name
	viperInstance.AddConfigPath(".")
	viperInstance.SetConfigName("config")
	viperInstance.SetConfigType("json")
	viperInstance.ReadInConfig()
}

func Hey() {
	fmt.Println(viperInstance.GetStringSlice("cluster_nodes"))

	client  := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: viperInstance.GetStringSlice("cluster_nodes")})
	client.Set("HEYYYYYY",42424,0)
	l := client.Get("HEYYYYYY")
	fmt.Println("SPACE")
	fmt.Println(l.Result())
}