package app

import (
	"strconv"

	"github.com/saifsabir97/pcap-analyser/pkg/session"
)

func (c *Client) Run() [][]string {
	sessionStream := map[session.Session]*session.Metrics{}
	for packet := range c.packetSource.Packets() {
		currentSession, err := session.NewSession(packet)
		if err != nil {
			continue
		}
		actualSession, metrics := session.GetSessionMetrics(sessionStream, *currentSession, packet.Metadata().Timestamp)
		sessionStream[actualSession] = metrics
	}
	var sessions []SessionDetails
	for currentSession, metric := range sessionStream {
		currentSessionDetails := SessionDetails{
			Session: currentSession,
			Metrics: *metric,
		}
		sessions = append(sessions, currentSessionDetails)
	}
	return transformSessionToMatrix(sessions)
}

func transformSessionToMatrix(sessions []SessionDetails) [][]string {
	var matrix [][]string
	matrix = append(
		matrix,
		[]string{"Source IP", "Destination IP",
			"Source Port", "Destination Port",
			"Protocol", "Total Packets Transferred",
			"Start Time", "End Time",
		},
	)
	for _, currentSessionDetails := range sessions {
		matrix = append(
			matrix,
			[]string{currentSessionDetails.Session.SourceIP, currentSessionDetails.Session.DestinationIP,
				currentSessionDetails.Session.SourcePort, currentSessionDetails.Session.DestinationPort,
				currentSessionDetails.Session.Protocol, strconv.Itoa(currentSessionDetails.Metrics.TotalPackets),
				currentSessionDetails.Metrics.StartTime.String(), currentSessionDetails.Metrics.EndTime.String(),
			},
		)
	}
	return matrix
}
