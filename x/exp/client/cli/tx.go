package cli

import (
	"errors"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channelutils "github.com/cosmos/ibc-go/v4/modules/core/04-channel/client/utils"
	"github.com/notional-labs/craft/x/exp/types"
)

const (
	flagPacketTimeoutHeight    = "packet-timeout-height"
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	flagAbsoluteTimeouts       = "absolute-timeouts"
)

// NewTxCmd returns a root CLI command handler for all x/exp transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Exp transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(NewMintExpCmd())
	txCmd.AddCommand(NewBurnExpCmd())
	txCmd.AddCommand(NewSpendIbcAssetForExpCmd())
	txCmd.AddCommand(NewFundToExpModule())
	txCmd.AddCommand(NewAdjustDaoTokenPrice())
	return txCmd
}

func NewMintExpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mintexp [dao_member_address] [amount]",
		Short: `Mint exp for a whitelisted DAO member (only execute by DAO account)`,
		Long: `Mint exp for a DAO member which is whitelisted, this only execute by DAO account.
You can check the DAO account address by following command:  craftd q params subspace exp daoAccount
Also you can check whitelist by following command: 		craftd q exp whitelist`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			toAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgMintAndAllocateExp(clientCtx.GetFromAddress(), toAddr, coins)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewBurnExpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burnexp [dao_member_address]",
		Short: `Burn exp and exit the DAO.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnAndRemoveMember(clientCtx.GetFromAddress(), args[0])

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewSpendIbcAssetForExpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "spend [coins]",
		Short: `Spend IBC asset to receive exp and join the DAO.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			srcPort := "ibc-exp"
			srcChannel := "channel-0"
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}

			timeoutHeightStr, err := cmd.Flags().GetString(flagPacketTimeoutHeight)
			if err != nil {
				return err
			}
			timeoutHeight, err := clienttypes.ParseHeight(timeoutHeightStr)
			if err != nil {
				return err
			}

			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}

			absoluteTimeouts, err := cmd.Flags().GetBool(flagAbsoluteTimeouts)
			if err != nil {
				return err
			}

			// if the timeouts are not absolute, retrieve latest block height and block timestamp
			// for the consensus state connected to the destination port/channel
			if !absoluteTimeouts {
				_, height, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
				if err != nil {
					return err
				}
				if !timeoutHeight.IsZero() {
					absoluteHeight := height
					absoluteHeight.RevisionNumber += timeoutHeight.RevisionNumber
					absoluteHeight.RevisionHeight += timeoutHeight.RevisionHeight
					timeoutHeight = absoluteHeight
				}

				// Original: https://github.com/notional-labs/craft/commit/d631b248d044cbe64f637a41519e0500c884110b
				// can't be <0 bc unsigned int
				if timeoutTimestamp == 0 {
					return errors.New("local clock time is not greater than Jan 1st, 1970 12:00 AM")
				}

				now := uint64(time.Now().UnixNano())
				// consensusStateTimestamp := consensusState.GetTimestamp()

				// if the timeout is less than current time, then error.
				// though, is timeoutTimestamp in nano or normal seconds? if so we need to multiply
				if timeoutTimestamp < now {
					return errors.New("timeoutTimestamp is not greater than current local clock time")
				}
			}
			msg := types.NewMsgSpendIbcAssetToExp(clientCtx.GetFromAddress().String(), coins, timeoutHeight, timeoutTimestamp)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(flagPacketTimeoutHeight, types.DefaultRelativePacketTimeoutHeight, "Packet timeout block height. The timeout is disabled when set to 0-0.")
	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, types.DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes. The timeout is disabled when set to 0.")
	cmd.Flags().Bool(flagAbsoluteTimeouts, false, "Timeout flags are used as absolute timeouts.")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewFundToExpModule() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fund [coins]",
		Short: `Send [coins] to the exp module.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgFundExpPool(clientCtx.GetFromAddress().String(), coins)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewAdjustDaoTokenPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "adjust [price]",
		Short: `adjust the DAO exp token price to [price].`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			price, err := sdk.NewDecFromStr(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgAdjustDaoTokenPrice(clientCtx.GetFromAddress().String(), price)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewSendCoinsToDAO() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [amount]",
		Short: `send [amount] from module escrow to the DAO address.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgSendCoinsFromModuleToDAO(clientCtx.GetFromAddress().String(), coins)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// This func only for testing
// func NewJoinDaoCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "join [dao_member_address]",
// 		Short: `Burn exp and exit dao. Note, the'--from' flag is ignored as it is implied from [from_key_or_address].`,
// 		Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			msg := types.NewMsgJoinDaoByNonIbcAsset(clientCtx.GetFromAddress(), 100000, "craft10d07y265gmmuvt4z0w9aw880jnsr700jm5qjn0")

// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)

// 	return cmd
// }
