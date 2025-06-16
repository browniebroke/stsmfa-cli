package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

var (
	profile    string
	mfaProfile string
	version    = "dev" // Set by build ldflags
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "awsmfa [token]",
		Short:   "Get a session token using AWS STS and a 2FA token",
		Long:    "A small CLI to help with creating AWS profile for MFA protected sessions",
		Args:    cobra.ExactArgs(1),
		RunE:    runCommand,
		Version: version,
	}

	rootCmd.Flags().StringVarP(&profile, "profile", "p", "default", "The profile to use for obtaining the session token")
	rootCmd.Flags().StringVar(&mfaProfile, "mfa-profile", "", "The profile to write the session data to. Defaults to <profile>-mfa")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runCommand(cmd *cobra.Command, args []string) error {
	token := args[0]

	// Get AWS credentials file path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		color.Red("Error getting home directory: %v", err)
		return err
	}

	credsPath := filepath.Join(homeDir, ".aws", "credentials")
	if _, err := os.Stat(credsPath); os.IsNotExist(err) {
		color.Red("Credentials file not found at %s", credsPath)
		return err
	}

	// Parse credentials file
	cfg, err := ini.Load(credsPath)
	if err != nil {
		color.Red("Error reading credentials file: %v", err)
		return err
	}

	// Check if profile exists
	section, err := cfg.GetSection(profile)
	if err != nil {
		color.Red("Profile %s not found in credentials file %s", profile, credsPath)
		return err
	}

	// Check if mfa_serial is configured
	mfaSerial := section.Key("mfa_serial").String()
	if mfaSerial == "" {
		color.Red("Profile %s does not have an mfa_serial configured", profile)
		return fmt.Errorf("mfa_serial not configured")
	}

	// Load AWS config with the specified profile
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile),
	)
	if err != nil {
		color.Red("Error loading AWS config: %v", err)
		return err
	}

	// Create STS client
	stsClient := sts.NewFromConfig(awsCfg)

	// Get session token
	input := &sts.GetSessionTokenInput{
		SerialNumber: aws.String(mfaSerial),
		TokenCode:    aws.String(token),
	}

	result, err := stsClient.GetSessionToken(context.TODO(), input)
	if err != nil {
		color.Red("Error getting session token: %v", err)
		return err
	}

	// Determine MFA profile name
	if mfaProfile == "" {
		mfaProfile = profile + "-mfa"
	}

	// Add or update MFA profile section
	mfaSection, err := cfg.NewSection(mfaProfile)
	if err != nil {
		// Section might already exist, get it instead
		mfaSection, err = cfg.GetSection(mfaProfile)
		if err != nil {
			color.Red("Error creating/getting MFA profile section: %v", err)
			return err
		}
	}

	// Set credentials
	mfaSection.Key("aws_access_key_id").SetValue(*result.Credentials.AccessKeyId)
	mfaSection.Key("aws_secret_access_key").SetValue(*result.Credentials.SecretAccessKey)
	mfaSection.Key("aws_session_token").SetValue(*result.Credentials.SessionToken)

	// Save credentials file
	err = cfg.SaveTo(credsPath)
	if err != nil {
		color.Red("Error saving credentials file: %v", err)
		return err
	}

	color.Green("All written to %s profile in %s", mfaProfile, credsPath)
	return nil
}
