// Weierstrass ECDH Curves
// 256-bit prime field Weierstrass curves y²=x³+ax+b.
package curves

import "crypto/elliptic"
import "math/big"

func strbig(s string) (i *big.Int) {
	i = new(big.Int)
	i.SetString(s, 0)
	return
}

var a256_1 = &elliptic.CurveParams{
	P:       strbig("0x00a6d2c41553ba68cfba46aa1281a6b23e95700c68c797367806f96b098d65af57"), // Prime
	N:       strbig("0x00a6d2c41553ba68cfba46aa1281a6b23ef4dc5fd9efc8b9f212b3be35c9fba89a"), // Order
	B:       strbig("0x003f6ebab31201e5576f5b0d18b1e7989ab274a734c0b3ee8a8d028d8c4601ebb5"), // B
	Gx:      strbig("0x4d6cacdb1d602c33aadb7218c8934888bb4f2b060a4465249462a7795bd1488b"),   // Generator X
	Gy:      strbig("0x9454bde43144d5078c9f60ebc07557220be7ad8d4af1be1d8bc73849dcf49881"),   // Generator Y
	BitSize: 256,
	Name:    "a256_1",
}

// a256v1() returns a Curve which implements a256v1
func A256v1() elliptic.Curve { return a256_1 }

var a256_2 = &elliptic.CurveParams{
	P:       strbig("0x0087d14dc3e41fedf2204981a026bf738ef2f4d8980a76ab396ec734c95aee2475"), // Prime
	N:       strbig("0x0087d14dc3e41fedf2204981a026bf7390143c205dc84df98e652bbb17486ad3ef"), // Order
	B:       strbig("0x006b686db550e8096e133540b5d3d41fd7a4181ab739264dbb001c2b6ed4677a69"), // B
	Gx:      strbig("0x0dc010e53beedbb939cbc7c8d92ae99b515694317bdfe862f2b82387d63b80af"),   // Generator X
	Gy:      strbig("0x4c21eac66867b583848866b178ec6d276bd76038bb5505410af9c05990e98ea4"),   // Generator Y
	BitSize: 256,
	Name:    "a256_2",
}

// a256v2() returns a Curve which implements a256v2
func A256v2() elliptic.Curve { return a256_2 }

var a256_3 = &elliptic.CurveParams{
	P:       strbig("0x00d4b31a2747dec22bc779637e32d2ee320aed34d54898050449a6e2128752e2f7"), // Prime
	N:       strbig("0x00d4b31a2747dec22bc779637e32d2ee309489dedbe9b854cf3b238ace2e4affa9"), // Order
	B:       strbig("0x00809c16abaa2a2e0b2c8b9abace15aeb586a558a4a0ace651877e76ba0f3067f1"), // B
	Gx:      strbig("0x6b775b74b0b48b59e0fa3aab3d09bda23e67f4446eb8a13db9e0a548dc08be4f"),   // Generator X
	Gy:      strbig("0xcde027bb7ae35bc84c7053d13427c776bb15d1808fd159908b2ec11d82a9cd48"),   // Generator Y
	BitSize: 256,
	Name:    "a256_3",
}

// a256v3() returns a Curve which implements a256v3
func A256v3() elliptic.Curve { return a256_3 }
