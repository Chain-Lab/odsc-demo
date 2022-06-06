package main

import "github.com/urfave/cli"

func main() {
	app := cli.NewApp()

	app.Name = "osc-cli"

	app.Commands = []cli.Command{
		{
			Name:  "certificate",
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
}

func certificatePublish(c *cli.Context) error {
	data := c.String("data")
	secret := c.String("secret")
	eth_priv := c.String("eth_priv")

}
