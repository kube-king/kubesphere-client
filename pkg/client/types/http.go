package types

import (
	"github.com/go-resty/resty/v2"
	"time"
)

type HttpClient struct {
	Request   *resty.Client // resty client
	Timeout   time.Duration // Request Timeout (default is 5 seconds)
	TlsVerify bool          // Enable Tls Verify (default is false)
}
