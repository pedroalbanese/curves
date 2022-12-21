package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"

	"github.com/pedroalbanese/curves"
	"github.com/pedroalbanese/eccrypt"
)

var (
	derive = flag.Bool("derive", false, "Derive shared secret key.")
	keygen = flag.Bool("keygen", false, "Generate keypair.")
	key    = flag.String("key", "", "Private/Public key depending on operation.")
	public = flag.String("pub", "", "Remote's side Public key. (for DERIVE)")
	enc    = flag.Bool("enc", false, "Encrypt with Publickey.")
	dec    = flag.Bool("dec", false, "Decrypt with Privatekey.")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage of", os.Args[0]+":")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var privatekey *ecdsa.PrivateKey
	var pubkey ecdsa.PublicKey
	var err error
	var pubkeyCurve elliptic.Curve

	pubkeyCurve = curves.A256v1()

	if *keygen {
		if *key != "" {
			privatekey, err = ReadPrivateKeyFromHex(*key)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			privatekey = new(ecdsa.PrivateKey)
			privatekey, err = ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for len(WritePrivateKeyToHex(privatekey)) != 64 {
				privatekey, err = ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				break
			}
			if len(WritePrivateKeyToHex(privatekey)) != 64 {
				log.Fatal("Private key too short!")
				os.Exit(1)
			}
		}
		pubkey = privatekey.PublicKey
		fmt.Println("Private= " + WritePrivateKeyToHex(privatekey))
		fmt.Println("Public= " + WritePublicKeyToHex(&pubkey))
		os.Exit(0)
	}

	if *derive {
		private, err := ReadPrivateKeyFromHex(*key)
		if err != nil {
			log.Fatal(err)
		}
		public, err := ReadPublicKeyFromHex(*public)
		if err != nil {
			log.Fatal(err)
		}

		b, _ := public.Curve.ScalarMult(public.X, public.Y, private.D.Bytes())
		shared := b.Bytes()
		fmt.Printf("Shared= %x\n", shared)
		os.Exit(0)
	}

	if *enc {
		public, err := ReadPublicKeyFromHexX(*key)
		if err != nil {
			log.Fatal(err)
		}
		buf := bytes.NewBuffer(nil)
		data := os.Stdin
		io.Copy(buf, data)
		scanner := string(buf.Bytes())
		ciphertxt, err := eccrypt.EncryptAsn1(public, []byte(scanner), rand.Reader)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", ciphertxt)
		os.Exit(0)
	}

	if *dec {
		private, err := ReadPrivateKeyFromHexX(*key)
		if err != nil {
			log.Fatal(err)
		}
		buf := bytes.NewBuffer(nil)
		data := os.Stdin
		io.Copy(buf, data)
		scanner := string(buf.Bytes())
		str := string(scanner)
		plaintxt, err := eccrypt.DecryptAsn1(private, []byte(str))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", plaintxt)
		os.Exit(0)
	}
}

func ReadPrivateKeyFromHex(Dhex string) (*ecdsa.PrivateKey, error) {
	c := curves.A256v1()
	d, err := hex.DecodeString(Dhex)
	if err != nil {
		return nil, err
	}
	k := new(big.Int).SetBytes(d)
	params := c.Params()
	one := new(big.Int).SetInt64(1)
	n := new(big.Int).Sub(params.N, one)
	if k.Cmp(n) >= 0 {
		return nil, errors.New("privateKey's D is overflow.")
	}
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	return priv, nil
}

func ReadPublicKeyFromHex(Qhex string) (*ecdsa.PublicKey, error) {
	q, err := hex.DecodeString(Qhex)
	if err != nil {
		return nil, err
	}
	if len(q) == 65 && q[0] == byte(0x04) {
		q = q[1:]
	}
	if len(q) != 64 {
		return nil, errors.New("publicKey is not uncompressed.")
	}
	pub := new(ecdsa.PublicKey)
	pub.Curve = curves.A256v1()
	pub.X = new(big.Int).SetBytes(q[:32])
	pub.Y = new(big.Int).SetBytes(q[32:])
	return pub, nil
}

func ReadPrivateKeyFromHexX(Dhex string) (*eccrypt.PrivateKey, error) {
	c := curves.A256v1()
	d, err := hex.DecodeString(Dhex)
	if err != nil {
		return nil, err
	}
	k := new(big.Int).SetBytes(d)
	params := c.Params()
	one := new(big.Int).SetInt64(1)
	n := new(big.Int).Sub(params.N, one)
	if k.Cmp(n) >= 0 {
		return nil, errors.New("privateKey's D is overflow.")
	}
	priv := new(eccrypt.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	return priv, nil
}

func ReadPublicKeyFromHexX(Qhex string) (*eccrypt.PublicKey, error) {
	q, err := hex.DecodeString(Qhex)
	if err != nil {
		return nil, err
	}
	if len(q) == 65 && q[0] == byte(0x04) {
		q = q[1:]
	}
	if len(q) != 64 {
		return nil, errors.New("publicKey is not uncompressed.")
	}
	pub := new(eccrypt.PublicKey)
	pub.Curve = curves.A256v1()
	pub.X = new(big.Int).SetBytes(q[:32])
	pub.Y = new(big.Int).SetBytes(q[32:])
	return pub, nil
}

func WritePrivateKeyToHex(key *ecdsa.PrivateKey) string {
	d := key.D.Bytes()
	if n := len(d); n < 32 {
		d = append(zeroByteSlice()[:64-n], d...)
	}
	c := []byte{}
	c = append(c, d...)
	return hex.EncodeToString(c)
}

func WritePublicKeyToHex(key *ecdsa.PublicKey) string {
	x := key.X.Bytes()
	y := key.Y.Bytes()
	if n := len(x); n < 32 {
		x = append(zeroByteSlice()[:32-n], x...)
	}
	if n := len(y); n < 32 {
		y = append(zeroByteSlice()[:32-n], y...)
	}
	c := []byte{}
	c = append(c, x...)
	c = append(c, y...)
	return hex.EncodeToString(c)
}

func zeroByteSlice() []byte {
	return []byte{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}
}
