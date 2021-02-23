package models

import (
	"reflect"
	"testing"
	"time"
)

func TestNivelCombustible_Validate(t *testing.T) {
	type fields struct {
		Generador string
		Nivel     float32
		Time      int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Invalid Generator Name Test", fields{"", 10.0, time.Now().Unix()}, true},
		{"Invalid Level Test", fields{"ECR Cdo", -10.0, time.Now().Unix()}, true},
		{"Invalid Time Value Test", fields{"ECR Cdo", 10, time.Now().Unix() + 650}, true},
		{"Invalid Time Value Test", fields{"ECR Cdo", 10, time.Now().Unix() - 650}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NivelCombustible{
				Generador: tt.fields.Generador,
				Nivel:     tt.fields.Nivel,
				Time:      tt.fields.Time,
			}
			if err := e.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNivelCombustible_Write(t *testing.T) {
	type fields struct {
		Generador string
		Nivel     float32
		Time      int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    *NivelCombustible
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NivelCombustible{
				Generador: tt.fields.Generador,
				Nivel:     tt.fields.Nivel,
				Time:      tt.fields.Time,
			}
			got, err := e.Write()
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Write() got = %v, Want %v", got, tt.want)
			}
		})
	}
}
