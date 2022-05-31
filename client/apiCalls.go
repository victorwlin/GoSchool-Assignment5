package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type friend struct {
	FriendName  string
	Group       string
	DesiredFreq int
	LastContact string
}

type friends struct {
	Friends []friend
}

const baseURL string = "http://localhost:5000/api/v1/friends"

func getFriends() (friends friends) {
	response, err := http.Get(baseURL)

	if err == nil {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		json.Unmarshal(data, &friends)

		response.Body.Close()
	} else {
		fmt.Println(err)
	}

	return friends
}

func getFriend(friendName string) (friends friends) {
	response, err := http.Get(baseURL + "/" + friendName)

	if err == nil {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		json.Unmarshal(data, &friends)

		response.Body.Close()
	} else {
		fmt.Println(err)
	}

	return friends
}

func addFriendAPI(friend friend) (res int) {
	byteFriend, err := json.Marshal(friend)
	if err != nil {
		fmt.Println(err)
	}
	jsonFriend := bytes.NewBuffer(byteFriend)

	response, err := http.Post(baseURL+"/"+friend.FriendName, "application/json", jsonFriend)

	if err == nil {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		res = response.StatusCode
		fmt.Println(string(data))

		response.Body.Close()
	} else {
		fmt.Println(err)
	}

	return res
}

func editFriendAPI(friend friend) (res int) {
	byteFriend, err := json.Marshal(friend)
	if err != nil {
		fmt.Println(err)
	}
	jsonFriend := bytes.NewBuffer(byteFriend)

	request, err := http.NewRequest(http.MethodPut, baseURL+"/"+friend.FriendName, jsonFriend)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err == nil {

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		res = response.StatusCode
		fmt.Println(string(data))

		response.Body.Close()

	} else {
		fmt.Println(err)
	}

	return res
}

func deleteFriendAPI(friend friend) (res int) {

	request, err := http.NewRequest(http.MethodDelete, baseURL+"/"+friend.FriendName, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err == nil {

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		res = response.StatusCode
		fmt.Println(string(data))

		response.Body.Close()

	} else {
		fmt.Println(err)
	}

	return res
}
