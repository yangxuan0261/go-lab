module go-lab

go 1.14

require (
	firebase.google.com/go v3.12.0+incompatible
	github.com/DataDog/zstd v1.4.4 // indirect
	github.com/buaazp/fasthttprouter v0.1.1
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/hpcloud/tail v1.0.0
	github.com/json-iterator/go v1.1.7
	github.com/kavu/go_reuseport v1.4.0 // indirect
	github.com/labstack/echo/v4 v4.1.14
	github.com/micro/go-micro v1.13.1 // indirect
	github.com/mitchellh/mapstructure v1.1.2
	github.com/olzhy/quote v1.0.0
	github.com/pkg/errors v0.8.1
	github.com/satori/go.uuid v1.2.0
	github.com/steakknife/bloomfilter v0.0.0-20180922174646-6819c0d2a570
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
	github.com/tidwall/evio v1.0.3
	github.com/valyala/fasthttp v1.6.0
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.2.0
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20191227163750-53104e6ec876
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	golang.org/x/text v0.3.2
	google.golang.org/api v0.8.0
	google.golang.org/grpc v1.24.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/yaml.v2 v2.2.2
)

// replace "go.uber.org/zap" => "./vendor/go.uber.org/zap"
