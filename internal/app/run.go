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
		actualSession, metrics := session.GetSessionMetrics(sessionStream, *currentSession,
			packet.Metadata().Timestamp, packet.Metadata().CaptureLength)
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
		[]string{"Client 1 IP", "Client 2 IP",
			"Client 1 Port", "Client 2 Port",
			"Protocol", "Total Packets Transferred",
			"Start Time", "End Time",
			"Data Transferred (in Bytes)",
		},
	)
	for _, currentSessionDetails := range sessions {
		matrix = append(
			matrix,
			[]string{currentSessionDetails.Session.Client1IP, currentSessionDetails.Session.Client2IP,
				currentSessionDetails.Session.Client1Port, currentSessionDetails.Session.Client2Port,
				currentSessionDetails.Session.Protocol, strconv.Itoa(currentSessionDetails.Metrics.TotalPackets),
				currentSessionDetails.Metrics.StartTime.String(), currentSessionDetails.Metrics.EndTime.String(),
				strconv.Itoa(currentSessionDetails.Metrics.TotalData),
			},
		)
	}
	return matrix
}
