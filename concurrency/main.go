package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type UserResponse struct {
	Data       []User `json:"data"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Support    struct {
		Url  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

func getUserIds() ([]int, error) {
	endpoint := "https://reqres.in/api/users"
	page := 1
	perPage := 100
	url := fmt.Sprintf("%s?page=%d&per_page=%d", endpoint, page, perPage)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return []int{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []int{}, err
	}
	var userResponse UserResponse
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		fmt.Println(err)
		return []int{}, err
	}
	if len(userResponse.Data) == 0 {
		fmt.Println("No users found")
		return []int{}, err
	}
	userIds := []int{}
	for _, user := range userResponse.Data {
		userIds = append(userIds, user.Id)
	}
	return userIds, nil
}

func getIndividualUser(userId int, wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()
	semaphore <- struct{}{}
	defer func() { <-semaphore }()
	log.Println("Fetching user", userId)
	url := "https://reqres.in/api/users/"
	endpoint := fmt.Sprintf("%s%d", url, userId)
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var wrapper struct {
		Data User `json:"data"`
	}
	err = json.Unmarshal(body, &wrapper)
	if err != nil {
		fmt.Println(err)
		return
	}
	user := wrapper.Data
	log.Println("Fetched user", userId, user)
	time.Sleep(2 * time.Second)
}

func main() {
	userIds, err := getUserIds()
	if err != nil {
		fmt.Println(err)
		return
	}
	semaphore := make(chan struct{}, 5)
	wg := sync.WaitGroup{}
	for _, userId := range userIds {
		wg.Add(1)
		go getIndividualUser(userId, &wg, semaphore)
	}
	wg.Wait()
}
