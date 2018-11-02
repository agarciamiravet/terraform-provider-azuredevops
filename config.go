package main

type Config struct {
	Organization string
	Pat          string
}

// Client represents the DNSimple provider client.
// This is a convenient container for the configuration and the underlying API client.
type Client struct {
	config *Config
}

// Client() returns a new client for accessing dnsimple.
func (c *Config) Client() (*Client, error) {

	provider := &Client{
		config: c,
	}

	return provider, nil
}
