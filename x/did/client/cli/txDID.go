package cli

import (
	"bufio"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	didcrypto "github.com/jim380/Re/x/did/client/crypto"
	"github.com/jim380/Re/x/did/types"
)

func CmdCreateDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-did",
		Short: "Create a DID",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			inBuf := bufio.NewReader(cmd.InOrStdin())

			mnemonic, bip39Passphrase, err := readBIP39ParamsFrom(viper.GetBool(flagInteractive), inBuf)
			if err != nil {
				return err
			}
			privKey, err := didcrypto.GenSecp256k1PrivKey(mnemonic, bip39Passphrase)
			if err != nil {
				return err
			}

			fromAddress := clientCtx.GetFromAddress()

			msg, err := newMsgCreateDID(fromAddress, privKey)

			if err != nil {
				return err
			}
			if err := savePrivKeyToKeyStore(msg.VerificationMethodId, privKey, inBuf); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Bool(flagInteractive, false, "Interactively prompt user for BIP39 mnemonic and passphrase")
	return cmd
}

func CmdUpdateDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-did [did] [key-id] [did-doc-path]",
		Short: "Update a DID Document",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			inBuf := bufio.NewReader(cmd.InOrStdin())

			did, err := types.ParseDID(args[0])
			if err != nil {
				return err
			}
			verificationMethodID, err := types.ParseVerificationMethodID(args[1], did)
			if err != nil {
				return err
			}
			// read an input file of DID document
			doc, err := readDIDDocFrom(args[2])
			if err != nil {
				return err
			}
			privKey, err := getPrivKeyFromKeyStore(verificationMethodID, inBuf)
			if err != nil {
				return err
			}
			// For proving that I know the private key. It signs on the DIDDocument.
			sig, err := signUsingCurrentSeq(clientCtx, did, privKey, &doc)
			if err != nil {
				return err
			}

			fromAddress := clientCtx.GetFromAddress()
			err = cmd.Flags().Set(flags.FlagFrom, fromAddress.String())
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDID(did, doc, verificationMethodID, sig, fromAddress.String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdDeactivateDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deactivate-did [did] [key-id]",
		Short: "Deactivate a DID Document",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			inBuf := bufio.NewReader(cmd.InOrStdin())

			did, err := types.ParseDID(args[0])
			if err != nil {
				return err
			}
			verificationMethodID, err := types.ParseVerificationMethodID(args[1], did)
			if err != nil {
				return err
			}
			privKey, err := getPrivKeyFromKeyStore(verificationMethodID, inBuf)
			if err != nil {
				return err
			}

			// For proving that I know the private key. It signs on the DIDDocument.
			doc := types.DIDDocument{
				Id: did,
			}
			sig, err := signUsingCurrentSeq(clientCtx, did, privKey, &doc)
			if err != nil {
				return err
			}

			fromAddress := clientCtx.GetFromAddress()
			msg := types.NewMsgDeactivateDID(did, verificationMethodID, sig, fromAddress.String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
