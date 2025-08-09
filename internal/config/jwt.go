package config

import "github.com/spf13/viper"

type JWTConfig struct {
	Issuer           string
	Secret           string
	SecretRefresh    string
	Expiration       uint
	ExpirationRefreh uint
}

func NewJWTManager(viper *viper.Viper) *JWTConfig {
	issuer := viper.GetString("JWT_ISSUER")
	secret := viper.GetString("JWT_SECRET")
	secretRefresh := viper.GetString("JWT_SECRET_REFRESH")
	exp := viper.GetUint("JWT_EXPIRATION")
	expRefresh := viper.GetUint("JWT_EXPIRATION_REFRESH")

	return &JWTConfig{
		Issuer:           issuer,
		Secret:           secret,
		Expiration:       exp,
		SecretRefresh:    secretRefresh,
		ExpirationRefreh: expRefresh,
	}
}
