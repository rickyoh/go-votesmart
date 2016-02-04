package votesmart

import (
	"log"
	"errors"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Apiserver    string
	Key  		 string
	Format   	 string
}

func Connect(Apiserver string, Key string, Format string) *Client {
	return &Client{
		Apiserver:  Apiserver,
		Key:   		Key,
		Format: 	Format,
	}
}

func (c *Client) Query(method string, args []map[string]string) interface{}{

	client := http.Client{}

	url := c.Apiserver
	url = url + "/" + method
	url = url + "?key=" + c.Key
	url = url + "&o=" + c.Format

	for _, arg := range args {
		url = url + "&" + arg["key"] + "=" + arg["value"]
	}
	log.Println(url)
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)

	data := map[string]interface{}{}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &data)

	if(err != nil) {
		log.Println(err)
		errors.New("Problem with votesmart")
	}
	return data
}