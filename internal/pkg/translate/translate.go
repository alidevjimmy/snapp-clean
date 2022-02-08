package translate

type Translate interface {
	Trans(s string, lang ...Language) string
}

type Language uint

const (
	_ Language = iota
	Fa
	En
)