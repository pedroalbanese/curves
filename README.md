# I&I Elliptic curves
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/curves/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/curves?status.png)](http://godoc.org/github.com/pedroalbanese/curves)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/curves)](https://goreportcard.com/report/github.com/pedroalbanese/curves)

### 128/256-bit prime field Weierstrass curves y²=x³+ax+b.

Elliptic-curve cryptography (ECC) is an approach to public-key cryptography based on the algebraic structure of elliptic curves over finite fields. ECC allows smaller keys compared to non-EC cryptography (based on plain Galois fields) to provide equivalent security.

Public-key cryptography, or asymmetric cryptography, is the field of cryptographic systems that use pairs of related keys. Each key pair consists of a public key and a corresponding private key. Key pairs are generated with cryptographic algorithms based on mathematical problems termed one-way functions. Security of public-key cryptography depends on keeping the private key secret; the public key can be openly distributed without compromising security.

In a public-key encryption system, anyone with a public key can encrypt a message, yielding a ciphertext, but only those who know the corresponding private key can decrypt the ciphertext to obtain the original message.

### Usage
```
Usage of a256v1:
  -decrypt
        Decrypt with Privatekey.
  -derive
        Derive shared secret key.
  -encrypt
        Encrypt with Publickey.
  -key string
        Private/Public key depending on operation.
  -keygen
        Generate keypair.
  -peerkey string
        Remote's side Public key. (for DERIVE)
```

### Usage
#### EC Diffie-Hellman:
```sh
./edgetk -derive -key $private -public $peerkey
```
#### Asymmetric Encryption/Decryption:
```sh
./edgetk -encrypt -key $public < plaintext.ext > ciphertext.ext
./edgetk -decrypt -key $private < ciphertext.ext > plaintext.ext
```
### TODO
- [ ] 512-bit

## License
This project is licensed under the ISC License.
##### Written for educational purposes. Copyright (c) 2020-2022, ALBANESE Research Lab.

