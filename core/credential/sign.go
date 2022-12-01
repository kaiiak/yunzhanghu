package credential

type Signer interface {
	Type() string
	Sign(params string) (string, error)
}
