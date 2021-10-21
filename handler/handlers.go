package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gpark1005/FlashCardsTeamOne/incomingdata"
)

type Service interface {
	PostNewInfo(Card incomingdata.Info) error
}

type InfoHandler struct {
	Svc Service
}

func NewInfoHandler(s Service) InfoHandler {
	return InfoHandler{
		Svc: s,
	}
}

func (ih InfoHandler) PostNewDeck(w http.ResponseWriter, r *http.Request) {

	deck := incomingdata.Deck{}

	err := json.NewDecoder(r.Body).Decode(&deck)
	if err != nil {
		log.Fatal(err)
	}

}

func (ih InfoHandler) PostNewInfo(w http.ResponseWriter, r *http.Request) {

	card := incomingdata.Info{}

	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		log.Fatal(err)
	}

	err = ih.Svc.PostNewInfo(card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // 
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
