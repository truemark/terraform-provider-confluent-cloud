package client

import (
	"fmt"
	"log"
	"net/url"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthSuccessResponse struct {
	Token string `json:"token"`
}

func (c *Client) Login() error {
	log.Printf("[INFO] ConfluentCloud client - Login()")

	rel, err := url.Parse("sessions")
	if err != nil {
		return err
	}

	u := c.BaseURL.ResolveReference(rel)

	response, err := c.NewRequest().
		SetBody(AuthRequest{Email: c.email, Password: c.password}).
		SetResult(&AuthSuccessResponse{}).
		Post(u.String())

	log.Println("Response from Login-Request(): %s\n", response)
	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", response.StatusCode())
	fmt.Println("  Status     :", response.Status())
	fmt.Println("  Proto      :", response.Proto())
	fmt.Println("  Time       :", response.Time())
	fmt.Println("  Received At:", response.ReceivedAt())
	fmt.Println("  Body       :\n", response)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := response.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	if ti.RemoteAddr != nil {
		fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
	}
	if err != nil {
		return err
	}

	if response.IsError() {
		return fmt.Errorf("login: %s", response.Error().(*ErrorResponse).Error.Message)
	}

	if response.IsSuccess() {
		c.token = response.Result().(*AuthSuccessResponse).Token
	}
	return nil
}
