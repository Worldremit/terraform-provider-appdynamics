package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/imroc/req"
)

type Collector struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Hostname  string `json:"hostname"`
	Port      string `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	AgentName string `json:"agentName"`
}

func (c *AppDClient) createCollectorBaseUrl() string {
	return fmt.Sprintf("%s/controller/rest/databases/collectors", c.BaseUrl)
}

func (c *AppDClient) createCollectorUrl() string {
	return fmt.Sprintf("%s/create", c.createCollectorBaseUrl())
}

func (c *AppDClient) updateCollectorUrl() string {
	return fmt.Sprintf("%s/update", c.createCollectorBaseUrl())
}

func (c *AppDClient) deleteCollectorUrl(collectorId int) string {
	return fmt.Sprintf("%s/%d", c.createCollectorBaseUrl(), collectorId)
}

func (c *AppDClient) getCollectorUrl(collectorId int) string {
	return fmt.Sprintf("%s/%d", c.createCollectorBaseUrl(), collectorId)
}

func (c *AppDClient) CreateCollector(collector *Collector) (string, error) {
	url := c.createCollectorUrl()
	auth := c.createAuthHeader()
	body := req.BodyJSON(collector)
	req.Debug = true
	resp, err := req.Post(url, auth, body)
	if err != nil {
		return "", err
	}
	if resp.Response().StatusCode != 201 {
		respString, _ := resp.ToString()
		return "", errors.New(fmt.Sprintf("Error creating Collector: %d, %s\nurl=[%s]", resp.Response().StatusCode, respString, c.createCollectorUrl()))
	}
	//eg.  http://worldremit-test.saas.appdynamics.com/controller/rest/databases/collectors/create/1540
	locationHeader := resp.Response().Header.Get("Location")
	id := locationHeader[strings.LastIndex(locationHeader, "/")+1:]
	fmt.Printf("id=%v\n", id)

	return id, nil
}

func (c *AppDClient) DeleteCollector(collectorId int) error {
	resp, err := req.Delete(c.deleteCollectorUrl(collectorId), c.createAuthHeader())
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return errors.New(fmt.Sprintf("Error deleting Collector: %d, %s", resp.Response().StatusCode, respString))
	}
	return nil
}

/*
func (c *AppDClient) UpdateCollector(dashboard Dashboard) (*Dashboard, error) {
	resp, err := req.Post(c.updateDashboardUrl(), c.createAuthHeader(), req.BodyJSON(dashboard))
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error updating Dashboard: %d, %s", resp.Response().StatusCode, respString))
	}
	updated := Dashboard{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}
	return &updated, err
}



func (c *AppDClient) GetCollector(dashboardId int) (*Dashboard, error) {
	resp, err := req.Get(c.getDashboardUrl(dashboardId), c.createAuthHeader())
	if err != nil {
		return nil, err
	}
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error getting dashboard: %d, %s", resp.Response().StatusCode, respString))
	}
	dashboard := Dashboard{}
	err = resp.ToJSON(&dashboard)
	if err != nil {
		return nil, err
	}
	dashboard.ID = dashboardId
	return &dashboard, nil
}
*/
