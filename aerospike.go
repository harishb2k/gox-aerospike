package dbaerospike

import (
    "fmt"
    "github.com/aerospike/aerospike-client-go"
    "github.com/harishb2k/gox-base/metrics"
    "github.com/harishb2k/gox-errors"
    "sync"
)
import "github.com/harishb2k/gox-db"
import "github.com/harishb2k/gox-base"

type Context struct {
    Client   *aerospike.Client;
    HostList string
    Port     int
    Keyspace string
    onlyOnce sync.Once
    *base.ApplicationContext
}

func New(context *Context) (idb db.IDb, err error) {
    // Setup application context with defaults
    if context.ApplicationContext == nil {
        context.ApplicationContext = &base.ApplicationContext{}
    }
    context.SetupDefaultIfMissing()
    return context, nil
}

func (context *Context) InitDatabase() (err error) {
    context.onlyOnce.Do(func() {

        client, err := aerospike.NewClient(context.HostList, context.Port)
        if err != nil {
            err = &errors.ErrorObj{
                Name: "failed_to_open_connection",
                Err:  err,
            }
        } else {
            context.Client = client;
        }

        // Register all matrices
        context.Metrics.RegisterCounter("aerospike_find_all", "aerospike_find_all help")
        context.Metrics.RegisterCounter("aerospike_find_all_error", "aerospike_find_all_error help")

    })
    return
}

func (context *Context) FindOne(queryString string, mapper db.RowMapper, val ...interface{}) (result interface{}, e error) {
    defer metrics.LogMetricFunc(context.Metrics, e, "aerospike_find_all", "aerospike_find_all_error")

    // Ensure we have a session before we make any call
    if e = context.ensureSession(); e != nil {
        return
    }

    if key, err := aerospike.NewKey(context.Keyspace, base.Stringify(val[0]), base.Stringify(val[1])); err != nil {
        e = &errors.ErrorObj{
            Name:        db.DatabaseErrorUnknown,
            Err:         err,
            Description: fmt.Sprintf("Failed to create new key [%s]", base.Stringify(val[0])),
        }
    } else {
        if record, err := context.Client.Get(nil, key); err != nil {
            e = &errors.ErrorObj{
                Name:        db.DatabaseErrorRecordNotFound,
                Err:         err,
                Description: fmt.Sprintf("Failed to read data for key [%s]", base.Stringify(val[0])),
            }
        } else {
            result, e = mapper.Map(record)
        }
    }
    return
}

func (context *Context) Persist(queryString string, val ...interface{}) (e error) {
    return errors.New("Not implemented")
}

func (context *Context) FindAll(queryString string, mapper db.RowMapper, val ...interface{}) (result []interface{}, e error) {
    return nil, errors.New("Not implemented")
}

func (context *Context) Execute(queryString string, val ...interface{}) (result interface{}, e error) {
    return nil, errors.New("Not implemented")
}

func (context *Context) ensureSession() (err error) {
    if context.Client == nil {
        return &errors.ErrorObj{
            Name: "db_session_is_null",
            Err:  errors.New("session is not created, you need to initiate it by calling InitDatabase() once"),
        }
    }
    return
}
