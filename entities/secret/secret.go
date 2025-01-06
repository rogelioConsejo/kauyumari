package secret

import "crypto/sha256"

type Secret string

func (s Secret) Hash() HashedSecret {
	h := sha256.New()
	h.Write([]byte(s))
	return HashedSecret(h.Sum(nil))
}

type HashedSecret string

func (hs HashedSecret) Compare(s Secret) bool {
	return hs == s.Hash()
}
