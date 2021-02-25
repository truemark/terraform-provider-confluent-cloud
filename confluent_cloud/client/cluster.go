package client

import (
	"fmt"
	"log"
	"net/url"
	"time"
)

type ClustersResponse struct {
	Clusters []Cluster `json:"clusters"`
}

type ClusterCreateDeploymentConfig struct {
	Sku       string `json:"sku"`
	AccountID string `json:"account_id"`
}

type ClusterCreateConfig struct {
	Name            string                        `json:"name"`
	AccountID       string                        `json:"accountId"`
	Storage         int                           `json:"storage"`
	NetworkIngress  int                           `json:"network_ingress"`
	NetworkEgress   int                           `json:"network_egress"`
	Region          string                        `json:"region"`
	ServiceProvider string                        `json:"serviceProvider"`
	Availability    string                        `json:"availability"`
	Deployment      ClusterCreateDeploymentConfig `json:"deployment"`
	Cku             int                           `json:"cku"`
}

type ClusterCreateRequest struct {
	Config ClusterCreateConfig `json:"config"`
}

type ClusterDeploymentNetworkAccess struct {
	PublicInternet []interface{} `json:"public_internet"`
	VpcPeering     []interface{} `json:"vpc_peering"`
	PrivateLink    []interface{} `json:"private_link"`
	TransitGateway []interface{} `json:"transit_gateway"`
}

type ClusterDeployment struct {
	ID            string                         `json:"id"`
	Created       time.Time                      `json:"created"`
	Modified      time.Time                      `json:"modified"`
	Deactivated   time.Time                      `json:"deactiviated"`
	AccountID     string                         `json:"account_id"`
	NetworkAccess ClusterDeploymentNetworkAccess `json:"network_access"`
	Sku           string                         `json:"sku"`
}

type Cluster struct {
	ID                       string            `json:"id"`
	Name                     string            `json:"name"`
	AccountID                string            `json:"account_id"`
	NetworkIngress           int               `json:"network_ingress"`
	NetworkEgress            int               `json:"network_egress"`
	Storage                  int               `json:"storage"`
	Durability               string            `json:"durability"`
	Status                   string            `json:"status"`
	Endpoint                 string            `json:"endpoint"`
	Region                   string            `json:"region"`
	ServiceProvider          string            `json:"service_provider"`
	OrganizationID           int               `json:"organization_id"`
	Enterprise               bool              `json:"enterprise"`
	K8sClusterID             string            `json:"k8s_cluster_id"`
	PhysicalClusterID        string            `json:"physical_cluster_id"`
	PricePerHour             string            `json:"prince_per_hour"`
	AccruedThisCycle         string            `json:"accrued_this_cycle"`
	Type                     string            `json:"type"`
	APIEndpoint              string            `json:"api_endpoint"`
	InternalProxy            bool              `json:"internal_proxy"`
	IsSLAEnabled             bool              `json:"is_sla_enabled"`
	IsSchedulable            bool              `json:"is_schedulable"`
	Dedicated                bool              `json:"dedicated"`
	NetworkIsolationDomainID string            `json:"network_isolation_domain_id"`
	MaxNetworkIngress        int               `json:"max_network_ingress"`
	MaxNetworkEgress         int               `json:"max_network_egress"`
	Deployment               ClusterDeployment `json:"deployment"`
	Cku                      int               `json:"cku"`
}

type ClusterResponse struct {
	Cluster Cluster `json:"cluster"`
}

func (c *Client) ListClusters(accountID string) ([]Cluster, error) {
	rel, err := url.Parse("clusters")
	if err != nil {
		return []Cluster{}, err
	}

	u := c.BaseURL.ResolveReference(rel)
	response, err := c.NewRequest().
		SetQueryParam("account_id", accountID).
		SetResult(&ClustersResponse{}).
		SetError(&ErrorResponse{}).
		Get(u.String())

	if err != nil {
		return []Cluster{}, err
	}

	if response.IsError() {
		return []Cluster{}, fmt.Errorf("clusters: %s", response.Error().(*ErrorResponse).Error.Message)
	}
	return response.Result().(*ClustersResponse).Clusters, nil
}

func (c *Client) CreateCluster(request ClusterCreateConfig) (*Cluster, error) {

	log.Println("into CreateCluster()")

	rel, err := url.Parse("clusters")
	if err != nil {
		log.Printf("Error occured on url parse to clusters: %s\n", err)
		return nil, err
	}
	log.Printf("rel was: %s\n", rel)

	log.Println("calling (c.BaseURL.ResolveReference)")
	u := c.BaseURL.ResolveReference(rel)
	log.Printf("ResolveRef was: %s ===\n", u)
	log.Printf("request: %s\n", request)

	response, err := c.NewRequest().
		SetBody(&ClusterCreateRequest{Config: request}).
		SetResult(&ClusterResponse{}).
		SetError(&ErrorResponse{}).
		Post(u.String())

	log.Printf("response from call to CreateCluster: \n")
	log.Printf("%s\n", response)

	// Explore response object
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", response.StatusCode())
	log.Println("  Status     :", response.Status())
	log.Println("  Proto      :", response.Proto())
	log.Println("  Time       :", response.Time())
	log.Println("  Received At:", response.ReceivedAt())
	log.Println("  Body       :\n", response)
	log.Println()

	// Explore trace info
	log.Println("Request Trace Info:")
	ti := response.Request.TraceInfo()
	log.Println("  DNSLookup     :", ti.DNSLookup)
	log.Println("  ConnTime      :", ti.ConnTime)
	log.Println("  TCPConnTime   :", ti.TCPConnTime)
	log.Println("  TLSHandshake  :", ti.TLSHandshake)
	log.Println("  ServerTime    :", ti.ServerTime)
	log.Println("  ResponseTime  :", ti.ResponseTime)
	log.Println("  TotalTime     :", ti.TotalTime)
	log.Println("  IsConnReused  :", ti.IsConnReused)
	log.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	log.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	log.Println("  RequestAttempt:", ti.RequestAttempt)
	log.Println("  RemoteAddr    :", ti.RemoteAddr.String())

	// client.OnRequestLog()

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, fmt.Errorf("clusters: %s", response.Error().(*ErrorResponse).Error.Message)
	}

	return &response.Result().(*ClusterResponse).Cluster, nil
}

func (c *Client) DeleteCluster(id, account_id string) error {
	rel, err := url.Parse(fmt.Sprintf("clusters/%s", id))
	if err != nil {
		return err
	}

	u := c.BaseURL.ResolveReference(rel)

	response, err := c.NewRequest().
		SetBody(
			map[string]interface{}{
				"cluster": map[string]interface{}{
					"id":        id,
					"accountId": account_id,
				},
			},
		).
		SetError(&ErrorResponse{}).
		Delete(u.String())

	if err != nil {
		return err
	}

	if response.IsError() {
		return fmt.Errorf("delete cluster: %s", response.Error().(*ErrorResponse).Error.Message)
	}

	log.Printf("[DEBUG] DeleteCluster Success(%s, %s)", id, account_id)
	return nil
}

func (c *Client) GetCluster(id, account_id string) (*Cluster, error) {
	rel, err := url.Parse(fmt.Sprintf("clusters/%s", id))
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	fmt.Println(rel.String())

	response, err := c.NewRequest().
		SetResult(&ClusterResponse{}).
		SetQueryParam("account_id", account_id).
		SetError(&ErrorResponse{}).
		Get(u.String())

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, fmt.Errorf("get cluster: %s", response.Error().(*ErrorResponse).Error.Message)
	}

	return &response.Result().(*ClusterResponse).Cluster, nil
}

func (c *Client) UpdateCluster(id, account_id, name string) error {
	rel, err := url.Parse(fmt.Sprintf("clusters/%s", id))
	if err != nil {
		return err
	}

	u := c.BaseURL.ResolveReference(rel)

	data, err := c.GetCluster(id, account_id)

	if err != nil {
		return err
	}

	data.Name = name

	response, err := c.NewRequest().
		SetBody(&ClusterResponse{Cluster: *data}).
		SetError(&ErrorResponse{}).
		Put(u.String())

	if err != nil {
		return err
	}

	if response.IsError() {
		return fmt.Errorf("update cluster: %s", response.Error().(*ErrorResponse).Error.Message)
	}

	return nil
}
