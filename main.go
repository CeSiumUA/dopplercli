package main

import (
	"bufio"
	"fmt"
	"github.com/CeSiumUA/dopplerptp/networking"
	"os"
)

func main() {
	err := networking.StartListener()
	if err != nil {
		fmt.Printf("error occured during listener creation: %s", err)
	}
	natAddresses, err := networking.GetLocalNetworkAddresses()
	if err != nil {
		fmt.Printf("error occured local addresses discovery: %s", err)
	}
	networking.ExploreAddresses(natAddresses)
	fmt.Printf("Enter q to exit:")
	cliReader := bufio.NewReader(os.Stdin)
	cmd, err := cliReader.ReadString('\n')
	if err != nil {
		fmt.Printf("error occured while reading cmd value: %s", err)
	}
	if cmd == "q" {
		return
	}
}
