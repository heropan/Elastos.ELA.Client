package info

import (
	"fmt"
	"strconv"

	"ELAClient/rpc"
	"github.com/urfave/cli"
)

func infoAction(c *cli.Context) error {
	if c.NumFlags() == 0 {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	if c.Bool("version") {
		result, err := rpc.CallAndUnmarshal("getversion")
		if err != nil {
			fmt.Println("error: get node version failed, ", err)
			return err
		}
		fmt.Println(result)
		return nil
	}

	if c.Bool("connections") {
		result, err := rpc.CallAndUnmarshal("getconnectioncount")
		if err != nil {
			fmt.Println("error: get node connections failed, ", err)
			return err
		}
		fmt.Println(result)
		return nil
	}

	if c.Bool("neighbor") {
		result, err := rpc.CallAndUnmarshal("getneighbor")
		if err != nil {
			fmt.Println("error: get node neighbors info failed, ", err)
			return err
		}
		fmt.Println(result)
		return nil
	}

	if c.Bool("state") {
		result, err := rpc.CallAndUnmarshal("getnodestate")
		if err != nil {
			fmt.Println("error: get node state info failed, ", err)
			return err
		}
		fmt.Println(result)
		return nil
	}

	if c.Bool("blockcount") {
		result, err := rpc.CallAndUnmarshal("getblockcount")
		if err != nil {
			fmt.Println("error: get block count failed, ", err)
			return err
		}
		fmt.Println(result)
		return nil
	}

	if height := c.Uint("getblockhash"); height >= 0 {
		result, err := rpc.CallAndUnmarshal("getblockhash", height)
		if err != nil {
			fmt.Println("error: get block hash failed, ", err)
			return err
		}
		fmt.Println(result.(string))
		return nil
	}

	if param := c.String("getblock"); param != "" {
		height, err := strconv.ParseInt(param, 10, 64)

		var result interface{}
		if err == nil {
			result, err = rpc.CallAndUnmarshal("getblock", height)
		} else {
			result, err = rpc.CallAndUnmarshal("getblock", param)
		}
		if err != nil {
			fmt.Println("error: get block failed, ", err)
			return err
		}
		fmt.Println(result)
		return nil
	}

	if param := c.String("gettransaction"); param != "" {
		result, err := rpc.CallAndUnmarshal("getrawtransaction", param)
		if err != nil {
			fmt.Println("error: get transaction failed, ", err)
			return err
		}
		fmt.Println(result)
		return nil
	}

	if c.Bool("bestblockhash") {
		result, err := rpc.CallAndUnmarshal("getbestblockhash")
		if err != nil {
			fmt.Println("error: get last block hash failed, ", err)
			return err
		}
		fmt.Println(result)
		return nil
	}

	// TODO format transactions in mem pool
	if c.Bool("showtxpool") {
		result, err := rpc.CallAndUnmarshal("getrawmempool")
		if err != nil {
			fmt.Println("error: get transaction pool failed, ", err)
			return err
		}
		fmt.Println(result)
		return nil
	}

	return nil
}

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:        "info",
		Usage:       "show blockchain information",
		Description: "With ela-cli info, you could look up blocks, transactions, etc.",
		ArgsUsage:   "[args]",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "version, v",
				Usage: "version of the connected node",
			},
			cli.BoolFlag{
				Name:  "connections, cs",
				Usage: "how many connections are holding by the connected node",
			},
			cli.BoolFlag{
				Name:  "neighbor, nb",
				Usage: "neighbor information of the connected node",
			},
			cli.BoolFlag{
				Name:  "state, s",
				Usage: "get the connected node's state",
			},
			cli.BoolFlag{
				Name:  "blockcount, bc",
				Usage: "current blocks in the blockchain",
			},
			cli.UintFlag{
				Name:  "getblockhash, gbh",
				Usage: "query a block's hash with it's height",
			},
			cli.StringFlag{
				Name:  "getblock, gb",
				Usage: "query a block with height or it's hash",
			},
			cli.BoolFlag{
				Name:  "bestblockhash, bbh",
				Usage: "get the latest block's hash",
			},
			cli.StringFlag{
				Name:  "gettransaction, gt",
				Usage: "query a transaction with it's hash",
			},
			cli.BoolFlag{
				Name:  "showtxpool, stp",
				Usage: "show the transactions in node's transaction pool",
			},
		},
		Action: infoAction,
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			return cli.NewExitError(err, 1)
		},
	}
}
