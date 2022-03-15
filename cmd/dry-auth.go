package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"golang.org/x/oauth2/clientcredentials"
)

var (
	daClientID     string
	daClientSecret string
	daTokenURL     string
)

func init() {
	dryAuthCmd.Flags().StringVarP(&daClientID, "client_id", "a", devClientID, "Override the existing client ID with new value.")
	// dryAuthCmd.MarkFlagRequired("client_id")
	dryAuthCmd.Flags().StringVarP(&daClientSecret, "client_secret", "b", devClientSecret, "Override the existing client secret with new value.")
	// dryAuthCmd.MarkFlagRequired("client_secret")
	dryAuthCmd.Flags().StringVarP(&daTokenURL, "token_url", "c", devTokenURL, "Override the existing token URL with new value.")
	// dryAuthCmd.MarkFlagRequired("token_url")
	rootCmd.AddCommand(dryAuthCmd)
}

var dryAuthCmd = &cobra.Command{
	Use:   "dry-auth",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Process the overrides.
		if daClientID == "" {
			daClientID = devClientID
		}
		if daClientSecret == "" {
			daClientSecret = devClientSecret
		}
		if daTokenURL == "" {
			daTokenURL = devTokenURL
		}

		log.Println("Beginning dry run of client credential based authorization...")

		cfg := clientcredentials.Config{
			ClientID:     daClientID,
			ClientSecret: daClientSecret,
			Scopes:       []string{"all"},
			TokenURL:     daTokenURL,
		}

		// https://github.com/go-oauth2/oauth2/blob/b208c14e621016995debae2fa7dc20c8f0e4f6f8/example/client/client.go#L116
		token, err := cfg.Token(context.Background())
		if err != nil {
			log.Println("Dry run failed with error:")
			log.Fatal(err)
		}

		// NOTE: https://pkg.go.dev/golang.org/x/oauth2#Token
		// log.Println("AccessToken", token.AccessToken)
		// log.Println("TokenType", token.TokenType)
		// log.Println("RefreshToken", token.RefreshToken)
		// log.Println("Expiry", token.Expiry)
		// log.Println("UserID", token.Extra("custom_parameter"))
		if token != nil {
			log.Println("Dry run was a success!")
		}
	},
}
