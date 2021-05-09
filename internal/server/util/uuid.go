package util

import (
	"encoding/hex"
)

type UUID [16]byte

// FromString returns UUID parsed from string input.
// Input is expected in a form accepted by UnmarshalText.
func FromString(input string) (u UUID, err error) {
	err = u.decodeHashLike([]byte(input))
	return
}

func (u *UUID) decodeHashLike(t []byte) (err error) {
	src := t[:]
	dst := u[:]

	if _, err = hex.Decode(dst, src); err != nil {
		return err
	}
	return
}
