package webutil

const (
	ENV_DEV     = "dev"
	ENV_PROD    = "prod"
	ENV_TEST    = "test"
	ENV_CI      = "ci"
	ENV_STAGING = "staging"
)

type Config interface {
	IsDev() bool
	IsProd() bool
	IsTest() bool
	IsCI() bool
	IsStaging() bool

	GetName() string
	GetHttpPort() string
	GetRpcPort() string

	SetConfigFile(path string)
	AfterInit()
}

// The base configuration file can be embedded by app's specified configuration
type BaseConfig struct {
	ConfigFile  string
	Env         string
	ServiceName string
	DbName      string
	DbDail      string
	RpcPort     string
	HttpPort    string
}

func (this *BaseConfig) IsDev() bool     { return this.Env == ENV_DEV }
func (this *BaseConfig) IsProd() bool    { return this.Env == ENV_PROD }
func (this *BaseConfig) IsTest() bool    { return this.Env == ENV_TEST }
func (this *BaseConfig) IsCI() bool      { return this.Env == ENV_CI }
func (this *BaseConfig) IsStaging() bool { return this.Env == ENV_STAGING }

func (this *BaseConfig) GetHttpPort() string { return this.HttpPort }
func (this *BaseConfig) GetRpcPort() string  { return this.RpcPort }

func (this *BaseConfig) SetConfigFile(path string) { this.ConfigFile = path }
func (this *BaseConfig) AfterInit()                {}
