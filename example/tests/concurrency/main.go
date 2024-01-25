package main

import (
	"fmt"
	"sync"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChan := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultsChan <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// receive expression
		result := <-resultsChan
		// 受け取ったresultをもとにmapを更新
		results[result.string] = result.bool
	}

	return results
}

type User struct {
	id   string
	name string
}

func concurrentUsers(ids []string, users []User) []User {
	channelLen := len(ids)
	ch := make(chan []User, channelLen)

	var wg sync.WaitGroup
	wg.Add(channelLen)

	for _, id := range ids {
		go func(id string) {
			ch <- []User{getUser(id, users)}
			wg.Done()
		}(id)
	}

	wg.Wait()
	close(ch)

	fetchedUsers := []User{}
	for res := range ch {
		fmt.Println("res: ", res)
		fetchedUsers = append(fetchedUsers, res...)
	}

	return fetchedUsers
}

func getUser(id string, users []User) User {
	for _, user := range users {
		if user.id == id {
			return user
		}
	}

	return User{}
}

func main() {
	users := []User{
		{id: "1", name: "hoge"},
		{id: "2", name: "fuga"},
		{id: "3", name: "piyo"},
	}

	u := concurrentUsers([]string{"1", "2"}, users)
	for _, s := range u {
		fmt.Println("u: ", s)
	}

}
