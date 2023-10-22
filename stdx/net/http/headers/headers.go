package headers

const (
	AcceptKey        HeaderKey = "Accept"
	AuthorizationKey HeaderKey = "Authorization"

	CookieKey HeaderKey = "Cookie"

	ContentTypeKey     HeaderKey = "Content-Type"
	ContentLengthKey   HeaderKey = "Content-Length"
	ContentEncodingKey HeaderKey = "Content-Encoding"
)

const (
	RetryAfterKey HeaderKey = "Retry-After"
)

const (
	// Host
	// Identifies the original host and optionally the port requested by the client in the Host HTTP request header.
	Host HeaderKey = "Host"

	// XRealIP
	// Identifies the client's IP address.
	XRealIP HeaderKey = "X-Real-IP"

	// XForwardedFor
	// Provides a list of connection IP addresses.
	// The load balancer appends the last remote peer address to the XForwardedFor field from the incoming request.
	// A comma and space precede the appended address.
	// If the client request header does not include an XForwardedFor field,
	// this value is equal to the XRealIP value.
	// The original requesting client is the first (left-most) IP address in the list,
	// assuming that the incoming field content is trustworthy.
	// The last address is the last (most recent) peer, that is, the machine from which the load balancer received the request.
	// The format is: `X-Forwarded-For: original_client, proxy1, proxy2`
	XForwardedFor HeaderKey = "X-Forwarded-For"

	// XForwardedHost
	// Identifies the original host and port requested by the client in the Host HTTP request header.
	// This header helps you determine the original host, since the hostname or port of the reverse proxy (load balancer) might differ from the original server handling the request.
	XForwardedHost HeaderKey = "X-Forwarded-Host"

	// XForwardedPort
	// Identifies the listener port number that the client used to connect to the load balancer.
	XForwardedPort HeaderKey = "X-Forwarded-Port"
)
