package seleccion

type Doparti struct {
	Gana bool
	DiferenciaGoles int
}

func Clasificamos(argentina, islandia Doparti) bool {
	if !argentina.Gana {
		return false
	} else if !islandia.Gana {
		return true
	} else {
		return argentina.DiferenciaGoles > islandia.DiferenciaGoles
	}
}