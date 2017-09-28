package webutil

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"reflect"

	"github.com/kobeld/goutils"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

type serverEngine struct {
	appConf Config
	rpc.Server
}

func NewServerEngine(conf Config, env string, confPath string) *serverEngine {

	val := reflect.ValueOf(conf)
	if val.Kind() != reflect.Ptr {
		panic(errors.New("The conf must be a pointer"))
	}

	loadConfigAndSetupEnv(conf, env, confPath)

	return &serverEngine{
		appConf: conf,
	}

}

func (this *serverEngine) RunRPC() {
	var (
		rpcPort = this.appConf.GetRpcPort()
	)

	if rpcPort == "" {
		panic("Error: No RPC port!")
	}

	log.Printf("Listening and serving RPC on port %s\n", rpcPort)
	listen, err := net.Listen("tcp", rpcPort)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if goutils.HasErrorAndPrintStack(err) {
			continue
		}

		go this.ServeConn(conn)
	}
}

func (this *serverEngine) RunEcho(engine *echo.Echo) {

	if engine == nil {
		engine = echo.New()
	}

	var (
		httpPort = this.appConf.GetHttpPort()
	)

	if httpPort == "" {
		panic("Error: No HTTP port!")
	}

	log.Printf("Listening and serving HTTP on port %s\n", httpPort)
	engine.Start(httpPort)
}

func LoadConfigAndSetupEnv(conf Config, env string, configPath string) {
	loadConfigAndSetupEnv(conf, env, configPath)
}

// ===== Private =====
func loadConfigAndSetupEnv(conf Config, env string, configPath string) {

	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey(fmt.Sprintf("%s.%s", conf.GetName(), env), conf)
	if err != nil {
		panic(err)
	}

	conf.SetConfigFile(configPath)

	conf.AfterInit()

	log.Printf("\nThe configuration is: \n%+v\n\n", conf)
}
