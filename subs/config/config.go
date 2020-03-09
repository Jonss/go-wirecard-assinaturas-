package config

// WirecardConfig is responsible to environment envs
var WirecardConfig wirecardConfig

// Env represents wirecard environment.
type Env string

const (
	// PROD sets production wirecard url
	PROD Env = "https://api.moip.com.br/assinaturas/v1"
	// SANDBOX sets production wirecard url
	SANDBOX Env = "https://sandbox.moip.com.br/assinaturas/v1"
)

type wirecardConfig struct {
	Env   Env
	Token string
	Key   string
}
