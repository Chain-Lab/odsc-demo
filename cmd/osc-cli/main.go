package main

import (
	"encoding/base64"
	"fmt"
	"github.com/decision2016/osc/internal/rpc"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"google.golang.org/protobuf/types/known/structpb"
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
				//{
				//	Name: "upload",
				//	Flags: []cli.Flag{
				//		cli.StringFlag{
				//			Name: "file-path",
				//		},
				//	},
				//	Action: walletUpload,
				//},
			},
		},
		{
			Name:  "data",
			Usage: "osc-cli data ...",
			Subcommands: []cli.Command{
				{
					Name: "create",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "data",
						},
					},
					Action: dataCreate,
				},
				{
					Name: "modify",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "data",
						},
						cli.StringFlag{
							Name: "key",
						},
					},
					Action: dataModify,
				},
			},
		},
		{
			Name:  "permission",
			Usage: "osc-cli permission ...",
			Subcommands: []cli.Command{
				{
					Name: "add",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "address",
						},
					},
					Action: permissionAdd,
				},
				{
					Name: "revoke",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "address",
						},
					},
					Action: permissionRevoke,
				},
			},
		},
		{
			Name: "genesis",
			Usage: "osc-cli genesis",
			Action: genesis,
		},
		{
			Name: "test",
			Usage: "osc-cli test",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "id",
				},
			},
			Action: test,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}

func genesis(c *cli.Context) error {
	params := map[string]*structpb.Value{}
	rpc.CallLocalRPC("genesis", "", &params)

	return nil
}

func dataCreate(c *cli.Context) error {
	filePath := c.String("data")
	data, err := os.ReadFile(filePath)

	if err != nil {
		logrus.Fatal(err)
	}

	dataBase64String := base64.StdEncoding.EncodeToString(data)
	params := map[string]*structpb.Value{
		"data": {Kind: &structpb.Value_StringValue{StringValue: dataBase64String}},
	}

	rpc.CallLocalRPC("data", "publish", &params)

	return nil
}

func dataModify(c *cli.Context) error {
	filePath := c.String("data")
	key := c.String("key")

	data, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	dataBase64String := base64.StdEncoding.EncodeToString(data)
	params := map[string]*structpb.Value{
		"data": {Kind: &structpb.Value_StringValue{StringValue: dataBase64String}},
		"key":  {Kind: &structpb.Value_StringValue{StringValue: key}},
	}

	rpc.CallLocalRPC("data", "modify", &params)

	return nil
}

func permissionAdd(c *cli.Context) error {
	address := c.String("address")

	params := map[string]*structpb.Value{
		"address": {Kind: &structpb.Value_StringValue{StringValue: address}},
	}

	rpc.CallLocalRPC("permission", "add", &params)

	return nil
}

func permissionRevoke(c *cli.Context) error {
	address := c.String("address")

	params := map[string]*structpb.Value{
		"address": {Kind: &structpb.Value_StringValue{StringValue: address}},
	}

	rpc.CallLocalRPC("permission", "revoke", &params)

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

//func walletUpload(c *cli.Context) (err error) {
//	file_path := c.String("file-path")
//
//	if "" == file_path {
//		return
//	}
//
//	keystoreJson, err := ioutil.ReadFile(file_path)
//
//	if err != nil {
//		log.Fatal(err)
//		return err
//	}
//
//	rpc.UploadKeystoreJson(string(keystoreJson))
//
//	return
//}

func test(c *cli.Context) (err error) {
	id := c.String("id")

	filepath := "./files/" + id
	fileBytes, err := os.ReadFile(filepath)

	if err != nil {
		logrus.Fatal(err)
	}

	fileBase64 := base64.StdEncoding.EncodeToString(fileBytes)

	//params := map[string]*structpb.Value{
	//	"data": {Kind: &structpb.Value_StringValue{StringValue: fileBase64}},
	//}
	//
	//rpc.CallLocalRPC("data", "create", &params)
	fmt.Println(fileBase64)

	return
}