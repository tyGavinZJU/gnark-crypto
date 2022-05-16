// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by consensys/gnark-crypto DO NOT EDIT

package bls24315

//Note: This only works for simple extensions

import (
	"github.com/consensys/gnark-crypto/ecc/bls24-315/fp"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
)

func g1IsogenyXNumerator(dst *fp.Element, x *fp.Element) {
	g1EvalPolynomial(dst,
		false,
		[]fp.Element{
			{11620002718874663739, 4984467296741409765, 9174718300976205935, 11374294140644765434, 331965326722599209},
			{12794915441326992831, 3515443655574390653, 6174257928039766159, 70148989344615692, 200953992158149919},
			{5852384876649512947, 11848499933279379168, 12693517207910261404, 4355336966086013201, 153982054162701797},
		},
		x)
}

func g1IsogenyXDenominator(dst *fp.Element, x *fp.Element) {
	g1EvalPolynomial(dst,
		true,
		[]fp.Element{
			{16605520835351066362, 4532778258980819953, 11041097066391022716, 6626569051763865297, 118015358745724890},
		},
		x)
}

func g1IsogenyYNumerator(dst *fp.Element, x *fp.Element, y *fp.Element) {
	var _dst fp.Element
	g1EvalPolynomial(&_dst,
		false,
		[]fp.Element{
			{9734843649657667679, 9905469488516037607, 12244225131002460472, 12160927269755757379, 293726840634836990},
			{332611309977308143, 8673449249147179720, 7968180610051701274, 525286427825436485, 27337445552095458},
			{5937151911073875102, 12114288429387176123, 10459089249045026662, 1691716757613274170, 129980835765506182},
			{6958041652386594810, 8306499057468875249, 14372428283824529086, 591175209446655968, 248441179553069595},
		},
		x)

	dst.Mul(&_dst, y)
}

func g1IsogenyYDenominator(dst *fp.Element, x *fp.Element) {
	g1EvalPolynomial(dst,
		true,
		[]fp.Element{
			{7466136663908942255, 5910124112997814042, 598236406339551119, 15948603688162126360, 216078840945103380},
			{16886107642822413408, 14927238232380936652, 17792216571653695247, 7051181952824703829, 174959651533410931},
			{4859375930510419181, 8833836595284088531, 17071951839434271380, 4605949628774745540, 11145771293737278},
		},
		x)
}

func g1Isogeny(p *G1Affine) {

	den := make([]fp.Element, 2)

	g1IsogenyYDenominator(&den[1], &p.X)
	g1IsogenyXDenominator(&den[0], &p.X)

	g1IsogenyYNumerator(&p.Y, &p.X, &p.Y)
	g1IsogenyXNumerator(&p.X, &p.X)

	den = fp.BatchInvert(den)

	p.X.Mul(&p.X, &den[0])
	p.Y.Mul(&p.Y, &den[1])
}

// g1SqrtRatio computes the square root of u/v and returns 0 iff u/v was indeed a quadratic residue
// if not, we get sqrt(Z * u / v). Recall that Z is non-residue
// The main idea is that since the computation of the square root involves taking large powers of u/v, the inversion of v can be avoided
func g1SqrtRatio(z *fp.Element, u *fp.Element, v *fp.Element) uint64 {

	// Taken from https://datatracker.ietf.org/doc/draft-irtf-cfrg-hash-to-curve/13/ F.2.1.1. for any field

	tv1 := fp.Element{11195128742969911322, 1359304652430195240, 15267589139354181340, 10518360976114966361, 300769513466036652} //tv1 = c6

	var tv2, tv3, tv4, tv5 fp.Element
	var exp big.Int
	// c4 = 1048575 = 2^20 - 1
	// q is odd so c1 is at least 1.
	exp.SetBytes([]byte{15, 255, 255})

	tv2.Exp(*v, &exp)
	tv3.Mul(&tv2, &tv2)
	tv3.Mul(&tv3, v)

	// line 5
	tv5.Mul(u, &tv3)

	// c3 = 18932887415653914611351818986134037849871398170907377879650252106493894621432467626129921
	exp.SetBytes([]byte{38, 17, 208, 21, 172, 54, 178, 134, 159, 186, 76, 95, 75, 226, 245, 126, 246, 14, 128, 213, 19, 208, 215, 2, 16, 247, 46, 210, 149, 239, 40, 19, 127, 64, 23, 250, 1})
	tv5.Exp(tv5, &exp)
	tv5.Mul(&tv5, &tv2)
	tv2.Mul(&tv5, v)
	tv3.Mul(&tv5, u)

	// line 10
	tv4.Mul(&tv3, &tv2)

	// c5 = 524288
	exp.SetBytes([]byte{8, 0, 0})
	tv5.Exp(tv4, &exp)

	isQNr := g1NotOne(&tv5)

	tv2.Mul(&tv3, &fp.Element{1141794007209116247, 256324699145650176, 2958838397954514392, 9976887947641032208, 153331829745922234})
	tv5.Mul(&tv4, &tv1)

	// line 15

	tv3.Select(int(isQNr), &tv3, &tv2)
	tv4.Select(int(isQNr), &tv4, &tv5)

	exp.Lsh(big.NewInt(1), 20-2)

	for i := 20; i >= 2; i-- {
		//line 20
		tv5.Exp(tv4, &exp)
		nE1 := g1NotOne(&tv5)

		tv2.Mul(&tv3, &tv1)
		tv1.Mul(&tv1, &tv1)
		tv5.Mul(&tv4, &tv1)

		tv3.Select(int(nE1), &tv3, &tv2)
		tv4.Select(int(nE1), &tv4, &tv5)

		exp.Rsh(&exp, 1)
	}

	*z = tv3
	return isQNr
}

func g1NotOne(x *fp.Element) uint64 {

	var one fp.Element
	return one.SetOne().NotEqual(x)

}

/*
// g1SetZ sets z to [13].
func g1SetZ(z *fp.Element) {
    z.Set( &fp.Element  { 8178485296672800069, 8476448362227282520, 14180928431697993131, 4308307642551989706, 120359802761433421 } )
}*/

// g1MulByZ multiplies x by [13] and stores the result in z
func g1MulByZ(z *fp.Element, x *fp.Element) {

	res := *x

	res.Double(&res)
	res.Add(&res, x)
	res.Double(&res)
	res.Double(&res)
	res.Add(&res, x)

	*z = res
}

// From https://datatracker.ietf.org/doc/draft-irtf-cfrg-hash-to-curve/13/ Pg 80
// sswuMapG1 implements the SSWU map
// No cofactor clearing
func sswuMapG1(u *fp.Element) G1Affine {

	var tv1 fp.Element
	tv1.Square(u)

	//mul tv1 by Z
	g1MulByZ(&tv1, &tv1)

	var tv2 fp.Element
	tv2.Square(&tv1)
	tv2.Add(&tv2, &tv1)

	var tv3 fp.Element
	//Standard doc line 5
	var tv4 fp.Element
	tv4.SetOne()
	tv3.Add(&tv2, &tv4)
	tv3.Mul(&tv3, &fp.Element{16058189711238232929, 8302337653269510588, 11411933349841587630, 8954038365926617417, 177308873523699836})

	tv2NZero := g1NotZero(&tv2)

	// tv4 = Z
	tv4 = fp.Element{8178485296672800069, 8476448362227282520, 14180928431697993131, 4308307642551989706, 120359802761433421}

	tv2.Neg(&tv2)
	tv4.Select(int(tv2NZero), &tv4, &tv2)
	tv2 = fp.Element{5402807948305211529, 9163880483319140034, 7646126700453841420, 11071466103913358468, 124200740526673728}
	tv4.Mul(&tv4, &tv2)

	tv2.Square(&tv3)

	var tv6 fp.Element
	//Standard doc line 10
	tv6.Square(&tv4)

	var tv5 fp.Element
	tv5.Mul(&tv6, &fp.Element{5402807948305211529, 9163880483319140034, 7646126700453841420, 11071466103913358468, 124200740526673728})

	tv2.Add(&tv2, &tv5)
	tv2.Mul(&tv2, &tv3)
	tv6.Mul(&tv6, &tv4)

	//Standards doc line 15
	tv5.Mul(&tv6, &fp.Element{16058189711238232929, 8302337653269510588, 11411933349841587630, 8954038365926617417, 177308873523699836})
	tv2.Add(&tv2, &tv5)

	var x fp.Element
	x.Mul(&tv1, &tv3)

	var y1 fp.Element
	gx1NSquare := g1SqrtRatio(&y1, &tv2, &tv6)

	var y fp.Element
	y.Mul(&tv1, u)

	//Standards doc line 20
	y.Mul(&y, &y1)

	x.Select(int(gx1NSquare), &tv3, &x)
	y.Select(int(gx1NSquare), &y1, &y)

	y1.Neg(&y)
	y.Select(int(g1Sgn0(u)^g1Sgn0(&y)), &y, &y1)

	//Standards doc line 25
	x.Div(&x, &tv4)

	return G1Affine{x, y}
}

// mapToG1 invokes the SSWU map, and guarantees that the result is in g1
func mapToG1(u fp.Element) G1Affine {
	res := sswuMapG1(&u)
	res.ClearCofactor(&res)
	return res
}

// EncodeToG1 maps a fp.Element to a point on the curve using the Simplified Shallue and van de Woestijne Ulas map
//https://datatracker.ietf.org/doc/draft-irtf-cfrg-hash-to-curve/13/#section-6.6.3
func EncodeToG1(msg, dst []byte) (G1Affine, error) {

	var res G1Affine
	u, err := hashToFp(msg, dst, 1)
	if err != nil {
		return res, err
	}

	res = sswuMapG1(&u[0])

	//this is in an isogenous curve
	g1Isogeny(&res)

	res.ClearCofactor(&res)

	return res, nil
}

// HashToG1 hashes a byte string to the G1 curve. Usable as a random oracle.
// https://tools.ietf.org/html/draft-irtf-cfrg-hash-to-curve-06#section-3
func HashToG1(msg, dst []byte) (G1Affine, error) {
	u, err := hashToFp(msg, dst, 2*1)
	if err != nil {
		return G1Affine{}, err
	}

	Q0 := sswuMapG1(&u[0])
	Q1 := sswuMapG1(&u[1])

	//TODO: Add in E' first, then apply isogeny
	g1Isogeny(&Q0)
	g1Isogeny(&Q1)

	var _Q0, _Q1 G1Jac
	_Q0.FromAffine(&Q0)
	_Q1.FromAffine(&Q1).AddAssign(&_Q0)

	_Q1.ClearCofactor(&_Q1)

	Q1.FromJacobian(&_Q1)
	return Q1, nil
}

// g1Sgn0 is an algebraic substitute for the notion of sign in ordered fields
// Namely, every non-zero quadratic residue in a finite field of characteristic =/= 2 has exactly two square roots, one of each sign
// Taken from https://datatracker.ietf.org/doc/draft-irtf-cfrg-hash-to-curve/ section 4.1
// The sign of an element is not obviously related to that of its Montgomery form
func g1Sgn0(z *fp.Element) uint64 {

	nonMont := *z
	nonMont.FromMont()

	return nonMont[0] % 2

}

func g1EvalPolynomial(z *fp.Element, monic bool, coefficients []fp.Element, x *fp.Element) {
	dst := coefficients[len(coefficients)-1]

	if monic {
		dst.Add(&dst, x)
	}

	for i := len(coefficients) - 2; i >= 0; i-- {
		dst.Mul(&dst, x)
		dst.Add(&dst, &coefficients[i])
	}

	z.Set(&dst)
}

func g1NotZero(x *fp.Element) uint64 {

	return x[0] | x[1] | x[2] | x[3] | x[4]

}

// hashToFp hashes msg to count prime field elements.
// https://tools.ietf.org/html/draft-irtf-cfrg-hash-to-curve-06#section-5.2
func hashToFp(msg, dst []byte, count int) ([]fp.Element, error) {

	// 128 bits of security
	// L = ceil((ceil(log2(p)) + k) / 8), where k is the security parameter = 128
	L := 64

	lenInBytes := count * L
	pseudoRandomBytes, err := ecc.ExpandMsgXmd(msg, dst, lenInBytes)
	if err != nil {
		return nil, err
	}

	res := make([]fp.Element, count)
	for i := 0; i < count; i++ {
		res[i].SetBytes(pseudoRandomBytes[i*L : (i+1)*L])
	}
	return res, nil
}
