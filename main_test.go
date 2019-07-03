package main

import (
	"encoding/json"
	"fmt"
	handlers "github.com/fervargas94/proxy-app/api/handlers"
	server "github.com/fervargas94/proxy-app/api/server"
	utils "github.com/fervargas94/proxy-app/api/utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
)

type Response struct {
	Status string `json:"status,omitempty"`
	Response string `json:"result,omitempty"`
	ResponseText []ResponseText `json:"res,omitempty"`
}

type ResponseText struct {
	/*
	response:
		{domain string}
		{domain string}
		{domain string}
	 */
	Domain string
}

func init() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		utils.LoadEnv()
		app := server.SetUp()
		handlers.HandlerRedirection(app)
		wg.Done()
		server.RunServer(app)
	}(wg)
	wg.Wait()
	fmt.Println("Server running...")

}

func TestAlgorithm(t *testing.T) {

	cases := []struct {
		//attributes
		Domain string
		Output string
	}{
		//structs
		{Domain: "alpha", Output: "[\"alpha\"]"},
		{Domain: "beta", Output: "[\"alpha\",\"beta\"]"},
		{Domain: "beta", Output: "[\"alpha\",\"beta\",\"beta\"]"},
		{Domain: "omega", Output: "[\"alpha\",\"omega\",\"beta\",\"beta\"]"},
		{Domain: "alpha", Output: "[\"alpha\",\"alpha\",\"omega\",\"beta\",\"beta\"]"},
		{Domain: "", Output: "error"},

	}

	for _, singleCase := range cases {
		valuesToCompare  := &Response{}
		client := http.Client{

		}
		req, err := http.NewRequest("GET","http://localhost:8080/ping", nil )
		req.Header.Set("domain", singleCase.Domain)


		response, _ := client.Do(req)

		bytes, err := ioutil.ReadAll(response.Body)

		json.Unmarshal(bytes, valuesToCompare)

		assert.Nil(t, err)
		assert.Equal(t, valuesToCompare.Response, singleCase.Output)
		assert.True(t, true)
	}


}

