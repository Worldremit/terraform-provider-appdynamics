package client

import (
	"fmt"
	"testing"
)

func TestCollector(t *testing.T) {
	fmt.Println("Testing colllector")

	baseURL := "" //os.Getenv("worldremit-production_2db6e402-09cb-4a7d-8a9a-0a60a7cc0c89")
	secret := ""  //os.Getenv("APPD_SECRET")

	client := AppDClient{
		BaseUrl: baseURL,
		Secret:  secret,
	}

	// fmt.Println(client)
	collector, err := client.CreateCollector(Collector{
		Name:      "name",
		Type:      "type",
		Hostname:  "host",
		Port:      33,
		Username:  "aaa",
		Password:  "bb",
		AgentName: "fake",
	})

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(collector)
}
