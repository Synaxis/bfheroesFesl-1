package fesl

import (
	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/Synaxis/bfheroesFesl/inter/mm"

	"github.com/Synaxis/bfheroesFesl/inter/network/codec"
	"github.com/sirupsen/logrus"
)

type reqStart struct {
	TXN        string `fesl:"TXN"`
	debugLevel string `fesl:"debugLevel"`
	Version    int    `fesl:"version"`
}

type statusPartition struct {
	ID        int    `fesl:"id"`
	Partition string `fesl:"partition"`
}

type ansStart struct {
	TXN string          `fesl:"TXN"`
	ID    statusPartition `fesl:"id"`
}

// Start handles pnow.Start
func (fm *Fesl) Start(event network.EvProcess) {
	logrus.Println("---START---")

	event.Client.Answer(&codec.Packet{
		Content: ansStart{
			TXN:  "Start",
			ID:    statusPartition{1, event.Process.Msg["eagames/bfwest-dedicated"]},
		},
		Send:    event.Process.HEX,
		Message: "pnow",
	})
	fm.Status(event)

}

type Status struct {
	TXN        string                 `fesl:"TXN"`
	ID         statusPartition        `fesl:"id"`
	State      string                 `fesl:"sessionState"`
	Properties map[string]interface{} `fesl:"props"`
}

type stGame struct {
	LobbyID int    `fesl:"lid"`
	Fit     int    `fesl:"fit"`
	GID     string `fesl:"gid"` //gameID to join
}

// Status comes after Start. tells info about desired server
func (fm *Fesl) Status(event network.EvProcess) {
	logrus.Println("--Status--")	
		
	// var gid string	
	// var err error

	// err = fm.db.stmtGetBookmark.QueryRow(event.Client.HashState.Get("uID")).Scan(&gid)
	// if err != nil {	
 	// 	logrus.Println("no game found for player")
	//  }	

	// continuos search
	for search := range mm.Games {
	gid := search
	gamesArray := []stGame{
		{
			GID:     gid,
			Fit:     1001,
			LobbyID: 1,
		},
	}		


	event.Client.Answer(&codec.Packet{
		Content: Status{
			TXN:    "Status",
			State:  "COMPLETE",
			ID:    statusPartition{1, event.Process.Msg["eagames/bfwest-dedicated"]},
			Properties: map[string]interface{}{
				"resultType": "JOIN",
				"sessionType": "FindServer",
				"games":      gamesArray},
		},
		Send:    0x80000000,
		Message: "pnow",
	})
}}
