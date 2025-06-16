package main

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/ini.v1"
)

func TestGetCredentialsPath(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}

	expected := filepath.Join(homeDir, ".aws", "credentials")
	actual := filepath.Join(homeDir, ".aws", "credentials")

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestParseCredentialsFile(t *testing.T) {
	// Create a temporary credentials file
	tmpDir := t.TempDir()
	credsFile := filepath.Join(tmpDir, "credentials")

	content := `[default]
aws_access_key_id = AKIAIOSFODNN7EXAMPLE
aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
mfa_serial = arn:aws:iam::123456789012:mfa/test-user

[test-profile]
aws_access_key_id = AKIAIOSFODNN7EXAMPLE2
aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY2
mfa_serial = arn:aws:iam::123456789012:mfa/test-user2
`

	err := os.WriteFile(credsFile, []byte(content), 0600)
	if err != nil {
		t.Fatal(err)
	}

	// Test parsing
	cfg, err := ini.Load(credsFile)
	if err != nil {
		t.Fatal(err)
	}

	// Test that sections exist
	defaultSection, err := cfg.GetSection("default")
	if err != nil {
		t.Fatal(err)
	}

	testSection, err := cfg.GetSection("test-profile")
	if err != nil {
		t.Fatal(err)
	}

	// Test that mfa_serial is readable
	defaultMfaSerial := defaultSection.Key("mfa_serial").String()
	if defaultMfaSerial != "arn:aws:iam::123456789012:mfa/test-user" {
		t.Errorf("Expected mfa_serial to be 'arn:aws:iam::123456789012:mfa/test-user', got '%s'", defaultMfaSerial)
	}

	testMfaSerial := testSection.Key("mfa_serial").String()
	if testMfaSerial != "arn:aws:iam::123456789012:mfa/test-user2" {
		t.Errorf("Expected mfa_serial to be 'arn:aws:iam::123456789012:mfa/test-user2', got '%s'", testMfaSerial)
	}
}
