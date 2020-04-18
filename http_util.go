package main;

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// http get request
func HttpGet(url string) string {
	log.Printf("HttpGetJSON, url:[ %v ]", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	log.Printf("HttpGetJSON, url:[ %v ], resp:[ %v ]", url, bodyString)
	return bodyString
}

// http get and convert response into a JSON object
func HttpGetJSON(url string, jsonObj interface{}) {
	log.Printf("HttpGetJSON, url:[ %v ]", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	log.Printf("HttpGetJSON, url:[ %v ], resp:[ %v ]", url, string(body))
	json.Unmarshal([]byte(body), &jsonObj)
}


func HttpPostForm(url string, data url.Values) string {
	req, _ := json.Marshal(data);
	log.Printf("HttpPostForm, url:[ %v ], params:[ %s ]", url, string(req))
	resp, err := http.PostForm(url, data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	log.Printf("HttpPostForm, url:[ %v ], params:[ %v ], resp:[ %v ]", url, string(req), bodyString)
	return bodyString
}

func HttpPostFormReturnJson(url string, data url.Values, jsonObj interface{})  {
	req, _ := json.Marshal(data);
	log.Printf("HttpPostForm, url:[ %v ], params:[ %s ]", url, string(req))
	resp, err := http.PostForm(url, data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	log.Printf("HttpPostForm, url:[ %v ], params:[ %v ], resp:[ %v ]", url, string(req), bodyString)
	json.Unmarshal([]byte(bodyString), &jsonObj)
}
