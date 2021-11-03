package client

import (
	"errors"
	"fmt"

	"github.com/imroc/req"
)

type Collector struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Hostname  string `json:"hostname"`
	Port      int    `json:"port"`
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

func (c *AppDClient) CreateCollector(collector Collector) (*Collector, error) {
	resp, err := req.Post(c.createCollectorUrl(), c.createAuthHeader(), req.BodyJSON(collector))
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error creating Collector: %d, %s", resp.Response().StatusCode, respString))
	}
	updated := Collector{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}
	return &updated, err
}

func (c *AppDClient) DeleteCollector(collectorId int) error {
	resp, err := req.Post(c.deleteCollectorUrl(), c.createAuthHeader(), req.BodyJSON(collectorId))
	if err != nil {
		return err
	}
	if resp.Response().StatusCode != 204 {
		respString, _ := resp.ToString()
		return errors.New(fmt.Sprintf("Error deleting Dashboard: %d, %s", resp.Response().StatusCode, respString))
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
