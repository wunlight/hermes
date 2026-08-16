package password

type Hasher interface {
	Hash(password string) (string, error)
	Compare(password, encodedHash string) error
}
