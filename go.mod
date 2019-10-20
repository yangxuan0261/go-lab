module GoLab

go 1.13

require (
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/garyburd/redigo v1.6.0
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/golang/protobuf v1.3.2
	github.com/olzhy/quote v1.0.0
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.2.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.24.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

// replace "go.uber.org/zap" => "./vendor/go.uber.org/zap"
