// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package upgrade_test

import (
	"testing"

	"github.com/vmware/cluster-api-upgrade-tool/pkg/upgrade"
)

func TestValidArgs(t *testing.T) {
	testcases := []struct {
		name string
		cfg  upgrade.Config
	}{
		{
			name: "simple",
			cfg: upgrade.Config{
				TargetCluster: upgrade.TargetClusterConfig{
					UpgradeScope: upgrade.ControlPlaneScope,
				},
				KubernetesVersion: "v1.12.1",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if err := upgrade.ValidateArgs(tc.cfg); err != nil {
				t.Fatalf("%+v", err)
			}
		})
	}
}

// TODO if we end up with a set of specific errors we can test them here.
// For now this Invalid test just ensures we raise errors with specific configs.

func TestInvalidArgs(t *testing.T) {
	testcases := []struct {
		name string
		cfg  upgrade.Config
	}{
		{
			name: "empty configuration",
			cfg:  upgrade.Config{},
		},
		{
			name: "invalid cluster upgrade scope",
			cfg: upgrade.Config{
				TargetCluster: upgrade.TargetClusterConfig{
					UpgradeScope: "some-invalid-upgrade-scope",
				},
			},
		},
		{
			name: "invalid kubernetes version",
			cfg: upgrade.Config{
				KubernetesVersion: "some-bad-version",
				TargetCluster: upgrade.TargetClusterConfig{
					UpgradeScope: upgrade.ControlPlaneScope,
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if err := upgrade.ValidateArgs(tc.cfg); err == nil {
				t.Fatal("Expected an error but didn't receive one")
			}
		})
	}

}
