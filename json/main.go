package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func getUsers() ([]User, error) {
	url := "https://reqres.in/api/users"
	page := 1
	perPage := 100
	endpoint := fmt.Sprintf("%s?page=%d&per_page=%d", url, page, perPage)
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return []User{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []User{}, err
	}
	var userResponse UserResponse
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		fmt.Println(err)
		return []User{}, err
	}
	if len(userResponse.Data) == 0 {
		fmt.Println("No users found")
		return []User{}, err
	}
	return userResponse.Data, nil
}

func writeUsersToFile(users []User) error {
	fileName := "users.json"
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	n, err := file.Write(data)
	if err != nil {
		return err
	}
	fmt.Println(n, "bytes were written to file:", fileName)
	return nil
}

func downloadImage(avatar string, name string) error {
	res, err := http.Get(avatar)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()
	file, err := os.Create(name + ".png")
	if err != nil {
		return err
	}
	defer file.Close()
	n, err := io.Copy(file, res.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Downloaded %v's image of %d bytes\n", name, n)
	return nil
}

func main() {
	users, err := getUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = writeUsersToFile(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range users {
		err := downloadImage(user.Avatar, user.FirstName)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
