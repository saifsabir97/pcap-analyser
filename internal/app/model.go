package app

import "hawk/pkg/session"

type SessionDetails struct {
	Session session.Session `json:"session"`
	Metrics session.Metrics `json:"metrics"`
}
