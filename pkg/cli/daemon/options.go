package daemon

import (
	"github.com/ylallemant/panopticon/pkg/api"
	"os"
)

// Current is the defacto Config used.
var Current = api.ConfigWithDefaults()

func LoadEnvVarsToConfig() {
	if key, ok := os.LookupEnv("CLUSTER_POLICY_UPDATE_KEY"); len(Current.ClusterPolicyUpdateKey) == 0 && ok {
		Current.ClusterPolicyUpdateKey = key
	}
}
