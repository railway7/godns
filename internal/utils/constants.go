package utils

import "time"

const (
	// PanicMax is the max allowed panic times.
	PanicMax = 5
	// DNSPOD for dnspod.cn.
	DNSPOD = "DNSPod"
	// HE for he.net.
	HE = "HE"
	// CLOUDFLARE for cloudflare.com.
	CLOUDFLARE = "Cloudflare"
	// ALIDNS for AliDNS.
	ALIDNS = "AliDNS"
	// GOOGLE for Google Domains.
	GOOGLE = "Google"
	// DIGITALOCEAN for DigitalOcean.
	DIGITALOCEAN = "DigitalOcean"
	// DUCK for Duck DNS.
	DUCK = "DuckDNS"
	// DREAMHOST for Dreamhost.
	DREAMHOST = "Dreamhost"
	// DYNV6 for Dynv6.
	DYNV6 = "Dynv6"
	// DYNU for Dynu.
	DYNU = "Dynu"
	// NOIP for NoIP.
	NOIP = "NoIP"
	// SCALEWAY for Scaleway.
	SCALEWAY = "Scaleway"
	// LINODE for Linode.
	LINODE = "Linode"
	// STRATO for Strato.
	STRATO = "Strato"
	// LOOPIASE for LoopiaSE.
	LOOPIASE = "LoopiaSE"
	// INFOMANIAK for Infomaniak.
	INFOMANIAK = "Infomaniak"
	// HETZNER for Hetzner.
	HETZNER = "Hetzner"
	// OVH for OVH.
	OVH = "OVH"
	// IONOS for IONOS.
	IONOS = "IONOS"
	// TransIP for TransIP.
	TRANSIP = "TransIP"
	// IPV4 for IPV4 mode.
	IPV4 = "IPV4"
	// IPV6 for IPV6 mode.
	IPV6 = "IPV6"
	// IPTypeA.
	IPTypeA = "A"
	// IPTypeAAAA.
	IPTypeAAAA = "AAAA"
	// RootDomain.
	RootDomain = "@"
	// Regex pattern to match IPV4 address.
	IPv4Pattern = `((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)`
	// Regex pattern to match IPV6 address.
	IPv6Pattern = `(([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|` +
		`(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|` +
		`(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|` +
		`(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
		`(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
		`(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
		`(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|` +
		`(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))`
	// Regex pattern to match IPV4 and IPV6 address.
	IPPattern = "(" + IPv4Pattern + ")|(" + IPv6Pattern + ")"

	// DefaultTimeout is the default timeout value, in seconds.
	DefaultTimeout = 10
)

type ProviderSetting struct {
	Name        string `json:"name" yaml:"name"`
	Username    bool   `json:"username" yaml:"username"`
	Email       bool   `json:"email" yaml:"email"`
	Password    bool   `json:"password" yaml:"password"`
	LoginToken  bool   `json:"login_token" yaml:"login_token"`
	AppKey      bool   `json:"app_key" yaml:"app_key"`
	AppSecret   bool   `json:"app_secret" yaml:"app_secret"`
	ConsumerKey bool   `json:"consumer_key" yaml:"consumer_key"`
}

var (
	// Version is current version of GoDNS.
	Version = "v0.1"
	// StartTime is the start time of GoDNS.
	StartTime = time.Now().Unix()

	// Providers is the list of supported DNS providers.
	Providers = []ProviderSetting{
		{
			Name:       DNSPOD,
			LoginToken: true,
		}, {
			Name:     HE,
			Password: true,
		},
		{
			Name:       CLOUDFLARE,
			LoginToken: true,
		},
		{
			Name:     ALIDNS,
			Email:    true,
			Password: true,
		},
		{
			Name:     GOOGLE,
			Email:    true,
			Password: true,
		},
		{
			Name:       DIGITALOCEAN,
			LoginToken: true,
		},
		{
			Name:       DUCK,
			LoginToken: true,
		},
		{
			Name:       DREAMHOST,
			LoginToken: true,
		},
		{
			Name:       DYNV6,
			LoginToken: true,
		},
		{
			Name:     DYNU,
			Password: true,
		},
		{
			Name:     NOIP,
			Email:    true,
			Password: true,
		},
		{
			Name:       SCALEWAY,
			LoginToken: true,
		},
		{
			Name:       LINODE,
			LoginToken: true,
		},
		{
			Name:     STRATO,
			Password: true,
		},
		{
			Name:     LOOPIASE,
			Email:    true,
			Password: true,
		},
		{
			Name:     INFOMANIAK,
			Email:    true,
			Password: true,
		},
		{
			Name:       HETZNER,
			LoginToken: true,
		},
		{
			Name:        OVH,
			AppKey:      true,
			AppSecret:   true,
			ConsumerKey: true,
		},
		{
			Name:       IONOS,
			LoginToken: true,
		},
		{
			Name:       TRANSIP,
			LoginToken: true,
		},
	}
)
