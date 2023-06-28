package frotel

import (
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"net/http"
)

func HttpClient(c *http.Client) *http.Client {
	if c == nil {
		c = http.DefaultClient
	}
	transport := c.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	return &http.Client{
		Transport:     otelhttp.NewTransport(transport, otelhttp.WithMeterProvider(otel.GetMeterProvider())),
		CheckRedirect: c.CheckRedirect,
		Jar:           c.Jar,
		Timeout:       c.Timeout,
	}
}
