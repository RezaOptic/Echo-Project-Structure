package config

import (
	"fmt"
	"github.com/Echo-Project-Structure/utils"
)


// Port port of server to open and listen to connections
var Port = utils.GetEnv("PORT_TEST", "5000")

// ClientID of JWT
var ClientID = utils.GetEnv("CLIENT_ID_TEST", "CLIENT_ID_TEST")

// ClientSecret of JWT
var ClientSecret = utils.GetEnv("CLIENT_SECRET_TEST", "CLIENT_SECRET_TEST")

// MongoHost host of mongodb to connect
var MongoHost = utils.GetEnv("MONGO_HOST_TEST", "localhost")

// MongoPort port of mongo to connect
var MongoPort = utils.GetEnv("MONGO_PORT_TEST", "27017")

// MongoDBName name of current project database
var MongoDBName = utils.GetEnv("MONGO_DB_NAME_TEST", "sample_db_name")

// MongoConnectionString connection string of mongo database to connect
var MongoConnectionString = fmt.Sprintf("mongodb://%s:%s/%s", MongoHost, MongoPort, MongoDBName)

var MongoURI = utils.GetEnv("MongoURI", MongoConnectionString)

// Init initialize configurations
func Init() {
	Port = utils.GetEnv("PORT", "5000")
	MongoHost = utils.GetEnv("MONGO_HOST", MongoHost)
	MongoPort = utils.GetEnv("MONGO_PORT", MongoPort)
	MongoDBName = utils.GetEnv("MONGO_DB_NAME", "sample_db_name")

	MongoConnectionString = fmt.Sprintf("mongodb://%s:%s/%s", MongoHost, MongoPort, MongoDBName)
	MongoURI = utils.GetEnv("MongoURI", MongoConnectionString)
}
