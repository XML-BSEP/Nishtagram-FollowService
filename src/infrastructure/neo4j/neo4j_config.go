package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/viper"
	"os"
)

func init_viper() {
	viper.SetConfigFile(`configurations/neo4j.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func NewNeo4jDriver() (neo4j.Driver, error) {
	init_viper()
	var host string
	if os.Getenv("DOCKER_ENV") == "" {
		host = viper.GetString("neo4j_uri_localhost")
	} else {
		host = viper.GetString("neo4j_uri_docker")

	}
	username := viper.GetString("username")
	password := viper.GetString("password")

	return neo4j.NewDriver(host, neo4j.BasicAuth(username, password, ""))

}
