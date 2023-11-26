package hasher

type Hasher interface {
	Hash(data string) string
}
