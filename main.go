package main

import (
	"fmt"

	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/util/uuid"
)

func main() {
	// DoLogrus()
	// DoCobra()
	DoViper()
	// DoPflag()
	// DoRuntime()
	DoUUID()
}

// func DoLogrus() {
// 	log := logrus.StandardLogger()
// 	log.Info("hello world!")
// }

// func DoCobra() {
// 	rootCmd := &cobra.Command{}
// 	rootCmd.Run = func(cmd *cobra.Command, args []string) {}
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		panic(err)
// 	}
// }

func DoViper() {
	if !viper.GetBool("somevar") {
		fmt.Println("hello")
	}
}

// func DoPflag() {
// 	pflag.Bool("foo", true, "bar")
// }

// func DoRuntime() {
// 	foo := runtime.Protocol(0)
// 	fmt.Println(foo)
// }

func DoUUID() {
	fmt.Println(uuid.NewUUID())
}
