package python

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pyNameTests = []struct {
	input    string
	expected string
	legacy   string
}{
	{"kubeletConfigKey", "kubelet_config_key", "kubelet_config_key"},
	{"podCIDR", "pod_cidr", "pod_cidr"},
	{"podCidr", "pod_cidr", "pod_cidr"},
	{"podCIDRs", "pod_cidrs", "pod_cid_rs"},
	{"podIPs", "pod_ips", "pod_i_ps"},
	{"nonResourceURLs", "non_resource_urls", "non_resource_ur_ls"},
	{"someTHINGsAREWeird", "some_things_are_weird", "some_thin_gs_are_weird"},
	{"podCIDRSet", "pod_cidr_set", "pod_cidr_set"},
	{"Sha256Hash", "sha256_hash", "sha256_hash"},
	{"SHA256Hash", "sha256_hash", "sha256_hash"},

	// Handle embedded underscores
	{"_NO_NAME", "_no_name", "_no_name"},
	{"_no_NAME", "_no_name", "_no_name"},
	{"_NO_name", "_no_name", "_no_name"},
	{"_no_name", "_no_name", "_no_name"},
	{"NO_NAME", "no_name", "no_name"},
	{"no_NAME", "no_name", "no_name"},
	{"NO_name", "no_name", "no_name"},
	{"no_name", "no_name", "no_name"},

	// PyName should return the legacy name for these:
	{"openXJsonSerDe", "open_x_json_ser_de", "open_x_json_ser_de"},
	{"GetPublicIPs", "get_public_i_ps", "get_public_i_ps"},
	{"GetUptimeCheckIPs", "get_uptime_check_i_ps", "get_uptime_check_i_ps"},
}

func TestPyName(t *testing.T) {
	t.Parallel()

	for _, tt := range pyNameTests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			result := PyName(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPyNameLegacy(t *testing.T) {
	t.Parallel()

	for _, tt := range pyNameTests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			result := PyNameLegacy(tt.input)
			assert.Equal(t, tt.legacy, result)
		})
	}
}
