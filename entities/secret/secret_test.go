package secret

import "testing"

func TestSecret_Hash(t *testing.T) {
	var s Secret = "some secret"
	var hs HashedSecret = s.Hash()
	if hs == "" {
		t.Errorf("HashedSecret should not be empty")
	}
	if !hs.Compare(s) {
		t.Errorf("HashedSecret.Compare should return true")
	}
}

func TestHashedSecret_Compare(t *testing.T) {
	var s Secret = "some secret"
	var hs HashedSecret = s.Hash()
	if !hs.Compare(s) {
		t.Fatal("HashedSecret.Compare should return true")
	}
	var anotherSecret Secret = "another secret"
	if hs.Compare(anotherSecret) {
		t.Fatal("HashedSecret.Compare should return false")
	}
}

type Key Secret

func (k Key) Hash() HashedKey {
	return HashedKey(Secret(k).Hash())
}

type HashedKey HashedSecret

func (h HashedKey) Compare(k Key) bool {
	return HashedSecret(h).Compare(Secret(k))
}

func TestSecretType(t *testing.T) {
	var k Key = "some key"
	var hk HashedKey = k.Hash()
	if hk == "" {
		t.Errorf("HashedKey should not be empty")
	}
	if !hk.Compare(k) {
		t.Errorf("HashedKey.Compare should return true")
	}
}
