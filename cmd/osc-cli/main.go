package main

import (
	"fmt"
	"github.com/decision2016/osc/internal/rpc"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "osc-cli"

	app.Commands = []cli.Command{
		{
			Name:  "wallet",
			Usage: "osc-cli wallet ...",
			Subcommands: []cli.Command{
				{
					Name: "create",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "phrase",
						},
						cli.StringFlag{
							Name: "out",
						},
					},
					Action: walletCreate,
				},
				{
					Name: "upload",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "file-path",
						},
					},
					Action: walletUpload,
				},
			},
		},
		{
			Name:  "data",
			Usage: "osc-cli certificate ...",
			Subcommands: []cli.Command{
				{
					Name: "publish",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "data",
						},
						cli.StringFlag{
							Name: "secret",
						},
						cli.StringFlag{
							Name: "eth_priv",
						},
					},
					Action: certificatePublish,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}

func certificatePublish(c *cli.Context) error {
	//data := c.String("data")
	//secret := c.String("secret")
	//eth_priv := c.String("eth_priv")

	return nil

}

func walletCreate(c *cli.Context) (err error) {
	phrase := c.String("phrase")
	out := c.String("out")

	var outputPath string

	if "" == phrase {
		log.Fatal("Phrase is empty.")
		return
	}

	if "" == out {
		outputPath = "./keystore"
	} else {
		outputPath = out
	}

	password := phrase

	ks := keystore.NewKeyStore(outputPath, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generate address: ", account.Address.Hex())

	return
}

func walletUpload(c *cli.Context) (err error) {
	file_path := c.String("file-path")

	if "" == file_path {
		return
	}

	keystoreJson, err := ioutil.ReadFile(file_path)

	if err != nil {
		log.Fatal(err)
		return err
	}

	rpc.UploadKeystoreJson(keystoreJson)

	return
}
