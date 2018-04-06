package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

var url string

func main() {
	lookupServiceWithConsul()

	fmt.Println("Starting Simple Client.")
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	callHelloEvery(5*time.Second, client)
}

func lookupServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	services, err := consul.Agent().Services()
	if err != nil {
		fmt.Println(err)
		return
	}

	service := services["simple-server"]
	address := service.Address
	port := service.Port

	url = fmt.Sprintf("htt://%s:%v/info", address, port)
}

func hello(t time.Duration, client *http.Client) {

	response, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%s. Time is %v\n", body, t)
}

func callHelloEvery(t time.Duration, client *http.Client) {
	if true {
		hello(t, client)
		time.Sleep(t)
	}
}
