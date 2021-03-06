// Code generated by goyacc -o fss.y.go -p fss fss.y. DO NOT EDIT.

//line fss.y:8
package fss

import __yyfmt__ "fmt"

//line fss.y:8

import "strings"

var (
	print  = __yyfmt__.Print
	printf = __yyfmt__.Printf
)

//line fss.y:20
type fssSymType struct {
	yys int
	// flow
	Flow        string
	Steps       []Step
	_step       Step
	_return     Return
	_returns    Returns
	_action     Action
	_string     string
	_identifier string
	_list       []interface{}
	_args       []Param
	_param      Param
	_dict       map[string]interface{}
	_variable   string
	_number     int64

	// action
	ActionStatement
	_addr []string
	_type ActionMethodType
	_grpc ActionMethodType
	_http ActionMethodType
}

const ILLEGAL = 57346
const EOL = 57347
const IDENTIFIER = 57348
const NUMBER_VALUE = 57349
const ID = 57350
const STRING_VALUE = 57351
const LIST = 57352
const DICT = 57353
const FLOW = 57354
const FLOW_END = 57355
const STEP = 57356
const ACTION = 57357
const ARGS = 57358
const DECI = 57359
const ACTION_END = 57360
const ADDR = 57361
const METHOD = 57362
const FLOW_RUN = 57363
const FLOW_RUN_END = 57364
const RETURN = 57365
const LPAREN = 57366
const RPAREN = 57367
const LSQUARE = 57368
const RSQUARE = 57369
const LCURLY = 57370
const RCURLY = 57371
const SEMICOLON = 57372
const COMMA = 57373
const COLON = 57374
const HTTP = 57375
const GRPC = 57376
const INT = 57377
const STR = 57378
const ASSIGN = 57379
const OR = 57380
const AND = 57381
const TO = 57382
const DEST = 57383

var fssToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ILLEGAL",
	"EOL",
	"IDENTIFIER",
	"NUMBER_VALUE",
	"ID",
	"STRING_VALUE",
	"LIST",
	"DICT",
	"FLOW",
	"FLOW_END",
	"STEP",
	"ACTION",
	"ARGS",
	"DECI",
	"ACTION_END",
	"ADDR",
	"METHOD",
	"FLOW_RUN",
	"FLOW_RUN_END",
	"RETURN",
	"LPAREN",
	"RPAREN",
	"LSQUARE",
	"RSQUARE",
	"LCURLY",
	"RCURLY",
	"SEMICOLON",
	"COMMA",
	"COLON",
	"HTTP",
	"GRPC",
	"INT",
	"STR",
	"ASSIGN",
	"OR",
	"AND",
	"TO",
	"DEST",
}

var fssStatenames = [...]string{}

const fssEofCode = 1
const fssErrCode = 2
const fssInitialStackSize = 16

//line fss.y:259

//line yacctab:1
var fssExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const fssPrivate = 57344

const fssLast = 133

var fssAct = [...]int{
	115, 108, 87, 68, 52, 64, 81, 44, 43, 45,
	42, 41, 86, 79, 119, 88, 107, 106, 88, 89,
	90, 117, 89, 90, 36, 113, 80, 92, 58, 47,
	38, 28, 48, 49, 117, 111, 103, 102, 85, 84,
	114, 83, 82, 76, 74, 91, 116, 54, 55, 56,
	51, 62, 61, 50, 40, 65, 69, 109, 53, 116,
	70, 71, 72, 73, 60, 23, 23, 46, 24, 24,
	105, 26, 17, 25, 21, 67, 7, 27, 63, 6,
	57, 104, 31, 23, 93, 5, 24, 37, 30, 23,
	29, 95, 24, 14, 78, 120, 16, 121, 101, 99,
	39, 100, 98, 15, 94, 69, 19, 20, 18, 110,
	97, 34, 32, 96, 35, 33, 13, 118, 11, 12,
	9, 10, 66, 8, 112, 77, 22, 75, 59, 4,
	3, 2, 1,
}

var fssPact = [...]int{
	64, -1000, -1000, -1000, -1000, 114, 112, 110, -1000, -1000,
	53, 53, -1000, -1000, 52, 51, 57, -6, 57, 75,
	69, -1000, -1000, 106, 105, -1000, 71, -7, 91, 71,
	-1000, -1000, -29, -30, -32, -33, 44, -8, -1, 23,
	44, 34, 34, 34, 34, 62, -9, 40, 22, 21,
	-1000, 60, 27, 50, 27, 27, 27, -1000, 34, 14,
	-1000, -1000, -1000, -1000, 13, 79, -12, -1000, -1000, -35,
	12, 11, 9, 8, -1000, -13, -1000, 16, -10, -1000,
	99, 98, -1000, -1000, -1000, -1000, -1000, -1000, -16, 104,
	93, -1000, 92, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	7, 6, 65, 54, -20, -21, 33, 33, 5, -1000,
	-5, -1000, 15, -1000, -1000, -1000, 28, -23, -1000, 88,
	-1000, -1000,
}

var fssPgo = [...]int{
	0, 132, 131, 130, 129, 96, 71, 24, 9, 128,
	127, 2, 4, 93, 126, 5, 125, 1, 124, 0,
	122, 3,
}

var fssR1 = [...]int{
	0, 1, 1, 1, 3, 3, 5, 5, 6, 6,
	6, 7, 7, 9, 10, 10, 11, 11, 11, 11,
	11, 8, 4, 4, 2, 2, 13, 13, 14, 14,
	14, 14, 15, 16, 16, 16, 16, 17, 18, 18,
	19, 19, 19, 12, 12, 20, 20, 21, 21,
}

var fssR2 = [...]int{
	0, 1, 1, 1, 7, 7, 0, 4, 0, 4,
	4, 0, 4, 3, 0, 2, 2, 2, 2, 2,
	2, 4, 4, 4, 4, 4, 0, 2, 6, 6,
	6, 6, 3, 8, 8, 4, 4, 3, 0, 2,
	2, 3, 3, 3, 2, 1, 3, 3, 1,
}

var fssChk = [...]int{
	-1000, -1, -2, -3, -4, 21, 15, 12, 9, 6,
	9, 6, 9, 6, -13, -13, -5, 19, -5, -13,
	-13, 22, -14, 14, 17, 22, -6, 20, 37, -6,
	13, 13, 6, 9, 6, 9, -7, 16, 37, 9,
	-7, 40, 40, 40, 40, -8, 23, 37, 33, 34,
	30, -8, -12, 24, -12, -12, -12, 18, 37, -9,
	24, 30, 30, 18, -15, 28, -20, 25, -21, 6,
	-15, -15, -15, -12, 30, -10, 30, -16, 15, 25,
	38, 41, 30, 30, 30, 30, 25, -11, 31, 35,
	36, 29, 37, -21, 6, -11, 9, 6, 9, 6,
	9, 6, 30, 30, 16, 16, 37, 37, -17, 24,
	-17, 30, -18, 30, 25, -19, 31, 6, -19, 37,
	7, 9,
}

var fssDef = [...]int{
	0, -2, 1, 2, 3, 0, 0, 0, 26, 26,
	6, 6, 26, 26, 0, 0, 8, 0, 8, 0,
	0, 24, 27, 0, 0, 25, 11, 0, 0, 11,
	22, 23, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	7, 0, 0, 0, 0, 0, 0, 4, 0, 0,
	14, 9, 10, 5, 0, 0, 0, 44, 45, 48,
	0, 0, 0, 0, 12, 0, 28, 0, 0, 43,
	0, 0, 30, 29, 31, 21, 13, 15, 0, 0,
	0, 32, 0, 46, 47, 16, 17, 19, 18, 20,
	0, 0, 35, 36, 0, 0, 0, 0, 0, 38,
	0, 33, 0, 34, 37, 39, 0, 0, 40, 0,
	41, 42,
}

var fssTok1 = [...]int{
	1,
}

var fssTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
}

var fssTok3 = [...]int{
	0,
}

var fssErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	fssDebug        = 0
	fssErrorVerbose = false
)

type fssLexer interface {
	Lex(lval *fssSymType) int
	Error(s string)
}

type fssParser interface {
	Parse(fssLexer) int
	Lookahead() int
}

type fssParserImpl struct {
	lval  fssSymType
	stack [fssInitialStackSize]fssSymType
	char  int
}

func (p *fssParserImpl) Lookahead() int {
	return p.char
}

func fssNewParser() fssParser {
	return &fssParserImpl{}
}

const fssFlag = -1000

func fssTokname(c int) string {
	if c >= 1 && c-1 < len(fssToknames) {
		if fssToknames[c-1] != "" {
			return fssToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func fssStatname(s int) string {
	if s >= 0 && s < len(fssStatenames) {
		if fssStatenames[s] != "" {
			return fssStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func fssErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !fssErrorVerbose {
		return "syntax error"
	}

	for _, e := range fssErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + fssTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := fssPact[state]
	for tok := TOKSTART; tok-1 < len(fssToknames); tok++ {
		if n := base + tok; n >= 0 && n < fssLast && fssChk[fssAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if fssDef[state] == -2 {
		i := 0
		for fssExca[i] != -1 || fssExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; fssExca[i] >= 0; i += 2 {
			tok := fssExca[i]
			if tok < TOKSTART || fssExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if fssExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += fssTokname(tok)
	}
	return res
}

func fsslex1(lex fssLexer, lval *fssSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = fssTok1[0]
		goto out
	}
	if char < len(fssTok1) {
		token = fssTok1[char]
		goto out
	}
	if char >= fssPrivate {
		if char < fssPrivate+len(fssTok2) {
			token = fssTok2[char-fssPrivate]
			goto out
		}
	}
	for i := 0; i < len(fssTok3); i += 2 {
		token = fssTok3[i+0]
		if token == char {
			token = fssTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = fssTok2[1] /* unknown char */
	}
	if fssDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", fssTokname(token), uint(char))
	}
	return char, token
}

func fssParse(fsslex fssLexer) int {
	return fssNewParser().Parse(fsslex)
}

func (fssrcvr *fssParserImpl) Parse(fsslex fssLexer) int {
	var fssn int
	var fssVAL fssSymType
	var fssDollar []fssSymType
	_ = fssDollar // silence set and not used
	fssS := fssrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	fssstate := 0
	fssrcvr.char = -1
	fsstoken := -1 // fssrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		fssstate = -1
		fssrcvr.char = -1
		fsstoken = -1
	}()
	fssp := -1
	goto fssstack

ret0:
	return 0

ret1:
	return 1

fssstack:
	/* put a state and value onto the stack */
	if fssDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", fssTokname(fsstoken), fssStatname(fssstate))
	}

	fssp++
	if fssp >= len(fssS) {
		nyys := make([]fssSymType, len(fssS)*2)
		copy(nyys, fssS)
		fssS = nyys
	}
	fssS[fssp] = fssVAL
	fssS[fssp].yys = fssstate

fssnewstate:
	fssn = fssPact[fssstate]
	if fssn <= fssFlag {
		goto fssdefault /* simple state */
	}
	if fssrcvr.char < 0 {
		fssrcvr.char, fsstoken = fsslex1(fsslex, &fssrcvr.lval)
	}
	fssn += fsstoken
	if fssn < 0 || fssn >= fssLast {
		goto fssdefault
	}
	fssn = fssAct[fssn]
	if fssChk[fssn] == fsstoken { /* valid shift */
		fssrcvr.char = -1
		fsstoken = -1
		fssVAL = fssrcvr.lval
		fssstate = fssn
		if Errflag > 0 {
			Errflag--
		}
		goto fssstack
	}

fssdefault:
	/* default state action */
	fssn = fssDef[fssstate]
	if fssn == -2 {
		if fssrcvr.char < 0 {
			fssrcvr.char, fsstoken = fsslex1(fsslex, &fssrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if fssExca[xi+0] == -1 && fssExca[xi+1] == fssstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			fssn = fssExca[xi+0]
			if fssn < 0 || fssn == fsstoken {
				break
			}
		}
		fssn = fssExca[xi+1]
		if fssn < 0 {
			goto ret0
		}
	}
	if fssn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			fsslex.Error(fssErrorMessage(fssstate, fsstoken))
			Nerrs++
			if fssDebug >= 1 {
				__yyfmt__.Printf("%s", fssStatname(fssstate))
				__yyfmt__.Printf(" saw %s\n", fssTokname(fsstoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for fssp >= 0 {
				fssn = fssPact[fssS[fssp].yys] + fssErrCode
				if fssn >= 0 && fssn < fssLast {
					fssstate = fssAct[fssn] /* simulate a shift of "error" */
					if fssChk[fssstate] == fssErrCode {
						goto fssstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if fssDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", fssS[fssp].yys)
				}
				fssp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if fssDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", fssTokname(fsstoken))
			}
			if fsstoken == fssEofCode {
				goto ret1
			}
			fssrcvr.char = -1
			fsstoken = -1
			goto fssnewstate /* try again in the same state */
		}
	}

	/* reduction by production fssn */
	if fssDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", fssn, fssStatname(fssstate))
	}

	fssnt := fssn
	fsspt := fssp
	_ = fsspt // guard against "declared and not used"

	fssp -= fssR2[fssn]
	// fssp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if fssp+1 >= len(fssS) {
		nyys := make([]fssSymType, len(fssS)*2)
		copy(nyys, fssS)
		fssS = nyys
	}
	fssVAL = fssS[fssp+1]

	/* consult goto table to find next state */
	fssn = fssR1[fssn]
	fssg := fssPgo[fssn]
	fssj := fssg + fssS[fssp].yys + 1

	if fssj >= fssLast {
		fssstate = fssAct[fssg]
	} else {
		fssstate = fssAct[fssj]
		if fssChk[fssstate] != -fssn {
			fssstate = fssAct[fssg]
		}
	}
	// dummy call; replaced with literal code
	switch fssnt {

	case 1:
		fssDollar = fssS[fsspt-1 : fsspt+1]
//line fss.y:65
		{
			flowRunSymPoolPut(fssDollar[1].Flow, fssDollar[1])
		}
	case 2:
		fssDollar = fssS[fsspt-1 : fsspt+1]
//line fss.y:69
		{
			actionSymPoolPut(fssDollar[1].ActionStatement.Name, fssDollar[1])
		}
	case 3:
		fssDollar = fssS[fsspt-1 : fsspt+1]
//line fss.y:73
		{
			flowSymPoolPut(fssDollar[1].Flow, fssDollar[1])
		}
	case 4:
		fssDollar = fssS[fsspt-7 : fsspt+1]
//line fss.y:80
		{
			fssVAL.ActionStatement = ActionStatement{
				Name:    fssDollar[2]._string,
				Addr:    fssDollar[3]._addr,
				Type:    fssDollar[4]._type,
				Args:    fssDollar[5]._args,
				Returns: fssDollar[6]._returns,
			}
		}
	case 5:
		fssDollar = fssS[fsspt-7 : fsspt+1]
//line fss.y:91
		{
			fssVAL.ActionStatement = ActionStatement{
				Name:    fssDollar[2]._identifier,
				Addr:    fssDollar[3]._addr,
				Type:    fssDollar[4]._type,
				Args:    fssDollar[5]._args,
				Returns: fssDollar[6]._returns,
			}
		}
	case 7:
		fssDollar = fssS[fsspt-4 : fsspt+1]
//line fss.y:103
		{
			fssVAL._addr = append(fssVAL._addr, strings.Split(strings.Trim(fssDollar[3]._string, "\""), ",")...)
		}
	case 9:
		fssDollar = fssS[fsspt-4 : fsspt+1]
//line fss.y:107
		{
			fssVAL._type = fssDollar[3]._http
		}
	case 10:
		fssDollar = fssS[fsspt-4 : fsspt+1]
//line fss.y:108
		{
			fssVAL._type = fssDollar[3]._grpc
		}
	case 12:
		fssDollar = fssS[fsspt-4 : fsspt+1]
//line fss.y:112
		{
			fssVAL._args = fssDollar[3]._args
		}
	case 13:
		fssDollar = fssS[fsspt-3 : fsspt+1]
//line fss.y:117
		{
			fssVAL = fssDollar[2]
		}
	case 15:
		fssDollar = fssS[fsspt-2 : fsspt+1]
//line fss.y:121
		{
			fssVAL._args = append(fssVAL._args, fssDollar[2]._param)
		}
	case 16:
		fssDollar = fssS[fsspt-2 : fsspt+1]
//line fss.y:125
		{
			fssVAL._param = fssDollar[2]._param
		}
	case 17:
		fssDollar = fssS[fsspt-2 : fsspt+1]
//line fss.y:126
		{
			fssVAL._param = Param{Name: fssDollar[2]._string, ParamType: NumberType}
		}
	case 18:
		fssDollar = fssS[fsspt-2 : fsspt+1]
//line fss.y:127
		{
			fssVAL._param = Param{Name: fssDollar[2]._string, ParamType: StringType}
		}
	case 19:
		fssDollar = fssS[fsspt-2 : fsspt+1]
//line fss.y:128
		{
			fssVAL._param = Param{Name: fssDollar[2]._identifier, ParamType: NumberType}
		}
	case 20:
		fssDollar = fssS[fsspt-2 : fsspt+1]
//line fss.y:129
		{
			fssVAL._param = Param{Name: fssDollar[2]._identifier, ParamType: StringType}
		}
	case 21:
		fssDollar = fssS[fsspt-4 : fsspt+1]
//line fss.y:134
		{
			fssVAL._returns = fssDollar[3]._returns
		}
	case 22:
		fssDollar = fssS[fsspt-4 : fsspt+1]
//line fss.y:141
		{
			fssVAL.Flow = fssDollar[2]._string
			fssVAL.Steps = fssDollar[3].Steps
		}
	case 23:
		fssDollar = fssS[fsspt-4 : fsspt+1]
//line fss.y:146
		{
			fssVAL.Flow = fssDollar[2]._identifier
			fssVAL.Steps = fssDollar[3].Steps
		}
	case 24:
		fssDollar = fssS[fsspt-4 : fsspt+1]
//line fss.y:153
		{
			fssVAL.Flow = fssDollar[2]._string
			fssVAL.Steps = fssDollar[3].Steps
		}
	case 25:
		fssDollar = fssS[fsspt-4 : fsspt+1]
//line fss.y:158
		{
			fssVAL.Flow = fssDollar[2]._identifier
			fssVAL.Steps = fssDollar[3].Steps
		}
	case 27:
		fssDollar = fssS[fsspt-2 : fsspt+1]
//line fss.y:166
		{
			fssVAL.Steps = append(fssVAL.Steps, fssDollar[2]._step)
		}
	case 28:
		fssDollar = fssS[fsspt-6 : fsspt+1]
//line fss.y:173
		{
			fssVAL._step = Step{Name: fssDollar[2]._identifier, Action: fssDollar[5]._action, Returns: fssDollar[4]._returns, StepType: Normal}
		}
	case 29:
		fssDollar = fssS[fsspt-6 : fsspt+1]
//line fss.y:177
		{
			fssVAL._step = Step{Name: fssDollar[2]._identifier, Action: fssDollar[5]._action, Returns: fssDollar[4]._returns, StepType: Decision}
		}
	case 30:
		fssDollar = fssS[fsspt-6 : fsspt+1]
//line fss.y:181
		{
			fssVAL._step = Step{Name: fssDollar[2]._string, Action: fssDollar[5]._action, Returns: fssDollar[4]._returns, StepType: Normal}
		}
	case 31:
		fssDollar = fssS[fsspt-6 : fsspt+1]
//line fss.y:185
		{
			fssVAL._step = Step{Name: fssDollar[2]._string, Action: fssDollar[5]._action, Returns: fssDollar[4]._returns, StepType: Decision}
		}
	case 32:
		fssDollar = fssS[fsspt-3 : fsspt+1]
//line fss.y:191
		{
			fssVAL = fssDollar[2]
		}
	case 33:
		fssDollar = fssS[fsspt-8 : fsspt+1]
//line fss.y:196
		{
			fssVAL._action = Action{Name: fssDollar[3]._string, Args: fssDollar[7]._args}
		}
	case 34:
		fssDollar = fssS[fsspt-8 : fsspt+1]
//line fss.y:201
		{
			fssVAL._action = Action{Name: fssDollar[3]._identifier, Args: fssDollar[7]._args}
		}
	case 35:
		fssDollar = fssS[fsspt-4 : fsspt+1]
//line fss.y:206
		{
			fssVAL._action = Action{Name: fssDollar[3]._string}
		}
	case 36:
		fssDollar = fssS[fsspt-4 : fsspt+1]
//line fss.y:211
		{
			fssVAL._action = Action{Name: fssDollar[3]._identifier}
		}
	case 37:
		fssDollar = fssS[fsspt-3 : fsspt+1]
//line fss.y:217
		{
			fssVAL = fssDollar[2]
		}
	case 39:
		fssDollar = fssS[fsspt-2 : fsspt+1]
//line fss.y:222
		{
			fssVAL._args = append(fssVAL._args, fssDollar[2]._param)
		}
	case 40:
		fssDollar = fssS[fsspt-2 : fsspt+1]
//line fss.y:228
		{
			fssVAL._param = fssDollar[2]._param
		}
	case 41:
		fssDollar = fssS[fsspt-3 : fsspt+1]
//line fss.y:229
		{
			fssVAL._param = Param{Name: fssDollar[1]._identifier, ParamType: NumberType, Value: fssDollar[3]._number}
		}
	case 42:
		fssDollar = fssS[fsspt-3 : fsspt+1]
//line fss.y:230
		{
			fssVAL._param = Param{Name: fssDollar[1]._identifier, ParamType: StringType, Value: fssDollar[3]._string}
		}
	case 43:
		fssDollar = fssS[fsspt-3 : fsspt+1]
//line fss.y:234
		{
			fssVAL = fssDollar[2]
		}
	case 44:
		fssDollar = fssS[fsspt-2 : fsspt+1]
//line fss.y:235
		{
			fssVAL._returns = append(fssVAL._returns, Return{State: "DONE", Next: ""})
		}
	case 45:
		fssDollar = fssS[fsspt-1 : fsspt+1]
//line fss.y:240
		{
			fssVAL._returns = append(fssVAL._returns, fssDollar[1]._return)
		}
	case 46:
		fssDollar = fssS[fsspt-3 : fsspt+1]
//line fss.y:244
		{
			fssVAL._returns = append(fssVAL._returns, fssDollar[3]._return)
		}
	case 47:
		fssDollar = fssS[fsspt-3 : fsspt+1]
//line fss.y:250
		{
			fssVAL._return = Return{State: fssDollar[1]._identifier, Next: fssDollar[3]._identifier}
		}
	case 48:
		fssDollar = fssS[fsspt-1 : fsspt+1]
//line fss.y:254
		{
			fssVAL._return = Return{State: fssDollar[1]._identifier}
		}
	}
	goto fssstack /* stack new state and value */
}
