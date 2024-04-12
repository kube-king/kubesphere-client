package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"kubesphere-client/pkg/client/deployment"
	"kubesphere-client/pkg/client/global"
	"kubesphere-client/pkg/client/types"
	"kubesphere-client/pkg/client/utils"
	"log"
	"sync"
	"time"
)

const (
	DefaultClientId     = "kubesphere"
	DefaultClientSecret = "kubesphere"
	GrantType           = "password"
	defaultTimeout      = 15 * time.Second
	defaultTokenExpire  = 30 * time.Second
)

const (
	ApiOAuthToken = "/oauth/token"
)

type Client struct {
	Host         string             // Request Host
	OAuthConfig  *types.OAuthConfig // OAuth Config
	HttpClient   *types.HttpClient
	TokenExpire  time.Duration // createToken expire time
	cache        sync.Map
	tokenContent *types.OAuthToken
}

type OptFunc func(c *Client)

// NewDefaultClient Init Default Client
func NewDefaultClient(opt ...OptFunc) (*Client, error) {
	c := &Client{
		TokenExpire: defaultTokenExpire,
		OAuthConfig: &types.OAuthConfig{
			GrantType:    GrantType,
			ClientId:     DefaultClientId,
			ClientSecret: DefaultClientSecret,
		},
		HttpClient: &types.HttpClient{
			TlsVerify: true,
			Request:   resty.New(),
			Timeout:   defaultTimeout,
		},
	}

	for _, op := range opt {
		op(c)
	}

	err := c.createToken()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			err := c.createToken()
			if err != nil {
				log.Println(err.Error())
			}
			time.Sleep(c.TokenExpire)
		}
	}()

	return c, nil
}

// WithHost Load Host Address
func WithHost(host string) OptFunc {
	return func(c *Client) {
		c.Host = host
	}
}

// WithPasswordAuth Oauth2 Password Setting
func WithPasswordAuth(username, password string) OptFunc {
	return func(c *Client) {
		c.OAuthConfig.Username = username
		c.OAuthConfig.Password = password
	}
}

// WithOAuthConfig Oauth2 Config
func WithOAuthConfig(oauthConfig *types.OAuthConfig) OptFunc {
	return func(c *Client) {
		c.OAuthConfig = oauthConfig
	}
}

func (c *Client) createToken() error {
	var err error
	c.tokenContent = &types.OAuthToken{}
	_, err = c.HttpClient.Request.R().SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}).SetResult(c.tokenContent).SetFormData(map[string]string{
		"grant_type":    c.OAuthConfig.GrantType,
		"username":      c.OAuthConfig.Username,
		"password":      c.OAuthConfig.Password,
		"client_id":     c.OAuthConfig.ClientId,
		"client_secret": c.OAuthConfig.ClientSecret,
	}).Post(utils.ParseUrl(c.Host, ApiOAuthToken, nil))
	if err != nil {
		return fmt.Errorf("create token error:%v", err.Error())
	}
	global.AuthToken = c.tokenContent
	return nil
}

func (c *Client) Deployment() *deployment.Deployment {
	fmt.Println(c.tokenContent)
	return &deployment.Deployment{
		Host:       c.Host,
		HttpClient: c.HttpClient,
	}
}
