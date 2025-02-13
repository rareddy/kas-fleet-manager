package ocm

import (
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/shared"
	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

const (
	MockModeStubServer            = "stub-server"
	MockModeEmulateServer         = "emulate-server"
	strimziOperatorAddonID        = "managed-kafka"
	kasFleetshardAddonID          = "kas-fleetshard-operator"
	ClusterLoggingOperatorAddonID = "cluster-logging-operator"
)

type OCMConfig struct {
	BaseURL                       string `json:"base_url"`
	AmsUrl                        string `json:"ams_url"`
	ClientID                      string `json:"client-id"`
	ClientIDFile                  string `json:"client-id_file"`
	ClientSecret                  string `json:"client-secret"`
	ClientSecretFile              string `json:"client-secret_file"`
	SelfToken                     string `json:"self_token"`
	SelfTokenFile                 string `json:"self_token_file"`
	TokenURL                      string `json:"token_url"`
	Debug                         bool   `json:"debug"`
	EnableMock                    bool   `json:"enable_mock"`
	MockMode                      string `json:"mock_type"`
	StrimziOperatorAddonID        string `json:"strimzi_operator_addon_id"`
	KasFleetshardAddonID          string `json:"kas_fleetshard_addon_id"`
	ClusterLoggingOperatorAddonID string `json:"cluster_logging_operator_addon_id"`
}

func NewOCMConfig() *OCMConfig {
	return &OCMConfig{
		BaseURL:                "https://api-integration.6943.hive-integration.openshiftapps.com",
		AmsUrl:                 "https://api.stage.openshift.com",
		TokenURL:               "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token",
		ClientIDFile:           "secrets/ocm-service.clientId",
		ClientSecretFile:       "secrets/ocm-service.clientSecret",
		SelfTokenFile:          "secrets/ocm-service.token",
		Debug:                  false,
		EnableMock:             false,
		MockMode:               MockModeStubServer,
		StrimziOperatorAddonID: strimziOperatorAddonID,
		KasFleetshardAddonID:   kasFleetshardAddonID,
	}
}

func (c *OCMConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.ClientIDFile, "ocm-client-id-file", c.ClientIDFile, "File containing OCM API privileged account client-id")
	fs.StringVar(&c.ClientSecretFile, "ocm-client-secret-file", c.ClientSecretFile, "File containing OCM API privileged account client-secret")
	fs.StringVar(&c.SelfTokenFile, "self-token-file", c.SelfTokenFile, "File containing OCM API privileged offline SSO token")
	fs.StringVar(&c.BaseURL, "ocm-base-url", c.BaseURL, "The base URL of the OCM API, integration by default")
	fs.StringVar(&c.AmsUrl, "ams-base-url", c.AmsUrl, "The base URL of the AMS API, integration by default")
	fs.StringVar(&c.TokenURL, "ocm-token-url", c.TokenURL, "The base URL that OCM uses to request tokens, stage by default")
	fs.BoolVar(&c.Debug, "ocm-debug", c.Debug, "Debug flag for OCM API")
	fs.BoolVar(&c.EnableMock, "enable-ocm-mock", c.EnableMock, "Enable mock ocm clients")
	fs.StringVar(&c.MockMode, "ocm-mock-mode", c.MockMode, "Set mock type")
	fs.StringVar(&c.StrimziOperatorAddonID, "strimzi-operator-addon-id", c.StrimziOperatorAddonID, "The name of the Strimzi operator addon")
	fs.StringVar(&c.KasFleetshardAddonID, "kas-fleetshard-addon-id", c.KasFleetshardAddonID, "The name of the kas-fleetshard operator addon")
	fs.StringVar(&c.ClusterLoggingOperatorAddonID, "cluster-logging-operator-addon-id", "", "The name of the cluster logging operator addon. An empty string indicates that the operator should not be installed")
}

func (c *OCMConfig) ReadFiles() error {
	err := shared.ReadFileValueString(c.ClientIDFile, &c.ClientID)
	if err != nil {
		glog.Warning(err)
	}
	err = shared.ReadFileValueString(c.ClientSecretFile, &c.ClientSecret)
	if err != nil {
		glog.Warning(err)
	}
	err = shared.ReadFileValueString(c.SelfTokenFile, &c.SelfToken)
	if err != nil && (c.ClientSecret == "" || c.ClientID == "") {
		return err
	}

	return nil
}
