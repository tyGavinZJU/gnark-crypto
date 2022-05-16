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

package bw6633

//Note: This only works for simple extensions

import (
	"github.com/consensys/gnark-crypto/ecc/bw6-633/fp"
	"math/big"
)

func g2IsogenyXNumerator(dst *fp.Element, x *fp.Element) {
	g2EvalPolynomial(dst,
		false,
		[]fp.Element{
			{12267713055837521825, 1212596092600778933, 5226395968079745974, 11928252065127869839, 16368632504593357993, 16063800155037832323, 9358718386866238663, 14845980894001477527, 15811357545249775034, 40557942204006272},
			{597834311555652830, 5676150712176383509, 8459519236066800431, 11690428270517348528, 11839864809966557220, 1185830464157066542, 5950198841798077595, 13670804634510857615, 7801381657215673717, 65313904097694188},
			{13784414184985853237, 15655785158186492397, 12101352116729629183, 16924728184753209936, 5262388757860076597, 3980668092298208488, 3082578409602424779, 12803009346779985710, 15448751926107558648, 54234452766130564},
		},
		x)
}

func g2IsogenyXDenominator(dst *fp.Element, x *fp.Element) {
	g2EvalPolynomial(dst,
		true,
		[]fp.Element{
			{11192702706658734929, 9471950201046594205, 5654150871742973517, 1064926869499066026, 7399057011292302262, 13680728223779956488, 6863773321185403869, 18152681008199425511, 13601441751122646453, 12667349172889989},
		},
		x)
}

func g2IsogenyYNumerator(dst *fp.Element, x *fp.Element, y *fp.Element) {
	var _dst fp.Element
	g2EvalPolynomial(&_dst,
		false,
		[]fp.Element{
			{4966347166805171785, 5901766004022529993, 8116344375614693226, 15488373205894574973, 13683749641818622675, 11193797679007774234, 17818803555973174377, 10275933887373468852, 9920810925744653786, 28792607450625124},
			{1823522777607236865, 12808208129182298389, 10249345152917524976, 2251134980320265253, 6728735963367895750, 10712028649958228879, 3518459129547408211, 8820432640482636080, 15390508760019465661, 22899278098927699},
			{13053911991000296256, 13326713889233037794, 15592466212398492321, 8452295026155763968, 4439663138758526957, 1683423806334571317, 12502996284216256697, 17552733145339727247, 11982958884480011642, 15108100407245694},
			{5425312849086906017, 10033334687036402837, 4599082379215668754, 13004037782186734380, 9291261417025692735, 6649681009527006396, 13587498249323805949, 3266555558175884538, 7583932763725528791, 68548604252713076},
		},
		x)

	dst.Mul(&_dst, y)
}

func g2IsogenyYDenominator(dst *fp.Element, x *fp.Element) {
	g2EvalPolynomial(dst,
		true,
		[]fp.Element{
			{18150455001590453128, 16667992893458764333, 11433476296464492694, 15623787756943135869, 5167287249901804111, 11533969929056753328, 11759209128489608181, 4071561834127893664, 874055937145507461, 36923056072234183},
			{1607212096617940568, 15885536245478043276, 3832630086595040712, 10941954510883014474, 10683649199422707633, 1996189670536633415, 6106362375082575102, 7453025031903452354, 3025605118685313536, 13974970650264338},
			{15131364046266653171, 9969106529430231000, 16962452615228920552, 3194780608497198078, 3750426960167355170, 4148696523920766233, 2144575889846659993, 17564554877179173302, 3910837105948836129, 38002047518669969},
		},
		x)
}

func g2Isogeny(p *G2Affine) {

	den := make([]fp.Element, 2)

	g2IsogenyYDenominator(&den[1], &p.X)
	g2IsogenyXDenominator(&den[0], &p.X)

	g2IsogenyYNumerator(&p.Y, &p.X, &p.Y)
	g2IsogenyXNumerator(&p.X, &p.X)

	den = fp.BatchInvert(den)

	p.X.Mul(&p.X, &den[0])
	p.Y.Mul(&p.Y, &den[1])
}

// g2SqrtRatio computes the square root of u/v and returns 0 iff u/v was indeed a quadratic residue
// if not, we get sqrt(Z * u / v). Recall that Z is non-residue
// The main idea is that since the computation of the square root involves taking large powers of u/v, the inversion of v can be avoided
func g2SqrtRatio(z *fp.Element, u *fp.Element, v *fp.Element) uint64 {

	// Taken from https://datatracker.ietf.org/doc/draft-irtf-cfrg-hash-to-curve/13/ F.2.1.3. q = 5 mod 8
	// TODO: Test correct use of Element.Select

	var tv1, tv2 fp.Element
	tv1.Square(v)
	tv2.Mul(&tv1, v)
	tv1.Square(&tv1)
	tv2.Mul(&tv2, u)
	tv1.Mul(&tv1, &tv2)

	var c1 big.Int
	// c1 = 2561809830520971834851673423317370187208698865113597259441094318876502093964723972342941881295000173427447071280550850682906970629308555147846507041252715387332745928667496467633282424373249
	c1.SetBytes([]byte{36, 204, 103, 152, 30, 107, 236, 127, 131, 66, 233, 224, 58, 229, 86, 181, 31, 155, 24, 235, 175, 58, 88, 233, 203, 46, 211, 91, 55, 123, 69, 240, 42, 84, 216, 31, 91, 212, 146, 23, 27, 83, 235, 208, 126, 175, 137, 47, 193, 209, 10, 29, 183, 180, 128, 250, 246, 185, 207, 87, 7, 56, 68, 167, 166, 211, 122, 98, 40, 254, 231, 154, 233, 34, 221, 72, 174, 0, 1})
	var y1 fp.Element
	y1.Exp(tv1, &c1)
	y1.Mul(&y1, &tv2)
	tv1.Mul(&y1, &fp.Element{7899625277197386435, 5217716493391639390, 7472932469883704682, 7632350077606897049, 9296070723299766388, 14353472371414671016, 14644604696869838127, 11421353192299464576, 237964513547175570, 46667570639865841})
	tv2.Square(&tv1)

	//Line 10 in std doc
	tv2.Mul(&tv2, v)

	y1.Select(int(tv2.NotEqual(u)), &tv1, &y1)

	tv2.Square(&y1)
	tv2.Mul(&tv2, v)

	//Line 15
	isQNr := tv2.NotEqual(u)
	var y2 fp.Element
	y2.Mul(&y1, &fp.Element{16212120288951005687, 11690167560162600414, 9845362566212292170, 5006379754746321817, 3559960229467473872, 1378556217976105943, 4841104984578141598, 15436992508257808297, 6778583767067406308, 4544728946065242})
	tv1.Mul(&y2, &fp.Element{7899625277197386435, 5217716493391639390, 7472932469883704682, 7632350077606897049, 9296070723299766388, 14353472371414671016, 14644604696869838127, 11421353192299464576, 237964513547175570, 46667570639865841})
	tv2.Square(&tv1)
	tv2.Mul(&tv2, v)

	var tv3 fp.Element
	//Line 20
	// Z = [2]
	g2MulByZ(&tv3, u)

	y2.Select(int(tv2.NotEqual(&tv3)), &tv1, &y2)

	z.Select(int(isQNr), &y1, &y2)
	return isQNr
}

/*
// g2SetZ sets z to [2].
func g2SetZ(z *fp.Element) {
    z.Set( &fp.Element  { 14263791471689722215, 10958139817512614717, 646289283071182148, 16194112285086178910, 12391927829343171647, 3698619178316197998, 14879001273850772332, 4646357410414107532, 14313982959885664825, 19561843432566578 } )
}*/

// g2MulByZ multiplies x by [2] and stores the result in z
func g2MulByZ(z *fp.Element, x *fp.Element) {

	res := *x

	res.Double(&res)

	*z = res
}

// From https://datatracker.ietf.org/doc/draft-irtf-cfrg-hash-to-curve/13/ Pg 80
// sswuMapG2 implements the SSWU map
// No cofactor clearing
func sswuMapG2(u *fp.Element) G2Affine {

	var tv1 fp.Element
	tv1.Square(u)

	//mul tv1 by Z
	g2MulByZ(&tv1, &tv1)

	var tv2 fp.Element
	tv2.Square(&tv1)
	tv2.Add(&tv2, &tv1)

	var tv3 fp.Element
	//Standard doc line 5
	var tv4 fp.Element
	tv4.SetOne()
	tv3.Add(&tv2, &tv4)
	tv3.Mul(&tv3, &fp.Element{4170590011558214244, 9101648159034903675, 4256739633972552875, 7483080556638609334, 12430228215152656439, 9977400640742476476, 15847011074743951739, 17768582661138350292, 10869631430819016060, 64187107279947172})

	tv2NZero := g2NotZero(&tv2)

	// tv4 = Z
	tv4 = fp.Element{14263791471689722215, 10958139817512614717, 646289283071182148, 16194112285086178910, 12391927829343171647, 3698619178316197998, 14879001273850772332, 4646357410414107532, 14313982959885664825, 19561843432566578}

	tv2.Neg(&tv2)
	tv4.Select(int(tv2NZero), &tv4, &tv2)
	tv2 = fp.Element{13503940466125084703, 3000707982748310797, 1529397070312683242, 9240962296298654443, 4577258595340312235, 16046828875439788343, 7236093083337192433, 2860564553402019540, 5160479239841632821, 65394042426465165}
	tv4.Mul(&tv4, &tv2)

	tv2.Square(&tv3)

	var tv6 fp.Element
	//Standard doc line 10
	tv6.Square(&tv4)

	var tv5 fp.Element
	tv5.Mul(&tv6, &fp.Element{13503940466125084703, 3000707982748310797, 1529397070312683242, 9240962296298654443, 4577258595340312235, 16046828875439788343, 7236093083337192433, 2860564553402019540, 5160479239841632821, 65394042426465165})

	tv2.Add(&tv2, &tv5)
	tv2.Mul(&tv2, &tv3)
	tv6.Mul(&tv6, &tv4)

	//Standards doc line 15
	tv5.Mul(&tv6, &fp.Element{4170590011558214244, 9101648159034903675, 4256739633972552875, 7483080556638609334, 12430228215152656439, 9977400640742476476, 15847011074743951739, 17768582661138350292, 10869631430819016060, 64187107279947172})
	tv2.Add(&tv2, &tv5)

	var x fp.Element
	x.Mul(&tv1, &tv3)

	var y1 fp.Element
	gx1NSquare := g2SqrtRatio(&y1, &tv2, &tv6)

	var y fp.Element
	y.Mul(&tv1, u)

	//Standards doc line 20
	y.Mul(&y, &y1)

	x.Select(int(gx1NSquare), &tv3, &x)
	y.Select(int(gx1NSquare), &y1, &y)

	y1.Neg(&y)
	y.Select(int(g2Sgn0(u)^g2Sgn0(&y)), &y, &y1)

	//Standards doc line 25
	x.Div(&x, &tv4)

	return G2Affine{x, y}
}

// mapToG2 invokes the SSWU map, and guarantees that the result is in g2
func mapToG2(u fp.Element) G2Affine {
	res := sswuMapG2(&u)
	res.ClearCofactor(&res)
	return res
}

// EncodeToG2 maps a fp.Element to a point on the curve using the Simplified Shallue and van de Woestijne Ulas map
//https://datatracker.ietf.org/doc/draft-irtf-cfrg-hash-to-curve/13/#section-6.6.3
func EncodeToG2(msg, dst []byte) (G2Affine, error) {

	var res G2Affine
	u, err := hashToFp(msg, dst, 1)
	if err != nil {
		return res, err
	}

	res = sswuMapG2(&u[0])

	//this is in an isogenous curve
	g2Isogeny(&res)

	res.ClearCofactor(&res)

	return res, nil
}

// HashToG2 hashes a byte string to the G2 curve. Usable as a random oracle.
// https://tools.ietf.org/html/draft-irtf-cfrg-hash-to-curve-06#section-3
func HashToG2(msg, dst []byte) (G2Affine, error) {
	u, err := hashToFp(msg, dst, 2*1)
	if err != nil {
		return G2Affine{}, err
	}

	Q0 := sswuMapG2(&u[0])
	Q1 := sswuMapG2(&u[1])

	//TODO: Add in E' first, then apply isogeny
	g2Isogeny(&Q0)
	g2Isogeny(&Q1)

	var _Q0, _Q1 G2Jac
	_Q0.FromAffine(&Q0)
	_Q1.FromAffine(&Q1).AddAssign(&_Q0)

	_Q1.ClearCofactor(&_Q1)

	Q1.FromJacobian(&_Q1)
	return Q1, nil
}

// g2Sgn0 is an algebraic substitute for the notion of sign in ordered fields
// Namely, every non-zero quadratic residue in a finite field of characteristic =/= 2 has exactly two square roots, one of each sign
// Taken from https://datatracker.ietf.org/doc/draft-irtf-cfrg-hash-to-curve/ section 4.1
// The sign of an element is not obviously related to that of its Montgomery form
func g2Sgn0(z *fp.Element) uint64 {

	nonMont := *z
	nonMont.FromMont()

	return nonMont[0] % 2

}

func g2EvalPolynomial(z *fp.Element, monic bool, coefficients []fp.Element, x *fp.Element) {
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

func g2NotZero(x *fp.Element) uint64 {

	return x[0] | x[1] | x[2] | x[3] | x[4] | x[5] | x[6] | x[7] | x[8] | x[9]

}
