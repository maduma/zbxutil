package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func api_call(jsreq JSONRequest) []byte {
	data, err := json.Marshal(jsreq) // data is a byte slice
	if err != nil {
		panic(err)
	}

	//fmt.Println("JSON data:", string(data))

	buf := bytes.NewBuffer(data) // convert slice bytes to io.Reader
	api_url := zabbix_server_url + "/zabbix/api_jsonrpc.php"
	req, err := http.NewRequest("POST", api_url, buf)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 2 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body) // resp.Body is an io.Reader
	if err != nil {
		panic(err)
	}

	return body
}
