/*
Copyright Gen Digital Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package localsuite

import (
	kmsapi "github.com/trustbloc/kms-go/spi/kms"
)

type signer interface {
	Sign(msg []byte, kh any) ([]byte, error)
}

type multiSigner interface {
	signer
	SignMulti(messages [][]byte, kh any) ([]byte, error)
}

type verifier interface {
	Verify(signature []byte, msg []byte, kh any) error
}

type signerVerifier interface {
	signer
	verifier
}

type keyGetter interface {
	Get(keyID string) (any, error)
}

type keyHandleFetcher interface {
	PubKeyBytesToHandle(pubKeyBytes []byte, keyType kmsapi.KeyType, opts ...kmsapi.KeyOpts) (any, error)
	ExportPubKeyBytes(keyID string) ([]byte, kmsapi.KeyType, error)
	keyGetter
}

type keyCreator interface {
	CreateAndExportPubKeyBytes(kt kmsapi.KeyType, opts ...kmsapi.KeyOpts) (string, []byte, error)
	ExportPubKeyBytes(id string) ([]byte, kmsapi.KeyType, error)
}

type keyManager interface {
	keyCreator
	keyHandleFetcher
}

type encDecrypter interface {
	Encrypt(msg, aad []byte, kh any) ([]byte, []byte, error)
	Decrypt(cipher, aad, nonce []byte, kh any) ([]byte, error)
}

type allCrypto interface {
	multiSigner
	verifier
	encDecrypter
}
