package session

import (
	"errors"
	"time"

	"github.com/google/gopacket"
)

type Session struct {
	SourceIP        string `json:"source_ip"`
	DestinationIP   string `json:"destination_ip"`
	SourcePort      string `json:"source_port"`
	DestinationPort string `json:"destination_port"`
	Protocol        string `json:"protocol"`
}

type Metrics struct {
	TotalPackets int       `json:"total_packets"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
}

func NewSession(packet gopacket.Packet) (*Session, error) {
	if packet.TransportLayer() == nil {
		return nil, errors.New("unknown protocol")
	}
	return &Session{
		SourceIP:        packet.NetworkLayer().NetworkFlow().Src().String(),
		DestinationIP:   packet.NetworkLayer().NetworkFlow().Dst().String(),
		SourcePort:      packet.TransportLayer().TransportFlow().Src().String(),
		DestinationPort: packet.TransportLayer().TransportFlow().Dst().String(),
		Protocol:        packet.TransportLayer().LayerType().String(),
	}, nil
}

func GetSessionMetrics(sessionStream map[Session]*Metrics, currentSession Session, packetTimestamp time.Time) (Session, *Metrics) {
	sessionMetrics, present := sessionStream[currentSession]
	if present {
		sessionMetrics.EndTime = packetTimestamp
		sessionMetrics.TotalPackets += 1
		return currentSession, sessionMetrics
	}
	mateSession := Session{
		SourceIP:        currentSession.DestinationIP,
		DestinationIP:   currentSession.SourceIP,
		SourcePort:      currentSession.DestinationPort,
		DestinationPort: currentSession.SourcePort,
		Protocol:        currentSession.Protocol,
	}
	sessionMetrics, present = sessionStream[mateSession]
	if present {
		sessionMetrics.EndTime = packetTimestamp
		sessionMetrics.TotalPackets += 1
		return mateSession, sessionMetrics
	} else {
		return currentSession, &Metrics{TotalPackets: 1, StartTime: packetTimestamp, EndTime: packetTimestamp}
	}
}
