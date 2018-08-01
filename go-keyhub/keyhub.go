package keyhub

import (
	"net/http"
	"time"

	"github.com/coreos/go-oidc"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/dghubble/sling"
)

type Client struct {
	Version *VersionService
	Groups  *GroupService
	Vaults  *VaultService
}

func NewClient(httpClient *http.Client, issuer string, clientID string, clientSecret string) (*Client, error) {
	if httpClient.Timeout == 0 {
		httpClient.Timeout = time.Duration(time.Second * 10)
	}

	ctx := oidc.ClientContext(context.Background(), httpClient)
	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, err
	}

	var appClientConf = clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{oidc.ScopeOpenID},
		TokenURL:     provider.Endpoint().TokenURL + "?authVault=access",
	}
	oauth2Client := appClientConf.Client(ctx)
	oauth2Client.Timeout = time.Duration(time.Second * 10)

	base := sling.New().Base(issuer).Set("Accept", "application/vnd.topicus.keyhub+json;version=latest")
	oauth2Sling := base.New().Client(oauth2Client)

	vaultClient := &http.Client{
		Transport: &Transport{
			Base: oauth2Client.Transport,
		},
	}

	return &Client{
		Version: newVersionService(base.New().Client(httpClient)),
		Groups:  newGroupService(oauth2Sling.New()),
		Vaults:  newVaultService(base.New().Client(vaultClient), vaultClient),
	}, nil
}
