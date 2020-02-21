package config

var WirecardConfig wirecardConfig

// Env represents wirecard environment.
type env string

const (
	PROD    env = "https://api.moip.com.br/assinaturas/v1"
	SANDBOX env = "https://sandbox.moip.com.br/assinaturas/v1"
)

type wirecardConfig struct {
	Env   env
	Token string
	Key   string
}
