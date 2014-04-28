package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	md5Cmd := &cobra.Command{
		Use:   "md5 <phrases>",
		Short: "encrypt using md5",
		Long:  "encrypt using md5",
		Run:   encrypt(md5.New()),
	}
	sha1Cmd := &cobra.Command{
		Use:   "sha1 <phrases>",
		Short: "encrypt using sha1",
		Run:   encrypt(sha1.New()),
	}
	sha256Cmd := &cobra.Command{
		Use:   "sha256 <phrases>",
		Short: "encrypt using sha256",
		Long:  "encrypt using sha256",
		Run:   encrypt(sha256.New()),
	}
	sha512Cmd := &cobra.Command{
		Use:   "sha512 <phrases>",
		Short: "encrypt using sha512",
		Long:  "encrypt using sha512",
		Run:   encrypt(sha512.New()),
	}

	crypterCmd := &cobra.Command{
		Use:   "crypter",
		Short: "An encryption program",
	}
	crypterCmd.AddCommand(md5Cmd)
	crypterCmd.AddCommand(sha1Cmd)
	crypterCmd.AddCommand(sha256Cmd)
	crypterCmd.AddCommand(sha512Cmd)
	err := crypterCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problems: %v\n", err)
	}
}

func encrypt(hashFunc hash.Hash) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintf(os.Stderr, "Need at least one phrase to encrypt\n")
			os.Exit(1)
		}
		for _, arg := range args {
			hashFunc.Write([]byte(arg))
			fmt.Printf("%x\n", hashFunc.Sum(nil))
			hashFunc.Reset()
		}
	}
}
