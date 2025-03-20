package argon2id

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

type ArgonParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var DefaultParams = ArgonParams{
	Memory:      64 * 1024,
	Iterations:  1,
	Parallelism: 4,
	SaltLength:  16,
	KeyLength:   32,
}

func EncodeHash(password string, p ArgonParams) (string, error) {
	salt := make([]byte, p.SaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return EncodeHashWithSalt(password, p, salt)
}

func EncodeHashWithSalt(password string, p ArgonParams, salt []byte) (string, error) {
	b64Hash := base64.RawStdEncoding.EncodeToString(argon2.IDKey(
		[]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength))
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func DecodeHash(encodedHash string) (p *ArgonParams, salt []byte, hash []byte, err error) {
	argonVals := strings.Split(encodedHash, "$")
	if len(argonVals) != 6 {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	var version int
	_, err = fmt.Sscanf(argonVals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &ArgonParams{}
	_, err = fmt.Sscanf(argonVals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(argonVals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(argonVals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))
	return p, salt, hash, nil
}
