package client

import (
	"fmt"
	"os"
	"strconv"
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
		Name:      "dummytest-nameaaaa9",
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
	fmt.Printf("Successfully created collector id = %v \n", id)

	id1, _ := strconv.Atoi(id)

	err1 := client.DeleteCollector(id1)

	if err1 != nil {
		fmt.Println(err1)
		t.Fail()
	}

	fmt.Printf("Successfully deleted collector id = %v \n", id1)

}
