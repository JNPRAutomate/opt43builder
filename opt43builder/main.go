package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/jnprautomate/vendorcoder"
)

func main() {
	var AllOptions []*vendorcoder.VendorSubOption
	flagOptions := flag.String("options", "", "DHCP Option format code:value[,code:value,...]")
	delim := flag.String("delim", ",", "Specify the delimiter to seperate the options. Default is \",\"")
	flag.Parse()

	//Parse Vendor Options
	options := strings.Split(*flagOptions, *delim)
	for item := range options {
		optionSplit := strings.Split(options[item], ":")
		if len(optionSplit) == 2 {
			val, err := strconv.Atoi(optionSplit[0])
			if err != nil {
				fmt.Println(err)
			}
			newOption := &vendorcoder.VendorSubOption{}
			newOption.SetOption(uint8(val), []byte(optionSplit[1]))
			AllOptions = append(AllOptions, newOption)
		} else {
			flag.PrintDefaults()
			return
		}
	}

	//Compile all vendor options
	var OptByteSpString string
	var OptByteString string

	var OptHexSpString string
	var OptHexString string

	var OptHexFullSpString string
	var OptHexFullString string

	for item := range AllOptions {
		OptByteSpString = fmt.Sprintf("%s%s", OptByteSpString, AllOptions[item].ByteString(" "))
		OptByteString = fmt.Sprintf("%s%s", OptByteString, AllOptions[item].ByteString(""))

		OptHexSpString = fmt.Sprintf("%s%s", OptHexSpString, AllOptions[item].Hex(false, " "))
		OptHexString = fmt.Sprintf("%s%s", OptHexString, AllOptions[item].Hex(false, ""))

		OptHexFullSpString = fmt.Sprintf("%s%s", OptHexFullSpString, AllOptions[item].Hex(true, " "))
		OptHexFullString = fmt.Sprintf("%s%s", OptHexFullString, AllOptions[item].Hex(true, ""))

	}

	fmt.Printf("Decimal (Spaces): %s\n", OptByteSpString)
	fmt.Printf("Decimal: %s\n", OptByteString)
	fmt.Printf("Hex (Spaces): %s\n", OptHexSpString)
	fmt.Printf("Hex: %s\n", OptHexString)
	fmt.Printf("Hex Full(Spaces): %s\n", OptHexFullSpString)
	fmt.Printf("Hex Full: %s\n", OptHexFullString)
}
