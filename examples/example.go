package examples

import "github.com/harishb2k/gox-aerospike"

func main() {

    db, _ := dbaerospike.New(
        &dbaerospike.Context{
            HostList: []string{"127.0.0.7"},
            Port:     3001,
            Keyspace: "test",
        },
    )
    db.InitDatabase()



}
