package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "BCC",
		Usage: "Encrypt or Decrypt with caeser cipher",
		Commands: []*cli.Command{
			{
				Name:    "encrypt",
				Aliases: []string{"e"},
				Usage:   "Encrypt text",
				Action: func(c *cli.Context) error {
					fmt.Printf("\033[38;5;202m ________  ________  ________   \n")
					fmt.Printf("|\\   __  \\|\\   ____\\|\\   ____\\    \n")
					fmt.Printf("\\ \\  \\|\\ /\\ \\  \\ ___| \\  \\___|    \n")
					fmt.Printf(" \\ \\   __  \\ \\  \\    \\ \\  \\       \n")
					fmt.Printf("  \\ \\  \\|\\  \\ \\  \\____\\ \\  \\____  \n")
					fmt.Printf("   \\ \\_______\\ \\_______\\ \\_______\\ \n")
					fmt.Printf("    \\|_______|\\|_______|\\|_______|\033[0m\n\n")

					fmt.Printf("[\033[38;5;126mENCRYPT\033[0m]\n")
					if c.NArg() > 0 {
						ls, _ := strconv.Atoi(c.Args().Get(1))
						fmt.Printf("%s \033[38;5;51m=>\033[0m %s \n", c.Args().Get(0), caeser_cipher_encrypt(c.Args().Get(0), ls))
					}

					return nil
				},
			},
			{
				Name:    "decrypt",
				Aliases: []string{"d"},
				Usage:   "Decrypt text",
				Action: func(c *cli.Context) error {
					fmt.Printf("\033[38;5;202m ________  ________  ________   \n")
					fmt.Printf("|\\   __  \\|\\   ____\\|\\   ____\\    \n")
					fmt.Printf("\\ \\  \\|\\ /\\ \\  \\ ___| \\  \\___|    \n")
					fmt.Printf(" \\ \\   __  \\ \\  \\    \\ \\  \\       \n")
					fmt.Printf("  \\ \\  \\|\\  \\ \\  \\____\\ \\  \\____  \n")
					fmt.Printf("   \\ \\_______\\ \\_______\\ \\_______\\ \n")
					fmt.Printf("    \\|_______|\\|_______|\\|_______|\033[0m\n\n")

					fmt.Printf("[\033[38;5;126mDECRYPT\033[0m]\n")
					if c.NArg() > 0 {
						ls, _ := strconv.Atoi(c.Args().Get(1))
						fmt.Printf("%s \033[38;5;51m=>\033[0m %s \n", c.Args().Get(0), caeser_cipher_decrypt(c.Args().Get(0), ls))
					}

					return nil
				},
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func caeser_cipher_encrypt(t string, ls int) string {
	var temp []string
	for _, l := range t {
		temp = append(temp, string(l+rune(ls)))
	}
	return strings.Join(temp, "")
}

func caeser_cipher_decrypt(t string, ls int) string {
	var temp []string
	for _, l := range t {
		temp = append(temp, string(l-rune(ls)))
	}
	return strings.Join(temp, "")
}

/*
func caeser_cipher_complex_decrypt(t string) string {
	var temp []string
	for i := 0; i < 26; i++ {
		for _, l := range t {
			temp = append(temp, string(l-rune(i)))
		}
		temp = append(temp, "\n")
	}
	return strings.Join(temp, "")
}
*/
