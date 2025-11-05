package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// MidtransConfig menyimpan konfigurasi API Midtrans
type MidtransConfig struct {
	ServerKey string
	ClientKey string
}

// LoadMidtransConfig memuat konfigurasi Midtrans dari environment variable
func LoadMidtransConfig() (*MidtransConfig, error) {
	viper.SetDefault("MIDTRANS_SERVER_KEY", "")
	viper.SetDefault("MIDTRANS_CLIENT_KEY", "")

	viper.AutomaticEnv()

	serverKey := viper.GetString("MIDTRANS_SERVER_KEY")
	clientKey := viper.GetString("MIDTRANS_CLIENT_KEY")

	if serverKey == "" || clientKey == "" {
		return nil, fmt.Errorf("MIDTRANS_SERVER_KEY atau MIDTRANS_CLIENT_KEY belum dikonfigurasi")
	}

	fmt.Println("Midtrans success ... !")
	return &MidtransConfig{
		ServerKey: serverKey,
		ClientKey: clientKey,
	}, nil
}
