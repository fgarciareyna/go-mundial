package seleccion

import "testing"

func TestClasificamos(t *testing.T) {
	cases := []struct {
		name                     string
		argentinaGana            bool
		argentinaDiferenciaGoles int
		islandiaGana             bool
		islandiaDiferenciaGoles  int
		clasificamos             bool
	}{
		{
			name:                     "Test no ganamos",
			argentinaGana:            false,
			argentinaDiferenciaGoles: 0,
			islandiaGana:             false,
			islandiaDiferenciaGoles:  0,
			clasificamos:             false,
		},
		{
			name:                     "Test Islandia no gana",
			argentinaGana:            true,
			argentinaDiferenciaGoles: 1,
			islandiaGana:             false,
			islandiaDiferenciaGoles:  0,
			clasificamos:             true,
		},
		{
			name:                     "Test nos faltan goles",
			argentinaGana:            true,
			argentinaDiferenciaGoles: 1,
			islandiaGana:             true,
			islandiaDiferenciaGoles:  1,
			clasificamos:             false,
		},
		{
			name:                     "Test no nos faltan goles",
			argentinaGana:            true,
			argentinaDiferenciaGoles: 2,
			islandiaGana:             true,
			islandiaDiferenciaGoles:  1,
			clasificamos:             true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			argentina := Doparti{
				Gana:            c.argentinaGana,
				DiferenciaGoles: c.argentinaDiferenciaGoles,
			}
			islandia := Doparti{
				Gana:            c.islandiaGana,
				DiferenciaGoles: c.islandiaDiferenciaGoles,
			}
			clasificamos := Clasificamos(argentina, islandia)
			if clasificamos != c.clasificamos {
				t.Fail()
			}
		})
	}
}
