package dbaerospike

type Config struct {
	HostList []string `yaml:"nodes"`
	Port     int      `yaml:"port"`
	Keyspace string   `yaml:"keyspace"`
	Timeout  int      `yaml:"timeout"`
}
