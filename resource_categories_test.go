// Copyright 2024, Pulumi Corporation.  All rights reserved.

package main

import (
	"testing"
)

func TestGetResourceKind(t *testing.T) {
	tests := []struct {
		resourceType string
		expectedKind ResourceKind
	}{
		{"aws:acmpca", Security},
		{"aws:aps", Observability},
		{"aws:arczonalshift", Network},
		{"aws:notfound", NotFound},           // not found
		{"aws-native:acmpca", Security},      // aws-native
		{"unsupported:service", ""},          // unsupported provider
		{"azure:appinsights", Observability}, // azure
		{"kubernetes:apps", Container},       // kubernetes
		{"cloud:bucket", Storage},            // cloud
		{"oci:ContainerEngine", Container},   // oci: by default uses PascalCase
		// upper case
		{"AWS:ACMPCA", Security},
		{"AZURE:APPINSIGHTS", Observability},
		{"KUBERNETES:APPS", Container},
	}

	for _, test := range tests {
		t.Run(test.resourceType, func(t *testing.T) {
			kind := GetResourceKind(test.resourceType)
			if kind != test.expectedKind {
				t.Errorf("expected %v, got %v", test.expectedKind, kind)
			}
		})
	}
}
