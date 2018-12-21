package sign

import (
	"github.com/gotopia/more/config"
	"github.com/gotopia/watcher/issuer"
	"github.com/gotopia/watcher/keychain"
)

var iss *issuer.Issuer

func init() {
	keychain := keychain.New(config.Auth.Sign.KeyDir())
	iss = issuer.New(config.Auth.Sign.Issuer(), keychain)
}

// Issuer returns a global issuer.
func Issuer() *issuer.Issuer {
	return iss
}
