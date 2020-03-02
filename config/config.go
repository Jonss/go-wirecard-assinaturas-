package config

// WirecardConfig is responsible to environment envs
var WirecardConfig wirecardConfig

// Env represents wirecard environment.
type env string

const (
	// PROD sets production wirecard url
	PROD env = "https://api.moip.com.br/assinaturas/v1"
	// SANDBOX sets production wirecard url
	SANDBOX env = "https://sandbox.moip.com.br/assinaturas/v1"
)

type wirecardConfig struct {
	Env   env
	Token string
	Key   string
}
