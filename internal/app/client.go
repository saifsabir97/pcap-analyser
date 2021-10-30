package app

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Client struct {
	packetSource *gopacket.PacketSource
}

func NewClient(filePath string) (*Client, error) {
	handle, err := pcap.OpenOffline(filePath)
	if err != nil {
		return nil, err
	}
	return &Client{
		packetSource: gopacket.NewPacketSource(handle, handle.LinkType()),
	}, nil
}
