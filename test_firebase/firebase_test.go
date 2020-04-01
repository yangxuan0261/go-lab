package test_firebase

import (
	"context"
	"fmt"
	"log"
	"testing"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func Test_firebase01(t *testing.T) {
	opt := option.WithCredentialsFile("./temp_test001.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	randTopic := "sdfesdfesdfesdfes"
	// This registration token comes from the client FCM SDKs.
	cliRegToken1 := "fhqXFqU-_K4:APA91bHn6PkgY4xeW0nquD4-KC0SCS0Ly4f2wzhWQAaUW7E2tdWkzCwc2gd3fkCb_dxlfHkQoDNi0AADt6mlk2imcLFRDhG4F9fyQKSDT43vyFo8w2HX2_8xWhtQ1q8pPWKAm5-AHVB8"
	//cliRegToken2 := "dnSPVyzwTFg:APA91bFNQ2qyajGVxYsdw38Z-8sse-YLGlZp60-_btD2pDMVFsR3aQ-43OKAt-1JmTuHw-qb2JUsC-XBc5ktrYEarJ0hX1DcZ_krzV1DYc335eVfZ56L4A51ahcp5cub-ozX-B9ON5PM"
	//cliRegToken3 := "ccqkWMM5CzI:APA91bEXQcd_N0LUOH2xr3WORvtXBTCHr7_FI7dmalQfJVNdIMJ6AVEGLefmmUBtsvEm-y_LrobjYXtirZh-PeUD-0SC0jzkD8Mu3itpFNglHi3MdDUBuC6y2EOdeASkRHZv-nzV5GaL"
	//cliRegToken4 := "eDAbTuRtJwg:APA91bGqM1d1W5v_VyG7j5wOVLPfjGVl2zgjQZPAl27PrxeqqMKsu2p9HPt3lB3ZU0B_GLa_qasbT98XOQfTExCeXJa9vkM1D5NhG3GABLEImaCPJZ2x_NoZcYirO-jzZnhXqsy6cfsN"

	// These registration tokens come from the client FCM SDKs.
	cliRegTokenArr := []string{
		cliRegToken1,
		//cliRegToken2,
		//cliRegToken3,
		//cliRegToken4,
		// ...
	}
	_ = randTopic
	_ = cliRegTokenArr

	subscribe(ctx, client, randTopic, cliRegTokenArr)
	sendMsgToTopic(ctx, client, randTopic)
	unsubscribe(ctx, client, randTopic, cliRegTokenArr) // 发送完就取消订阅

	//sendMsgToToken(ctx, client, cliRegToken2)

	// createCustomToken(ctx, app)
	//sendMsgToToken(ctx, client, cliRegToken2)
}

func subscribe(ctx context.Context, client *messaging.Client, topic string, registrationTokens []string) {
	// Subscribe the devices corresponding to the registration tokens to the topic.
	response, err := client.SubscribeToTopic(ctx, registrationTokens, topic)
	if err != nil {
		log.Fatalln(err)
	}
	// See the TopicManagementResponse reference documentation for the contents of response.
	fmt.Println(response.SuccessCount, "tokens were subscribed successfully")

}

func unsubscribe(ctx context.Context, client *messaging.Client, topic string, registrationTokens []string) {
	response, err := client.UnsubscribeFromTopic(ctx, registrationTokens, topic)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response.SuccessCount, "tokens were unsubscribed successfully")
}

// 对多个 token 发送, 设备必须先订阅某个主题
func sendMsgToTopic(ctx context.Context, client *messaging.Client, topic string) {
	notification := &messaging.Notification{
		Title: "RMG Rummy Station",
		Body:  "Nice to meet you~111\nNice to meet you~222\nNice to meet you~333\n",
		//ImageURL: "https://img-blog.csdnimg.cn/20200216122912287.png", // 必须是 https
	}

	cnt := 9

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"type": "1",
			"msg":  "大王叫我来巡山",
		},
		Notification: notification,
		Topic:        topic,
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Color:                 "#c9a63e",
				DefaultVibrateTimings: true, // 震动
				Priority:              messaging.PriorityMax,
				NotificationCount:     &cnt,

				//ImageURL:"https://img-blog.csdnimg.cn/20200330171658746.jpg", // 覆盖
				Icon: "https://img-blog.csdnimg.cn/20200401153051574.png",
			},
		},
	}

	// Send a message to the devices subscribed to the provided topic.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}

// 对单个 token 发送
func sendMsgToToken(ctx context.Context, client *messaging.Client, registrationToken string) {
	// See documentation on defining a message payload.

	notification := &messaging.Notification{
		Title: "Title002",
		Body:  "Nice to meet you~",
	}

	// timestampMillis := int64(12345)

	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Notification: notification,
		//Webpush: &messaging.WebpushConfig{
		//	Notification: &messaging.WebpushNotification{
		//		Title: "title",
		//		Body:  "body",
		//		//      Icon: "icon",
		//	},
		//	FcmOptions: &messaging.WebpushFcmOptions{
		//		Link: "https://fcm.googleapis.com/",
		//	},
		//},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}

func createCustomToken(ctx context.Context, app *firebase.App) {
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := authClient.CustomToken(ctx, "25696773511053390")
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	log.Printf("Got custom token: %v\n", token)
}
