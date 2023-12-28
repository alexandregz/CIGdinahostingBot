package comandos

import (
	"testing"
)

var isComandoEParametroDeAdmin = []struct {
	comando   string
	parametro string
	expected  bool
}{
	{"msg", "", true},
	{"igualdade", "actas", true},
	{"fake2", "", false},
	{"contasanuais", "2022", true},
	{"temas", "privados", true},
	{"temas", "listado", false},
}

var isComandoDeAdmin = []struct {
	comando  string
	expected bool
}{
	{"msg", true},
	{"igualdade", true},
	{"fake1", false},
	{"fake2", false},
	{"contasanuais", true},
}

func TestComandoDeAdmin(t *testing.T) {
	for _, tt := range isComandoDeAdmin {
		t.Run(tt.comando, func(t *testing.T) {
			got := ComandoDeAdmin(tt.comando)
			if got != tt.expected {
				t.Errorf("Agardado: %v, got: %v", tt.expected, got)
			}
		})
	}
}

func TestComandoEParamDeAdmin(t *testing.T) {
	for _, tt := range isComandoEParametroDeAdmin {
		t.Run(tt.comando, func(t *testing.T) {
			got := ComandoEParamDeAdmin(tt.comando, tt.parametro)
			if got != tt.expected {
				t.Errorf("Agardado (%v %v): %v, got: %v", tt.comando, tt.parametro, tt.expected, got)
			}
		})
	}
}
