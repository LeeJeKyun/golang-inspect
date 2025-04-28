package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"net"
	"net/http"
	"strings"
	"time"
)

func ClientConfig() *opensearch.Client {
	cfg := opensearch.Config{
		Addresses: []string{"https://localhost:9200"},
		Username:  "admin",
		Password:  "admin",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	client, err := opensearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	return client
}

func main() {
	//insert
	client := ClientConfig()

	info := client.Info

	fmt.Println("client.info:", info)

	document := strings.NewReader(`{
		"name": "test1",
		"age": 19
	}`)

	docId := "1"
	req := opensearchapi.IndexRequest{
		Index:      "go-test-index1",
		DocumentID: docId,
		Body:       document,
	}

	insert, err1 := req.Do(context.Background(), client)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println("insert:", insert.String())

	//select
	query := strings.NewReader(`{
		"size": 5,
		"query": {
			"match_all": {}
		}
	}`)

	search := opensearchapi.SearchRequest{
		Index: []string{"go-test-index1"},
		Body:  query,
	}

	result, err2 := search.Do(context.Background(), client)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("select:", result.String())

	//delete
	delete := opensearchapi.DeleteRequest{
		Index:      "go-test-index1",
		DocumentID: "1",
	}

	response, err3 := delete.Do(context.Background(), client)
	if err3 != nil {
		panic(err3)
	}

	fmt.Println("delete:", response.String())

}
