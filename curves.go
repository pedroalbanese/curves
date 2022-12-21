package curves

import "crypto/elliptic"
import "math/big"

func strbig(s string) (i *big.Int) {
	i = new(big.Int)
	i.SetString(s,0)
	return
}

var a128 = &elliptic.CurveParams{
	P: strbig("0x008b996f0533f8417866c799fc61310415"), // Prime
	N: strbig("0x008b996f0533f84179acba3b5cb109661e"), // Order
	B: strbig("0x003ef5802fd915198e7a7173f5acce0545"), // B
	Gx: strbig("0x1e95f559b16d399a7300627c6267deb6"),  // Generator X
	Gy: strbig("0x0447d6ff6002fda175090cb0fe32ccbe"),  // Generator Y
	BitSize: 128,
	Name: "a128",
}

// a128v1() returns a Curve which implements a128v1
func A128v1() elliptic.Curve { return a128 }

var a256 = &elliptic.CurveParams{
	P: strbig("0x00a6d2c41553ba68cfba46aa1281a6b23e95700c68c797367806f96b098d65af57"), // Prime
	N: strbig("0x00a6d2c41553ba68cfba46aa1281a6b23ef4dc5fd9efc8b9f212b3be35c9fba89a"), // Order
	B: strbig("0x003f6ebab31201e5576f5b0d18b1e7989ab274a734c0b3ee8a8d028d8c4601ebb5"), // B
	Gx: strbig("0x4d6cacdb1d602c33aadb7218c8934888bb4f2b060a4465249462a7795bd1488b"),  // Generator X
	Gy: strbig("0x9454bde43144d5078c9f60ebc07557220be7ad8d4af1be1d8bc73849dcf49881"),  // Generator Y
	BitSize: 256,
	Name: "a256",
}

// a256v1() returns a Curve which implements a256v1
func A256v1() elliptic.Curve { return a256 }
