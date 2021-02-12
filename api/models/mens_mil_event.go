package models

import (
	"api-horus/api/event_db"
	"api-horus/api/helpers"
	"errors"
	"strings"
	"time"
)

var (
	ValidClasifSeg                 = []string{"publico", "reservado", "confidencial", "secreto"}
	ValidPrecedencia               = []string{"rutina", "prioridad", "inmediato", "flash"}
	InvalidClasifSegMensMilEvent   = errors.New("MensMilEvent Model: Invalid Clasificacion de Seguridad")
	InvalidPrecedenciaMensMilEvent = errors.New("MensMilEvent Model: Invalid Precedencia")
	InvalidCifradoMensMilEvent     = errors.New("MensMilEvent Model: Invalid value for cifrado")
	InvalidDestinoMensMilEvent     = errors.New("MensMilEvent Model: Invalid value for destino")
	InvalidOrigenMensMilEvent      = errors.New("MensMilEvent Model: value for origen")
	InvalidEventoMensMilEvent      = errors.New("MensMilEvent Model: Invalid value for evento")
	InvalidTimeMensMilEvent        = errors.New("MensMilEvent Model: Invalid Time Value, out of +- 600 sec range")
)

type MensMilEvent struct {
	ClasifSeg   string `json:"clasificacion" binding:"required"`
	Precedencia string `json:"precedencia" binding:"required"`
	Cifrado     bool   `json:"cifrado" binding:"required"`
	Destino     string `json:"destino" binding:"required"`
	Origen      string `json:"origen" binding:"required"`
	Evento      string `json:"evento" binding:"required"`
	Time        int64  `json:"timestamp"`
}

func (e MensMilEvent) Write() (*MensMilEvent, error) {
	client := event_db.GetEventDB()
	stmt := `INSERT INTO mm_events (clasif_seg, precedencia, cifrado, destino, origen, evento, gfh) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := client.Exec(stmt, e.ClasifSeg, e.Precedencia, e.Cifrado, e.Destino, e.Origen, e.Evento, e.Time)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e MensMilEvent) Validate() error {
	if !helpers.ValidString(e.ClasifSeg, ValidClasifSeg) {
		return InvalidClasifSegMensMilEvent
	}
	if !helpers.ValidString(e.Precedencia, ValidPrecedencia) {
		return InvalidPrecedenciaMensMilEvent
	}
	if strings.TrimSpace(e.Destino) == "" {
		return InvalidDestinoMensMilEvent
	}
	if strings.TrimSpace(e.Origen) == "" {
		return InvalidOrigenMensMilEvent
	}
	if strings.TrimSpace(e.Evento) == "" {
		return InvalidEventoMensMilEvent
	}
	diffTime := time.Now().Unix() - e.Time
	if diffTime > 600 || diffTime < -600 {
		return InvalidTime
	}
	return nil
}