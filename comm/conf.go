package comm

type Config struct {
	Base appConfig `yaml:"base" json:"base"`
}

type appConfig struct {
	Host        string   `json:"host" yaml:"host"`
	Network     string   `json:"network" yaml:"network"`
	ServerID    int32    `json:"serverID" yaml:"serverID"`
	ServerName  string   `json:"serverName" yaml:"serverName"`
	RpcPath     string   `json:"rpcPath" yaml:"rpcPath"`
	Version     string   `json:"version" yaml:"version"`
	EtcdAddress []string `json:"etcdAddress" yaml:"etcdAddress"`
}
