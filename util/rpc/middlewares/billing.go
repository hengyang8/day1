package middlewares

import (
	"github.com/Conflux-Chain/go-conflux-util/viper"
	web3pay "github.com/Conflux-Chain/web3pay-service/client"
	"github.com/openweb3/go-rpc-provider"
	"github.com/scroll-tech/rpc-gateway/util/rpc/handlers"
	"github.com/sirupsen/logrus"
)

func MustNewWeb3PayClient() (*web3pay.Client, bool) {
	var config struct {
		web3pay.ClientConfig `mapstructure:",squash"`
		Enabled              bool
	}
	viper.MustUnmarshalKey("web3pay", &config)

	if !config.Enabled {
		return nil, false
	}

	client, err := web3pay.NewClient(config.ClientConfig)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to new Web3Pay client")
	}

	return client, true
}

func Billing(client *web3pay.Client) rpc.HandleCallMsgMiddleware {
	mwoption := web3pay.NewOw3BillingMiddlewareOptionWithClient(client)
	mwoption.ApiKeyProvider = handlers.GetAccessTokenFromContext
	return web3pay.Openweb3BillingMiddleware(mwoption)
}
