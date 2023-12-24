package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	credsFilePath := "./gcp-credentials.json"

	// Create a new client with the credentials
	client, err := firestore.NewClient(ctx, firestore.DetectProjectID, option.WithCredentialsFile(credsFilePath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	return client
}

type User struct {
	Id          string    `firestore:"id"`
	Name        string    `firestore:"name"`
	Age         int       `firestore:"age"`
	CreatedAt   time.Time `firestore:"created_at"`
	Delete_flag bool      `firestore:"delete_flag"`
}

func main() {

	ctx := context.Background()
	fmt.Println(ctx)
	userId := "0027cc0b-8d5c-4bae-9fa4-7562f19eee28"
	userId2 := "0027f178-fe4b-4890-a420-8736c69ec8ea"
	client := createClient(ctx)

	// 処理の開始時刻を記録
	startTime := time.Now()
	ch := make(chan []User, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go getUser(ctx, client, userId, &wg, ch)
	go getUser(ctx, client, userId2, &wg, ch)

	wg.Wait()
	close(ch)
	for data := range ch {
		fmt.Println("受信したデータ:", data)
	}
	// 処理の終了時刻を記録し、経過時間を計算
	elapsed := time.Since(startTime)

	// 経過時間を出力
	fmt.Printf("総合計時間: %s\n", elapsed)

}

func getUser(ctx context.Context, client *firestore.Client, userId string, wg *sync.WaitGroup, ch chan []User) {
	// 処理の開始時刻を記録
	startTime := time.Now()

	// ここで初めてclientが評価される（遅延評価）
	// Goは初期化に約150msで、Nodeは1.4s。
	// getUserは、90msでNodeと同じ。
	snapshot, err := client.Collection("users").Doc(userId).Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(snapshot.Data())
	var user User
	user.Id = snapshot.Ref.ID // == userId
	if err := snapshot.DataTo(&user); err != nil {
		log.Fatalf("Failed to map data to User struct: %v", err)
	}
	ch <- []User{user}

	// 処理の終了時刻を記録し、経過時間を計算
	elapsed := time.Since(startTime)

	// 経過時間を出力
	fmt.Printf("実行時間: %s\n", elapsed)
	wg.Done()
}
