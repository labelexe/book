package config

type (
	Domain struct {
		OneXStavka string `json:"one_x_stavka"`
	}
	DB struct {
		Name     string `json:"name" env:"DB_NAME"`
		Password string `json:"password" env:"DB_PASSWORD"`
		Port     string `json:"port" env:"DB_PORT"`
		Host     string `json:"host" env:"DB_HOST"`
		User     string `json:"user" env:"DB_USER"`
		SslMode  string `json:"ssl_mode" env:"DB_SSL_MODE"`
	}
)
