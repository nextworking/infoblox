package infoblox

import (
	"crypto/tls"
	"golang.org/x/net/publicsuffix"
	"log"
	"net/http"
	"net/http/cookiejar"
)

var (
	// WapiVersion specifies the version of the Infoblox REST API to target
	WapiVersion = "2.6.1"

	// BasePath specifies the default path prefix to all WAPI actions
	BasePath = "/wapi/v" + WapiVersion + "/"

	// Debug mode flag
	Debug = false
)

type Client struct {
	Host       string
	Password   string
	Username   string
	HTTPClient *http.Client
	UseCookies bool
}

func NewClient(host, username, password string, sslVerify, useCookies bool) *Client {

  var (
	  req, _    = http.NewRequest("GET", host, nil)
	  proxy, _  = http.ProxyFromEnvironment(req)
	  transport *http.Transport
	  tlsconfig *tls.Config
  )

	tlsconfig = &tls.Config{
		InsecureSkipVerify: !sslVerify,
	}
	if tlsconfig.InsecureSkipVerify {
		log.Printf("WARNING: SSL cert verification  disabled\n")
	}
	transport = &http.Transport{
		TLSClientConfig: tlsconfig,
	}
	if proxy != nil {
		transport.Proxy = http.ProxyURL(proxy)
	}

	client := &Client{
		Host: host,
		HTTPClient: &http.Client{
			Transport: transport,
		},
		Username:   username,
		Password:   password,
		UseCookies: useCookies,
	}
	if useCookies {
		options := cookiejar.Options{
			PublicSuffixList: publicsuffix.List,
		}
		jar, _ := cookiejar.New(&options)
		client.HTTPClient.Jar = jar
	}

	return client

}