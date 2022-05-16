module go-lab

go 1.14

require (
	firebase.google.com/go v3.12.0+incompatible
	github.com/buaazp/fasthttprouter v0.1.1
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/emersion/go-imap v1.0.4
	github.com/emersion/go-message v0.11.1
	github.com/fsnotify/fsnotify v1.4.7 // indirect
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/hpcloud/tail v1.0.0
	github.com/json-iterator/go v1.1.7
	github.com/kavu/go_reuseport v1.4.0 // indirect
	github.com/labstack/echo/v4 v4.1.14
	github.com/mitchellh/mapstructure v1.1.2
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/olzhy/quote v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0
	github.com/steakknife/bloomfilter v0.0.0-20180922174646-6819c0d2a570
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
	github.com/tidwall/evio v1.0.3
	github.com/valyala/fasthttp v1.6.0
	go.mongodb.org/mongo-driver v1.5.1
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.2.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/text v0.3.5
	google.golang.org/api v0.8.0
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55 // indirect
	google.golang.org/grpc v1.24.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/go-playground/validator.v9 v9.30.0 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

// replace "go.uber.org/zap" => "./vendor/go.uber.org/zap"
