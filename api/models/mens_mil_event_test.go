package models

import (
	"reflect"
	"testing"
	"time"
)

func TestMensMilEvent_Validate(t *testing.T) {
	type fields struct {
		ClasifSeg   string
		Precedencia string
		Cifrado     bool
		Destino     string
		Origen      string
		Evento      string
		Time        int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Valid Test", fields{"secreto", "rutina", true, "ECR Cdo", "CMD", "sale", time.Now().Unix()}, false},
		{"Invalid Clasif Seg Test", fields{"rutinario", "rutina", true, "ECR Cdo", "CMD", "sale", time.Now().Unix()}, true},
		{"Invalid Precedencia Test", fields{"publico", "secreto", true, "ECR Cdo", "CMD", "sale", time.Now().Unix()}, true},
		{"Invalid Origen Test", fields{"rutinario", "rutina", false, "ECR Cdo", "", "sale", time.Now().Unix()}, true},
		{"Invalid Destino Test", fields{"rutinario", "rutina", false, "  ", "CMD", "sale", time.Now().Unix()}, true},
		{"Invalid Evento Test", fields{"secreto", "rutina", true, "ECR Cdo", "CMD", "", time.Now().Unix()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := MensMilEvent{
				ClasifSeg:   tt.fields.ClasifSeg,
				Precedencia: tt.fields.Precedencia,
				Cifrado:     tt.fields.Cifrado,
				Destino:     tt.fields.Destino,
				Origen:      tt.fields.Origen,
				Evento:      tt.fields.Evento,
				Time:        tt.fields.Time,
			}
			if err := e.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMensMilEvent_Write(t *testing.T) {
	type fields struct {
		ClasifSeg   string
		Precedencia string
		Cifrado     bool
		Destino     string
		Origen      string
		Evento      string
		Time        int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    *MensMilEvent
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := MensMilEvent{
				ClasifSeg:   tt.fields.ClasifSeg,
				Precedencia: tt.fields.Precedencia,
				Cifrado:     tt.fields.Cifrado,
				Destino:     tt.fields.Destino,
				Origen:      tt.fields.Origen,
				Evento:      tt.fields.Evento,
				Time:        tt.fields.Time,
			}
			got, err := e.Write()
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Write() got = %v, want %v", got, tt.want)
			}
		})
	}
}
