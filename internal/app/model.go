package app

import "github.com/saifsabir97/pcap-analyser/pkg/session"

type SessionDetails struct {
	Session session.Session `json:"session"`
	Metrics session.Metrics `json:"metrics"`
}
