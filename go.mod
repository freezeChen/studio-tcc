module studio-tcc

go 1.12

require (
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/freezeChen/studio-library v0.0.16
	github.com/gin-gonic/gin v1.3.0
	github.com/go-sql-driver/mysql v1.4.0
	github.com/go-xorm/xorm v0.7.1
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.3.1
	github.com/json-iterator/go v1.1.6
	github.com/micro/go-micro v1.7.0
	google.golang.org/grpc v1.21.1
)

replace (
	github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.4
	github.com/ugorji/go/codec => github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8

)
