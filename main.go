package main

import (
	"encoding/hex"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 2 {

		packetData, err := hex.DecodeString(os.Args[1])

		fmt.Println("Raw Data:")
		for i := 0; i < len(packetData); i++ {
			fmt.Printf("%02X", packetData[i])

			if (i+1)%16 == 0 {
				fmt.Printf("\n")
			} else {
				fmt.Printf(" ")
			}
		}

		fmt.Printf("\n\n")

		if err != nil {
			log.Fatal(err)
		}
		packet := gopacket.NewPacket(packetData, layers.LayerTypeEthernet, gopacket.Default)

		fmt.Println("Decoded Packet:")
		fmt.Println(packet)
	} else {
		fmt.Printf("Usage: %s <byte string>\n", os.Args[0])
	}
}

// Local Variables:
// go-run-args: "28993A8F3C49001C73B57BC308004500009CA8BD40004006B5CE0A2A632C0A2A6450C1AC401D837BB94DFACDD51C5018FFFFB2BA00000072554F4D41454749414145413239333636134900CA5300000000000000040020D2580300373832323233202020202020202020204145474941414541202020204D2020202020202020202020202020202020202020202020202020202020202020202000000000000000000000590000000000"
// End:
