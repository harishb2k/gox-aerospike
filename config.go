package dbaerospike

type Config struct {
    Nodes    []string `yaml:"nodes"`
    Keyspace string   `yaml:"keyspace"`
    Timeout  int      `yaml:"timeout"`
}
