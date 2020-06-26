package dbaerospike

type Config struct {
    HostList string `yaml:"nodes"`
    Keyspace string `yaml:"keyspace"`
    Timeout  int    `yaml:"timeout"`
}
