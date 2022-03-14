package cmd

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"

	"golang.org/x/oauth2/clientcredentials"
)

var (
	daClientID     string
	daClientSecret string
	daTokenURL     string
)

func init() {
	dryAuthCmd.Flags().StringVarP(&daClientID, "client_id", "a", os.Getenv("BIOMETRICSCLOUD_PI_CLIENT_ID"), "-")
	// dryAuthCmd.MarkFlagRequired("client_id")
	dryAuthCmd.Flags().StringVarP(&daClientSecret, "client_secret", "b", os.Getenv("BIOMETRICSCLOUD_PI_HOST"), "-")
	// dryAuthCmd.MarkFlagRequired("client_secret")
	dryAuthCmd.Flags().StringVarP(&daTokenURL, "token_url", "c", os.Getenv("BIOMETRICSCLOUD_PI_HOST"), "-")
	// dryAuthCmd.MarkFlagRequired("token_url")
	rootCmd.AddCommand(dryAuthCmd)
}

var dryAuthCmd = &cobra.Command{
	Use:   "dry-auth",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if daClientID == "" {
			log.Fatal("Missing `client_id` value.")
		}
		if daClientSecret == "" {
			log.Fatal("Missing `client_secret` value.")
		}
		if daTokenURL == "" {
			log.Fatal("Missing `token_url` value.")
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
