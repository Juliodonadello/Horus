package models

import (
	"reflect"
	"testing"
	"time"
)

func TestEstadoServicio_Validate(t *testing.T) {
	type fields struct {
		Facilidad string
		Estado    string
		Time      int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Invalid Facility Name Test", fields{"ECR Cdo", "en servicio", time.Now().Unix()}, false},
		{"Invalid Facility Name Test", fields{"", "en servicio", time.Now().Unix()}, true},
		{"Invalid Facility Name Test", fields{"    ", "en servicio", time.Now().Unix()}, true},
		{"Invalid State Name Test", fields{"ECR Cdo", "sin servicio", time.Now().Unix()}, true},
		{"Invalid Time Value Test", fields{"ECR Cdo", "en servicio", time.Now().Unix() + 650}, true},
		{"Invalid Time Value Test", fields{"ECR Cdo", "en servicio", time.Now().Unix() - 650}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EstadoServicio{
				Facilidad: tt.fields.Facilidad,
				Estado:    tt.fields.Estado,
				Time:      tt.fields.Time,
			}
			if err := e.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEstadoServicio_Write(t *testing.T) {
	type fields struct {
		Facilidad string
		Estado    string
		Time      int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    *EstadoServicio
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EstadoServicio{
				Facilidad: tt.fields.Facilidad,
				Estado:    tt.fields.Estado,
				Time:      tt.fields.Time,
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
