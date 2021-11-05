package client

import (
	"fmt"
	"os"
	"testing"
)

func TestCollector(t *testing.T) {
	fmt.Println("Testing colllector")

	baseURL := os.Getenv("APPD_BASE_URL")
	secret := os.Getenv("APPD_COLLECTOR_SECRET")

	client := AppDClient{
		BaseUrl: baseURL,
		Secret:  secret,
	}

	// fmt.Println(client)
	id, err := client.CreateCollector(&Collector{
		Name:      "dummytest-nameaaaa6",
		Type:      "MYSQL",
		Hostname:  "host",
		Port:      "33",
		Username:  "aaa",
		Password:  "bb",
		AgentName: "dbagent",
	})

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("Success %v", id)
}
