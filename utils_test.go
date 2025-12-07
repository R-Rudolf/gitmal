package main

import (
	"testing"

	"github.com/antonmedv/gitmal/pkg/git"
)

func TestRefToFileName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"main", "main"},
		{"master", "master"},
		{"release/v1.0", "release-v1.0"},
		{"feature/add-login", "feature-add-login"},
		{"bugfix\\windows\\path", "bugfix-windows-path"},
		{"1.0.0", "1.0.0"},
		{"1.x", "1.x"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := refToFileName(git.Ref(tt.in))
			if got != tt.want {
				t.Fatalf("refToFileName(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
