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

	randTopic := "25696773511053390"
	// This registration token comes from the client FCM SDKs.
	cliRegToken := "d-stxUmDf2w:APA91bHO7rEIoWkYe1Sck7rdvHye5S2TlEsMcWDTmudjylNeUR8MT0KswScpsCUA8U2c23g3VH82gou_6iaHKD4zCgAVCwV9G6AFg3Nq04MOybmEIlr-P7VHlc1pr6O6dKFkikxVDRdS"

	// These registration tokens come from the client FCM SDKs.
	cliRegTokenArr := []string{
		cliRegToken,
		// ...
	}

	subscribe(ctx, client, randTopic, cliRegTokenArr)
	sendMsgToTopic(ctx, client, randTopic)
	unsubscribe(ctx, client, randTopic, cliRegTokenArr) // 发送完就取消订阅
	// createCustomToken(ctx, app)
	//sendMsgToToken(ctx, client, cliRegToken)
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
		Title: "Title001",
		Body:  "Nice to meet you~",
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		//Data: map[string]string{
		//	"score": "88888",
		//	"time":  "2:45",
		//},
		Notification: notification,
		Topic:        topic,
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
		// Data: map[string]string{
		//  "score": "850",
		//  "time": "2:45",
		// },
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
