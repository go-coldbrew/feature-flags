package engines

// Context specifies the context in which a feature toggle should be considered
// to be enabled or not.
type Context struct {
	// UserId is the the id of the user.
	UserId string

	// SessionId is the id of the session.
	SessionId string

	// RemoteAddress is the IP address of the machine.
	RemoteAddress string

	// Environment is the environment this application is running in.
	Environment string

	// AppName is the application name.
	AppName string

	// Properties is a map of additional properties.
	Properties map[string]string
}

type Payload struct {
	// Type is the type of the payload
	Type string `json:"type"`
	// Value is the value of the payload type
	Value string `json:"value"`
}

type Variant struct {
	// Name is the value of the variant name.
	Name string `json:"name"`
	// Payload is the value of the variant payload
	Payload Payload `json:"payload"`
	// Enabled indicates whether the feature which is extend by this variant was enabled or not.
	Enabled bool `json:"enabled"`
}

type FeatureFlag interface {
	IsEnabled(name string, ctx Context) bool
	GetVariant(name string, ctx Context) *Variant
}
