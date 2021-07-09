module FollowService

go 1.16

replace github.com/jelena-vlajkov/logger/logger => ../../Nishtagram-Logger/

require (
	github.com/casbin/casbin/v2 v2.31.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/go-resty/resty/v2 v2.6.0
	github.com/google/uuid v1.2.0
	github.com/jelena-vlajkov/logger/logger v1.0.0
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/neo4j/neo4j-go-driver/v4 v4.3.2
	github.com/spf13/viper v1.7.1
	github.com/ugorji/go v1.2.6 // indirect
	go.mongodb.org/mongo-driver v1.5.2
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210525143221-35b2ab0089ea // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)
