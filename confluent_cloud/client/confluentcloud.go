package client

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	resty "github.com/go-resty/resty/v2"
)

const (
	defaultBaseURL = "https://confluent.cloud/api/"
	libraryVersion = "0.1"
	userAgent      = "go-confluent-cloud " + libraryVersion
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string
	email     string
	password  string
	token     string
	client    *resty.Client
}

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error ErrorMessage `json:"error"`
}

func sortHeaderKeys(hdrs http.Header) []string {
	keys := make([]string, 0, len(hdrs))
	for key := range hdrs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func copyHeaders(hdrs http.Header) http.Header {
	nh := http.Header{}
	for k, v := range hdrs {
		nh[k] = v
	}
	return nh
}

func composeHeaders(c *resty.Client, r *resty.Request, hdrs http.Header) string {
	str := make([]string, 0, len(hdrs))
	for _, k := range sortHeaderKeys(hdrs) {
		var v string
		if k == "Cookie" {
			cv := strings.TrimSpace(strings.Join(hdrs[k], ", "))
			if c.GetClient().Jar != nil {
				for _, c := range c.GetClient().Jar.Cookies(r.RawRequest.URL) {
					if cv != "" {
						cv = cv + "; " + c.String()
					} else {
						cv = c.String()
					}
				}
			}
			v = strings.TrimSpace(fmt.Sprintf("%25s: %s", k, cv))
		} else {
			v = strings.TrimSpace(fmt.Sprintf("%25s: %s", k, strings.Join(hdrs[k], ", ")))
		}
		if v != "" {
			str = append(str, "\t"+v)
		}
	}
	return strings.Join(str, "\n")
}

const debugRequestLogKey = "__restyDebugRequestLog"

func logReq(c *resty.Client, r *resty.Request) error {
	log.Printf("request: ")
	// rl := &resty.RequestLog{Header: copyHeaders(r.RawRequest.Header)} // Body: r.Body.(string)}
	// rl := &resty.RequestLog{Header: r.RawRequest.Header} // Body: r.Body.(string)}
	// if r.RawRequest != nil {
	reqLog := "\n==============================================================================\n" +
		"~~~ REQUEST ~~~\n" +
		// fmt.Sprintf("%s  %s  %s\n", r.Method, r.RawRequest.URL.RequestURI(), r.RawRequest.Proto) +
		// fmt.Sprintf("HOST   : %s\n", r.RawRequest.URL.Host) +
		// fmt.Sprintf("HEADERS:\n%s\n", composeHeaders(c, r, r.Header)) +
		// fmt.Sprintf("BODY   :\n%v\n", r.Body) +

		fmt.Sprintf("URL   : %s\n", r.URL) +
		fmt.Sprintf("Method   : %s\n", r.Method) +
		fmt.Sprintf("Token   : %s\n", r.Token) +
		fmt.Sprintf("AuthScheme   : %s\n", r.AuthScheme) +
		fmt.Sprintf("QueryParam    : %s\n", r.QueryParam) +
		fmt.Sprintf("FormData   : %s\n", r.FormData) +
		fmt.Sprintf("Header   : %s\n", r.Header) +
		fmt.Sprintf("Time   : %s\n", r.Time) +
		fmt.Sprintf("Body   : %s\n", r.Body) +
		fmt.Sprintf("Result   : %s\n", r.Result) +
		fmt.Sprintf("Error   : %s\n", r.Error) +
		fmt.Sprintf("RawRequest   : %s\n", r.RawRequest) +
		fmt.Sprintf("SRV   : %s\n", r.SRV) +
		fmt.Sprintf("USERINFO   : %s\n", r.UserInfo) +

		"------------------------------------------------------------------------------\n"
	log.Println(reqLog)
	// }
	return nil
}

func NewClient(email, password string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	client := resty.New()
	client.SetDebug(true)
	client.EnableTrace()
	// client.OnRequestLog(resty.requestLog)
	// client.OnResponseLog(resty.responseLog)
	client.OnBeforeRequest(logReq)
	client.OnAfterResponse(func(c *resty.Client, res *resty.Response) error {
		rl := &resty.ResponseLog{Header: copyHeaders(res.Header()), Body: string(res.Body())}
		// if c.responseLog != nil {
		// 	if err := c.responseLog(rl); err != nil {
		// 		return err
		// 	}
		// }
		debugLog := "" // res.Request.values[debugRequestLogKey].(string)
		debugLog += "~~~ RESPONSE ~~~\n" +
			fmt.Sprintf("STATUS       : %s\n", res.Status()) +
			fmt.Sprintf("PROTO        : %s\n", res.RawResponse.Proto) +
			fmt.Sprintf("RECEIVED AT  : %v\n", res.ReceivedAt().Format(time.RFC3339Nano)) +
			fmt.Sprintf("TIME DURATION: %v\n", res.Time()) +
			"HEADERS      :\n" +
			composeHeaders(c, res.Request, rl.Header) + "\n"
		debugLog += fmt.Sprintf("BODY         :\n%v\n", rl.Body)
		debugLog += "==============================================================================\n"
		log.Println(debugLog)
		// c.log.Debugf("%s", debugLog)
		return nil
	})

	c := &Client{BaseURL: baseURL, email: email, password: password, UserAgent: userAgent}
	c.client = client
	return c
}

func (c *Client) NewRequest() *resty.Request {
	return c.client.R().
		SetHeader("User-Agent", c.UserAgent).
		SetError(&ErrorResponse{})
}
