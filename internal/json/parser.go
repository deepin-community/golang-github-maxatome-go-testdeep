// Code generated by goyacc -l -o parser.go parser.y. DO NOT EDIT.
// Copyright (c) 2020-2022, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package json

import __yyfmt__ "fmt"

type Placeholder struct {
	Num  int
	Name string
}

type Operator struct {
	Name   string
	Params []any
}

type member struct {
	key   string
	value any
}

func finalize(l yyLexer, value any) {
	l.(*json).value = value
}

type yySymType struct {
	yys    int
	object map[string]any
	member member
	array  []any
	string string
	value  any
}

const TRUE = 57346
const FALSE = 57347
const NULL = 57348
const NUMBER = 57349
const PLACEHOLDER = 57350
const SUB_PARSER = 57351
const STRING = 57352
const OPERATOR = 57353

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TRUE",
	"FALSE",
	"NULL",
	"NUMBER",
	"PLACEHOLDER",
	"SUB_PARSER",
	"STRING",
	"OPERATOR",
	"'{'",
	"'}'",
	"','",
	"':'",
	"'['",
	"']'",
	"'('",
	"')'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 86

var yyAct = [...]int{
	22, 2, 9, 10, 11, 8, 12, 6, 7, 15,
	13, 21, 24, 37, 14, 27, 5, 39, 38, 9,
	10, 11, 8, 12, 6, 7, 15, 13, 34, 23,
	36, 14, 25, 26, 30, 18, 31, 4, 36, 9,
	10, 11, 8, 12, 6, 7, 15, 13, 17, 3,
	1, 14, 35, 9, 10, 11, 8, 12, 6, 7,
	15, 13, 33, 0, 0, 14, 20, 9, 10, 11,
	8, 12, 6, 7, 15, 13, 0, 19, 29, 14,
	32, 28, 19, 0, 0, 16,
}

var yyPact = [...]int{
	63, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 72, 49, -6, -1000, 19, -1000, 0,
	-1000, 64, -1000, -1000, 15, -1000, 67, 63, -1000, 35,
	-1000, -1, -1000, -1000, -1000, -1000, -1000, -2, -1000, -1000,
}

var yyPgo = [...]int{
	0, 50, 49, 48, 35, 37, 11, 29, 0, 16,
}

var yyR1 = [...]int{
	0, 1, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 2, 2, 2, 3, 3, 4, 5, 5,
	5, 6, 6, 7, 7, 7, 9, 9,
}

var yyR2 = [...]int{
	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 3, 4, 1, 3, 3, 2, 3,
	4, 1, 3, 2, 3, 4, 2, 1,
}

var yyChk = [...]int{
	-1000, -1, -8, -2, -5, -9, 9, 10, 7, 4,
	5, 6, 8, 12, 16, 11, 13, -3, -4, 10,
	17, -6, -8, -7, 18, 13, 14, 15, 17, 14,
	19, -6, 13, -4, -8, 17, -8, 14, 19, 19,
}

var yyDef = [...]int{
	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 0, 0, 27, 12, 0, 15, 0,
	18, 0, 21, 26, 0, 13, 0, 0, 19, 0,
	23, 0, 14, 16, 17, 20, 22, 0, 24, 25,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	18, 19, 3, 3, 14, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 15, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 16, 3, 17, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 12, 3, 13,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = yyDollar[1].value
			finalize(yylex, yyVAL.value)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = yyDollar[1].object
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = yyDollar[1].array
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = yyDollar[1].value
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = yyDollar[1].value
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = yyDollar[1].string
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.object = map[string]any{}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.object = yyDollar[2].object
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.object = yyDollar[2].object
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.object = map[string]any{
				yyDollar[1].member.key: yyDollar[1].member.value,
			}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].object[yyDollar[3].member.key] = yyDollar[3].member.value
			yyVAL.object = yyDollar[1].object
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.member = member{
				key:   yyDollar[1].string,
				value: yyDollar[3].value,
			}
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.array = []any{}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.array = yyDollar[2].array
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.array = yyDollar[2].array
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.array = []any{yyDollar[1].value}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.array = append(yyDollar[1].array, yyDollar[3].value)
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.array = []any{}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.array = yyDollar[2].array
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.array = yyDollar[2].array
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			op := yylex.(*json).newOperator(yyDollar[1].string, yyDollar[2].array)
			if op == nil {
				return 1
			}
			yyVAL.value = op
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			op := yylex.(*json).newOperator(yyDollar[1].string, nil)
			if op == nil {
				return 1
			}
			yyVAL.value = op
		}
	}
	goto yystack /* stack new state and value */
}
