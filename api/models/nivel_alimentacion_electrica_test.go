package models

import (
	"reflect"
	"testing"
)

func TestTensionGenerador_Validate(t *testing.T) {
	type fields struct {
		Generador string
		Tension   float32
		Corriente float32
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Invalid Generator Name Test", fields{"", 10.0, 2.0}, true},
		{"Invalid Tension Test", fields{"generador 1", -10.0, 2.0}, true},
		{"Invalid Tension Test", fields{"generador 1", 700.0, 2.0}, true},
		{"Invalid Corriente Test", fields{"generador 1", 220.0, -2.0}, true},
		{"Invalid Corriente Test", fields{"generador 1", 220.0, 52.0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Alimentacion{
				Generador: tt.fields.Generador,
				Tension:   tt.fields.Tension,
				Corriente: tt.fields.Corriente,
			}
			if err := e.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTensionGenerador_Write(t *testing.T) {
	type fields struct {
		Generador string
		Tension   float32
		Corriente float32
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Alimentacion
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Alimentacion{
				Generador: tt.fields.Generador,
				Tension:   tt.fields.Tension,
				Corriente: tt.fields.Corriente,
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
