package models

import (
	"api/event_db"
	"api/helpers"
	"errors"
	"strings"
	"time"
)

var (
	ValidClasifSeg                       = []string{"publico", "reservado", "confidencial", "secreto"}
	ValidPrecedencia                     = []string{"rutina", "prioridad", "inmediato", "flash"}
	ErrFooInvalidClasifSegMensMilEvent   = errors.New("MensMilEvent Model: Invalid Clasificacion de Seguridad")
	ErrFooInvalidPrecedenciaMensMilEvent = errors.New("MensMilEvent Model: Invalid Precedencia")
	ErrFooInvalidNroMMMensMilEvent       = errors.New("MensMilEvent Model: Invalid value for nro mm")
	ErrFooInvalidDestinoMensMilEvent     = errors.New("MensMilEvent Model: Invalid value for destino")
	ErrFooInvalidOrigenMensMilEvent      = errors.New("MensMilEvent Model: value for origen")
	ErrFooInvalidEventoMensMilEvent      = errors.New("MensMilEvent Model: Invalid value for evento")
	ErrFooInvalidMensajeMensMilEvent     = errors.New("MensMilEvent Model: Invalid value for mensaje")
	ErrFooInvalidTimeMensMilEvent        = errors.New("MensMilEvent Model: Invalid Time Value, out of +- 600 sec range")
)

type MensMilEvent struct {
	NroMM       int    `json:"nro_mm" binding:"required"`
	ClasifSeg   string `json:"clasificacion" binding:"required"`
	Precedencia string `json:"precedencia" binding:"required"`
	Cifrado     *bool  `json:"cifrado" binding:"required"`
	Destino     string `json:"destino" binding:"required"`
	Origen      string `json:"origen" binding:"required"`
	Evento      string `json:"evento" binding:"required"`
	Mensaje     string `json:"mensaje" binding:"required"`
	Time        int64  `json:"timestamp"`
}

func (e MensMilEvent) Write() (*MensMilEvent, error) {
	client := event_db.GetEventDB()
	stmt := `INSERT INTO mm_events (nro_mm, clasif_seg, precedencia, cifrado, destino, origen, evento, mensaje, gfh) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := (*client).Exec(stmt, e.NroMM, e.ClasifSeg, e.Precedencia, e.Cifrado, e.Destino, e.Origen, e.Evento, e.Mensaje, time.Unix(e.Time, 0))
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e MensMilEvent) Validate() error {
	if e.NroMM <= 0 {
		return ErrFooInvalidNroMMMensMilEvent
	}
	if !helpers.ValidString(e.ClasifSeg, ValidClasifSeg) {
		return ErrFooInvalidClasifSegMensMilEvent
	}
	if !helpers.ValidString(e.Precedencia, ValidPrecedencia) {
		return ErrFooInvalidPrecedenciaMensMilEvent
	}
	if strings.TrimSpace(e.Destino) == "" {
		return ErrFooInvalidDestinoMensMilEvent
	}
	if strings.TrimSpace(e.Origen) == "" {
		return ErrFooInvalidOrigenMensMilEvent
	}
	if strings.TrimSpace(e.Evento) == "" {
		return ErrFooInvalidEventoMensMilEvent
	}
	if strings.TrimSpace(e.Mensaje) == "" {
		return ErrFooInvalidMensajeMensMilEvent
	}
	diffTime := time.Now().Unix() - e.Time
	if diffTime > 600 || diffTime < -600 {
		return ErrFooInvalidTimeMensMilEvent
	}
	return nil
}

func (e MensMilEvent) Clear() (*MensMilEvent, error) {
	client := event_db.GetEventDB()
	stmt := `DELETE FROM mm_events`
	_, err := (*client).Exec(stmt)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
