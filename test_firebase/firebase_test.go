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
	//cliRegToken1 := "cZeBocpAO8c:APA91bELuBaRgBnz2qlqt8hVcR0TPjf7yDHAxExI7PKC-mBquIAtZFQuxa-aVACyt1HEGaMfNB3tTyluAdswP_ClF8nvGD1Wa_ALOV7tDHdUebsvcs7sJTueiaR3jDyhOgEGaCcztz5y"
	//cliRegToken2 := "eeaAQ6tdrrY:APA91bEXmLUYWd2Go1oGe2t1hjhJIik2CrvshAtc8976dM0WYT6jezRcpCmfttXTI-JKsVI1zsGl50U4j4jLrHZ9yPubYokBeV3Sf0xfgu0b9j2sqZNnoDN16qMLGYxhyb3jvU7PrX54"
	//cliRegToken3 := "f8GFKqYVGLo:APA91bGp0Ze6tRzsevqliCKjL16CPGy7S6nU5-1Ufbtalw4pX6ORIyuNkC4qZVgJsZMynuE04EcbTu2hyRio5bRvhYu4P08C47I2qwyWLTTA-zQ8EqGoM5W2wyAG2VJ7914lg80vc8bN"
	cliRegToken4 := "cE2sUszKbvs:APA91bGTEkdgQAEPFCALT_71xXwA-_i-dCKKnGw6bINvo-4_EuURkUI2J-7ug3b0hCe9ux35-2g-k_fMOZzuXiCe5guOYmiho4_rgpjQzqDQnFJxmm7h4G-hMFkbVeCXB1KEqiKz_zql"
	cliRegToken5 := "f_UVCtvIoDU:APA91bFVPfrDqBS1B6qKYTcDS7fajNZYyAGv5k1kSxdfR4FckIJ0c9NNV3LkO21H6me61NOxKpwyzXDvXEANYNsqVsPjoegkuOFCave1kq7TsNiALu1ao8MYAAlHGDV6ooHWCIsPgj3z"

	// These registration tokens come from the client FCM SDKs.
	cliRegTokenArr := []string{
		//cliRegToken1,
		//cliRegToken2,
		//cliRegToken3,
		cliRegToken4,
		cliRegToken5,
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
		Title:    "RMG Rummy Station",
		Body:     "Nice to meet you~111\nNice to meet you~222\nNice to meet you~333\n",
		ImageURL: "https://img-blog.csdnimg.cn/20200216122912287.png", // 必须是 https
	}

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
