// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"wash-bonus/internal/api/restapi/client/permission"
	"wash-bonus/internal/api/restapi/client/role"
	"wash-bonus/internal/api/restapi/client/standard"
	"wash-bonus/internal/api/restapi/client/user"
	"wash-bonus/internal/api/restapi/client/wash_server"
	"wash-bonus/internal/api/restapi/client/wash_session"
)

// Default wash bonus HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new wash bonus HTTP client.
func NewHTTPClient(formats strfmt.Registry) *WashBonus {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new wash bonus HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *WashBonus {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new wash bonus client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *WashBonus {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(WashBonus)
	cli.Transport = transport
	cli.Permission = permission.New(transport, formats)
	cli.Role = role.New(transport, formats)
	cli.Standard = standard.New(transport, formats)
	cli.User = user.New(transport, formats)
	cli.WashServer = wash_server.New(transport, formats)
	cli.WashSession = wash_session.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// WashBonus is a client for wash bonus
type WashBonus struct {
	Permission permission.ClientService

	Role role.ClientService

	Standard standard.ClientService

	User user.ClientService

	WashServer wash_server.ClientService

	WashSession wash_session.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *WashBonus) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Permission.SetTransport(transport)
	c.Role.SetTransport(transport)
	c.Standard.SetTransport(transport)
	c.User.SetTransport(transport)
	c.WashServer.SetTransport(transport)
	c.WashSession.SetTransport(transport)
}
