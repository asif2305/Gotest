package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Set the device name to capture packets from
	deviceName := "enp0s3"

	// Open the device
	handle, err := pcap.OpenLive(deviceName, 1514, false, pcap.Snaplen_32)
	if err != nil {
		fmt.Println("Error opening device:", err)
		return
	}
	defer handle.Close()

	// Set the filter to capture only ARP packets
	filter := "arp"
	err = handle.SetFilter(filter)
	if err != nil {
		fmt.Println("Error setting filter:", err)
		return
	}

	// Start capturing packets
	packetSource := gopacket.NewPacketSource(handle, gopacket.NewPacketSource.NoCopy)
	packetSource.Packets().ForEach(func(packet gopacket.Packet) {
		// Check if the packet is an ARP packet
		if arpLayer := packet.Layer(gopacket.LayerTypeARP); arpLayer != nil {
			// Parse the ARP packet
			arp := arpLayer.(*gopacket.ARP)
			fmt.Println("ARP packet:")
			fmt.Printf("Sender Hardware Address: %s\n", arp.SourceHardwareAddr)
			fmt.Printf("Sender IP Address: %s\n", arp.SenderProtAddr)
			fmt.Printf("Target Hardware Address: %s\n", arp.TargetHardwareAddr)
			fmt.Printf("Target IP Address: %s\n", arp.TargetProtAddr)
			fmt.Println("Operation:", arp.Operation)
		}
	})
}
