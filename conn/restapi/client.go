package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)
const (
	MaxIdleConns int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout int = 90
)

var (
	httpClient *http.Client
)

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:	 time.Duration(IdleConnTimeout) * time.Second,
		},
	}
	return client
}

func Init() {
	httpClient = createHTTPClient()
}

func DailServer(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	rebots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", rebots)
}

func Conn() {
	//response, err := http.Get("https://httpbin.org/ip")
	//if err != nil {
	//	fmt.Printf("The HTTP request failed with error %s\n", err)
	//} else {
	//	data, _ := ioutil.ReadAll(response.Body)
	//	fmt.Println(1111, string(data))
	//}
	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(2222, string(data))
	}
}
