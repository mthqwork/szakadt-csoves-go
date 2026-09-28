// Machine-translated from CSOVES.EXE (BASIC main module); maintained as source.

package game

import "math"

// mainProgram (1010:0030)
func (g *Game) mainProgram() {
	var l0c int
	_ = l0c
	var fcarry float64
	goto L0030
gosubReturn:
	switch g.gosubPop() {
	case 23:
		goto R019c
	case 22:
		goto R01e9
	case 7:
		goto R0392
	case 8:
		goto R0553
	case 9:
		goto R05da
	case 10:
		goto R0661
	case 11:
		goto R0769
	case 12:
		goto R0819
	case 13:
		goto R0871
	case 14:
		goto R08b1
	case 15:
		goto R0981
	case 16:
		goto R09d9
	case 17:
		goto R0aab
	case 18:
		goto R0c82
	case 19:
		goto R0cc4
	case 20:
		goto R0e28
	case 21:
		goto R0ec2
	case 3:
		goto R1022
	case 4:
		goto R144a
	case 94:
		goto R168d
	case 95:
		goto R178e
	case 24:
		goto R1962
	case 25:
		goto R1a19
	case 88:
		goto R1d1d
	case 89:
		goto R1eeb
	case 90:
		goto R1fb0
	case 91:
		goto R2095
	case 83:
		goto R226a
	case 84:
		goto R2447
	case 85:
		goto R25f9
	case 86:
		goto R261e
	case 87:
		goto R270b
	case 78:
		goto R28c8
	case 79:
		goto R2aff
	case 80:
		goto R2c78
	case 81:
		goto R2d5d
	case 82:
		goto R2e0d
	case 72:
		goto R2fd2
	case 73:
		goto R3150
	case 74:
		goto R3214
	case 75:
		goto R3277
	case 76:
		goto R336d
	case 77:
		goto R3454
	case 61:
		goto R354b
	case 62:
		goto R376c
	case 63:
		goto R3791
	case 64:
		goto R3805
	case 65:
		goto R3853
	case 66:
		goto R389f
	case 67:
		goto R38d0
	case 68:
		goto R39ae
	case 69:
		goto R3ab4
	case 70:
		goto R3ae5
	case 71:
		goto R3bef
	case 48:
		goto R3e0b
	case 49:
		goto R3e9a
	case 50:
		goto R3f18
	case 51:
		goto R3f42
	case 52:
		goto R3fc3
	case 53:
		goto R402e
	case 54:
		goto R409a
	case 55:
		goto R40bf
	case 56:
		goto R4110
	case 57:
		goto R4152
	case 58:
		goto R419e
	case 59:
		goto R4215
	case 60:
		goto R42c3
	case 40:
		goto R4450
	case 41:
		goto R46d0
	case 42:
		goto R4701
	case 43:
		goto R4752
	case 44:
		goto R4794
	case 45:
		goto R47e0
	case 46:
		goto R4857
	case 47:
		goto R48e0
	case 26:
		goto R4aa2
	case 27:
		goto R4b89
	case 28:
		goto R4c5c
	case 29:
		goto R4cf8
	case 30:
		goto R4dd8
	case 31:
		goto R4dfe
	case 32:
		goto R4e99
	case 35:
		goto R4ecb
	case 33:
		goto R4f3a
	case 34:
		goto R4f60
	case 36:
		goto R5036
	case 37:
		goto R510e
	case 38:
		goto R51d3
	case 39:
		goto R5210
	case 5:
		goto R5333
	case 6:
		goto R5388
	case 2:
		goto R5b5b
	case 1:
		goto R5c3e
	case 92:
		goto R5d2d
	}
	panic("bad gosub return")
L0030:
	g.sub_title()
	g.sub_initart()
	g.sub_7a66()
	g.sub_6ee3()
	goto L521b
L0047:
	g.s13e = g.inputS(1)
L005f:
	g.f142 = float32(0)
	g.s14e = mid(g.a0da[cint(float64(g.f146))], cint(float64(g.f14a)), 1)
	g.f152 = float32(0)
	g.f15a = float32(float64(len(g.s156)))
	fcarry = 1
	goto L0102
L00c4:
	if mid(g.s156, cint(float64(g.f15e)), 1) == g.s14e {
		goto L00ec
	}
	goto L00f8
L00ec:
	g.f152 = float32(float64(g.f15e))
L00f8:
	fcarry = (float64(g.f15e) + 1)
L0102:
	g.f15e = float32(fcarry)
	if float64(g.f15a) >= float64(g.f15e) {
		goto L00c4
	}
	g.locate(20, 1)
	g.color(7)
	switch cint(float64(g.f152)) {
	case 1:
		goto L18f2
	case 2:
		goto L016d
	case 3:
		goto L01c6
	case 4:
		goto L020d
	case 5:
		goto L0f41
	case 6:
		goto L158b
	case 7:
		goto L5ab5
	case 8:
		goto L57a8
	case 9:
		goto L5b05
	case 10:
		goto L5c1f
	}
	g.caseElseError()
L016d:
	g.sub_status()
	g.printS("Falnak ment\x82l,persze fejjel.", nl)
	g.f162 = float32((float64(g.f162) + -3))
	g.s166 = "fal"
	g.gosub = append(g.gosub, 23)
	goto L5633
R019c:
	g.sub_7a66()
	g.f14a = float32((float64(g.f14a) - float64(g.f16a)))
	g.f146 = float32((float64(g.f146) - float64(g.f16e)))
	goto L0047
L01c6:
	g.sub_status()
	g.printS("Meg\x94lelt\x82l egy oszlopot.", nl)
	g.f162 = float32((float64(g.f162) + -5))
	g.gosub = append(g.gosub, 22)
	goto L5633
R01e9:
	g.f14a = float32((float64(g.f14a) - float64(g.f16a)))
	g.f146 = float32((float64(g.f146) - float64(g.f16e)))
	goto L0047
L020d:
	g.sub_status()
	g.color(7, 0)
	g.printS("Egy ujs\xa0g- \x82s jegy\xa0rus b\xa2d\x82j\xa0n\xa0l vagy.", nl)
L0233:
	g.s13e = g.inputS(1)
	g.sub_status()
	g.f172 = float32(math.Floor((g.rnd() * 5)))
	g.color(15)
	if 1 == float64(g.f142) {
		goto L0294
	}
	goto L02a0
L0294:
	g.f172 = float32(1)
L02a0:
	g.f176 = float32(float64(g.f172))
	if float64(g.f176) == 0 {
		goto L02c1
	}
	goto L02ce
L02c1:
	g.printS("-Esti H\xa1rlap! Esti H\xa1rlap!", nl)
	goto L0326
L02ce:
	if float64(g.f176) == 1 {
		goto L02e3
	}
	goto L02fc
L02e3:
	g.printS("-Takarodj innen,te szem\x82t cs\x94ves!", nl)
	g.f142 = float32(1)
	goto L0326
L02fc:
	if float64(g.f176) == 2 {
		goto L0311
	}
	goto L031d
L0311:
	g.printS("-Mit adhatok az urnak?", nl)
	goto L0326
L031d:
	g.printS("-Megjelent az Esti H\xa1rlap!", nl)
L0326:
	g.sub_6a7f()
	g.f17e = float32(float64(g.f17a))
	if float64(g.f17e) != 1 {
		goto L034e
	}
	goto L038f
L034e:
	if float64(g.f17e) != 2 {
		goto L0364
	}
	goto L038f
L0364:
	if float64(g.f17e) != 3 {
		goto L037a
	}
	goto L038f
L037a:
	if float64(g.f17e) == 4 {
		goto L038f
	}
	goto L0396
L038f:
	g.gosub = append(g.gosub, 7)
	goto L56b6
R0392:
	goto L005f
L0396:
	if float64(g.f17e) == 5 {
		goto L03ac
	}
	goto L03c0
L03ac:
	g.f14a = float32((float64(g.f14a) + 1))
	goto L005f
L03c0:
	if float64(g.f17e) == 6 {
		goto L03d5
	}
	goto L0441
L03d5:
	g.sub_7fc2()
	g.sub_status()
	g.color(15, 0)
	g.printS("-Elm\x82sz \xa0m a j\xa2 any\xa0dba! ", semi)
	g.color(7)
	g.printS("-ord\xa1tja", nl)
	g.printS("", nl)
	g.printS("ut\xa0nad.", nl)
	g.f14a = float32((float64(g.f14a) + 1))
	goto L0047
L0441:
	if float64(g.f17e) != 7 {
		goto L0457
	}
	goto L046c
L0457:
	if float64(g.f17e) == 8 {
		goto L046c
	}
	goto L0496
L046c:
	g.sub_status()
	g.color(15, 0)
	g.printS("-Nem \xa0rulok,csak jegyet \x82s ujs\xa0got!", nl)
	goto L0233
L0496:
	if float64(g.f17e) == 12 {
		goto L04ab
	}
	goto L04e1
L04ab:
	g.sub_status()
	g.color(15, 0)
	g.printS("-Most mit \xa0csorogsz itt,menj a francba!", nl)
	g.f142 = float32(1)
	goto L0233
L04e1:
	if float64(g.f17e) == 14 {
		goto L04f6
	}
	goto L0568
L04f6:
	g.sub_status()
	g.color(15, 0)
	g.printS("-Itt nem szip\xa2zhatsz,te farokgyerek!", nl)
	g.color(7)
	g.printS("", nl)
	g.printS("J\xa2l seggbe rug.", nl)
	g.f162 = float32((float64(g.f162) + -2))
	g.gosub = append(g.gosub, 8)
	goto L5633
R0553:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L0233
L0568:
	if float64(g.f17e) == 15 {
		goto L057d
	}
	goto L05ef
L057d:
	g.sub_status()
	g.color(15, 0)
	g.printS("-Itt nem pi\xa0lhatsz,te h\x81lyegyerek!", nl)
	g.color(7)
	g.printS("", nl)
	g.printS("Arr\x82bb tasz\xa1t.", nl)
	g.f162 = float32((float64(g.f162) + -2))
	g.gosub = append(g.gosub, 9)
	goto L5633
R05da:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L0233
L05ef:
	if float64(g.f17e) == 16 {
		goto L0604
	}
	goto L0676
L0604:
	g.sub_status()
	g.color(15, 0)
	g.printS("-Te kis buzi!", nl)
	g.color(7)
	g.printS("", nl)
	g.printS("Gyomron v\xa0g.", nl)
	g.f162 = float32((float64(g.f162) + -5))
	g.gosub = append(g.gosub, 10)
	goto L5633
R0661:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L0233
L0676:
	if float64(g.f17e) == 17 {
		goto L068b
	}
	goto L069d
L068b:
	g.sub_status()
	g.printS("Hov\xa0?", nl)
	goto L0233
L069d:
	if float64(g.f17e) == 18 {
		goto L06b2
	}
	goto L077e
L06b2:
	g.sub_status()
	if 2 > math.Floor((g.rnd() * 8)) {
		goto L06de
	}
	goto L0711
L06de:
	g.printS("Siker\x81lt \x82szrev\x82tlen\x81l ellopnod a p\x82nz\x82t", nl)
	g.f182 = float32((math.Floor((g.rnd() * 25)) + float64(g.f182)))
	g.sub_7a66()
	goto L0233
L0711:
	g.color(15, 0)
	g.printS("-Megvagy,csirkefog\xa2!", nl)
	g.color(7)
	g.printS("", nl)
	g.printS("J\xa2l nyakon ver.", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 11)
	goto L5633
R0769:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L0233
L077e:
	g.f18a = float32(float64(g.f186))
	if float64(g.f18a) != 1 {
		goto L07a0
	}
	goto L07e2
L07a0:
	if float64(g.f18a) != 2 {
		goto L07b6
	}
	goto L07e2
L07b6:
	if float64(g.f18a) != 3 {
		goto L07cc
	}
	goto L07e2
L07cc:
	if float64(g.f18a) == 4 {
		goto L07e2
	}
	goto L08f2
L07e2:
	if 1 == float64(g.f142) {
		goto L07f7
	}
	goto L0821
L07f7:
	g.sub_status()
	g.printS("L\xa0tja,miben s\xa0ntik\xa0lsz \x82s egy vascs\x94vel lesujt.", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 12)
	goto L5633
R0819:
	g.sub_7a66()
	goto L0233
L0821:
	g.sub_status()
	if 80 < float64(g.f18e) {
		goto L083c
	}
	goto L0885
L083c:
	g.color(7, 0)
	g.printS("Tul lassu vagy,j\xa2l orrba ver.", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 13)
	goto L5633
R0871:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L0233
L0885:
	g.color(7, 0)
	g.f162 = float32((float64(g.f162) + -5))
	g.gosub = append(g.gosub, 14)
	goto L5633
R08b1:
	g.printS("Leverted a fiut.A p\x82nz\x82t zsebre v\xa0god.", nl)
	g.f182 = float32(((math.Floor((g.rnd() * 15)) + float64(g.f182)) + 10))
	g.sub_7a66()
	g.sub_8cdf()
	goto L0047
L08f2:
	if float64(g.f18a) != 5 {
		goto L0908
	}
	goto L094a
L0908:
	if float64(g.f18a) != 6 {
		goto L091e
	}
	goto L094a
L091e:
	if float64(g.f18a) != 7 {
		goto L0934
	}
	goto L094a
L0934:
	if float64(g.f18a) == 8 {
		goto L094a
	}
	goto L0a46
L094a:
	if 1 == float64(g.f142) {
		goto L095f
	}
	goto L0989
L095f:
	g.sub_status()
	g.printS("L\xa0tja,miben s\xa0ntik\xa0lsz \x82s egy vascs\x94vel lesujt.", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 15)
	goto L5633
R0981:
	g.sub_7a66()
	goto L0233
L0989:
	g.sub_status()
	if 40 < float64(g.f18e) {
		goto L09a4
	}
	goto L09ed
L09a4:
	g.color(7, 0)
	g.printS("Tul lassu vagy,j\xa2l t\x94k\x94n rug.", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 16)
	goto L5633
R09d9:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L0047
L09ed:
	g.color(7, 0)
	g.printS("Leverted a fiut.A p\x82nz\x82t zsebre v\xa0god.", nl)
	g.f182 = float32(((math.Floor((g.rnd() * 15)) + float64(g.f182)) + 10))
	g.sub_7a66()
	g.sub_8cdf()
	goto L0047
L0a46:
	if float64(g.f18a) == 9 {
		goto L0a5b
	}
	goto L0b23
L0a5b:
	g.sub_status()
	if 80 < float64(g.f18e) {
		goto L0a76
	}
	goto L0abf
L0a76:
	g.color(7, 0)
	g.printS("Lassu vagy,ugy t\x94k\x94n ver,hogy att\xa2l     koldulsz.", nl)
	g.f162 = float32((float64(g.f162) + -15))
	g.gosub = append(g.gosub, 17)
	goto L5633
R0aab:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L0233
L0abf:
	g.color(4, 0)
	g.printS("Kinyiffantottad a nyavaly\xa0st.           A p\x82nz\x82t ellopod.", nl)
	g.f182 = float32(((math.Floor((g.rnd() * 25)) + float64(g.f182)) + 20))
	g.f192 = float32(1)
	g.sub_7a66()
	g.sub_8cdf()
	goto L0047
L0b23:
	g.sub_status()
	g.color(7, 0)
	if (b2i((!(1 != float64(g.f142)))) & b2i((!(0 == float64(g.f196))))) != 0 {
		goto L0b79
	}
	goto L0b97
L0b79:
	g.color(15)
	g.printS("-Takarodj innen,te diszn\xa2!", nl)
	goto L0233
L0b97:
	g.f19a = float32(float64(g.f196))
	if float64(g.f19a) != 1 {
		goto L0bb9
	}
	goto L0be5
L0bb9:
	if float64(g.f19a) != 2 {
		goto L0bcf
	}
	goto L0be5
L0bcf:
	if float64(g.f19a) == 3 {
		goto L0be5
	}
	goto L0cdc
L0be5:
	if 2 < math.Floor((g.rnd() * 4)) {
		goto L0c0b
	}
	goto L0c96
L0c0b:
	g.f172 = float32((math.Floor((g.rnd() * 10)) + 1))
	g.f182 = float32((float64(g.f182) + float64(g.f172)))
	g.color(15)
	g.printS("-Nesze itt van", semi)
	g.printF(float64(g.f172), semi)
	g.printS(" Ft.", nl)
	g.f162 = float32((float64(g.f162) + -5))
	g.gosub = append(g.gosub, 18)
	goto L5633
R0c82:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L0233
L0c96:
	g.color(15)
	g.printS("-Na ne r\x94h\x94gtess\x82l m\xa0r,haver!", nl)
	g.f162 = float32((float64(g.f162) + -5))
	g.gosub = append(g.gosub, 19)
	goto L5633
R0cc4:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L0233
L0cdc:
	if float64(g.f19a) != 4 {
		goto L0cf2
	}
	goto L0d5f
L0cf2:
	if float64(g.f19a) != 5 {
		goto L0d08
	}
	goto L0d5f
L0d08:
	if float64(g.f19a) != 6 {
		goto L0d1e
	}
	goto L0d5f
L0d1e:
	if float64(g.f19a) != 7 {
		goto L0d34
	}
	goto L0d5f
L0d34:
	if float64(g.f19a) != 8 {
		goto L0d4a
	}
	goto L0d5f
L0d4a:
	if float64(g.f19a) == 9 {
		goto L0d5f
	}
	goto L0e30
L0d5f:
	g.color(15)
	g.f172 = float32(math.Floor((g.rnd() * 3)))
	if 1 == float64(g.f172) {
		goto L0da3
	}
	goto L0dad
L0da3:
	g.printS("-A kurva any\xa0ddal sz\xa2rakozz\xa0l!", nl)
L0dad:
	if 2 == float64(g.f172) {
		goto L0dc2
	}
	goto L0dcc
L0dc2:
	g.printS("-El nem takarodsz,te faszkalap?!?", nl)
L0dcc:
	if 0 == float64(g.f172) {
		goto L0de1
	}
	goto L0dea
L0de1:
	g.printS("T\x81nj el,te fat\x94k\x81!", nl)
L0dea:
	g.color(7, 0)
	g.printS("", nl)
	g.printS("Fejbever egy cs\x94vel.", nl)
	g.f162 = float32((float64(g.f162) + -15))
	g.gosub = append(g.gosub, 20)
	goto L5633
R0e28:
	g.sub_7a66()
	goto L0047
L0e30:
	g.sub_status()
	g.f1a2 = float32(float64(g.f19e))
	if float64(g.f1a2) == 1 {
		goto L0e58
	}
	goto L0ecb
L0e58:
	g.f142 = float32(1)
	g.color(15, 0)
	g.printS("-Any\xa0d pics\xa0j\xa0t,te gan\x82j!", nl)
	g.printS("Mindig ut\xa0ltam az ilyen kis geciket.", nl)
	g.color(7, 0)
	g.printS("Fejbe ver egy Esti H\xa1rlappal.", nl)
	g.f162 = float32((float64(g.f162) + -5))
	g.gosub = append(g.gosub, 21)
	goto L5633
R0ec2:
	g.sub_7a66()
	goto L0233
L0ecb:
	if float64(g.f1a2) == 2 {
		goto L0ee0
	}
	goto L0f3e
L0ee0:
	g.color(7, 0)
	g.printS("D\x81h\x82ben sz\xa2hoz sem jut,majd", nl)
	g.printS("elszalad rend\x94r\x82rt.Te k\x94zben a p\x82nzt    elrakod.", nl)
	g.sub_8cdf()
	g.f182 = float32(((math.Floor((g.rnd() * 5)) + float64(g.f182)) + 5))
	g.sub_7a66()
	goto L0047
L0f3e:
	goto L0047
L0f41:
	g.sub_status()
	if (b2i((!(9 <= float64(g.f1a6)))) | b2i((!(19 >= float64(g.f1a6))))) != 0 {
		goto L0f7f
	}
	goto L0fad
L0f7f:
	g.printS("A Sk\xa0la-Metr\xa2 z\xa0rt ajtaja el\x94tt \xa0llsz.", nl)
	g.f14a = float32((float64(g.f14a) - float64(g.f16a)))
	g.f146 = float32((float64(g.f146) - float64(g.f16e)))
	goto L0fb6
L0fad:
	g.printS("Benn vagy a Sk\xa0l\xa0ban.", nl)
L0fb6:
	g.sub_6a7f()
	g.f1aa = float32(float64(g.f17a))
	if float64(g.f1aa) != 1 {
		goto L0fde
	}
	goto L101f
L0fde:
	if float64(g.f1aa) != 2 {
		goto L0ff4
	}
	goto L101f
L0ff4:
	if float64(g.f1aa) != 3 {
		goto L100a
	}
	goto L101f
L100a:
	if float64(g.f1aa) == 4 {
		goto L101f
	}
	goto L1062
L101f:
	g.gosub = append(g.gosub, 3)
	goto L56b6
R1022:
	goto L005f
L1062:
	if float64(g.f1aa) != 5 {
		goto L1078
	}
	goto L108d
L1078:
	if float64(g.f1aa) == 6 {
		goto L108d
	}
	goto L1106
L108d:
	g.sub_status()
	if (b2i((!(9 <= float64(g.f1a6)))) | b2i((!(19 >= float64(g.f1a6))))) != 0 {
		goto L10cb
	}
	goto L10e8
L10cb:
	g.printS("Hov\xa0 futsz,te szerencs\x82tlen?!?", nl)
	g.f14a = float32((float64(g.f14a) + 1))
	goto L1102
L10e8:
	g.printS("Megbotr\xa0nkozva n\x82znek r\xa0d.", nl)
	g.f14a = float32((float64(g.f14a) + 1))
L1102:
	goto L0047
L1106:
	if float64(g.f1aa) == 7 {
		goto L111b
	}
	goto L1278
L111b:
	g.sub_status()
	if (b2i((!(9 <= float64(g.f1a6)))) | b2i((!(19 >= float64(g.f1a6))))) != 0 {
		goto L1159
	}
	goto L1166
L1159:
	g.printS("Nem tudsz.", nl)
	goto L1274
L1166:
	if 30 > float64(g.f182) {
		goto L117b
	}
	goto L1188
L117b:
	g.printS("Nincs r\xa0 p\x82nzed.", nl)
	goto L0fb6
L1188:
	g.f182 = float32((float64(g.f182) + -30))
	g.sub_7a66()
	g.sub_status()
	g.printS("Vett\x82l egy \x81veg ", semi)
	if 1 == float64(g.f1ae) {
		goto L11c1
	}
	goto L11cb
L11c1:
	g.printS("csukam\xa0jolajat.", nl)
L11cb:
	if 2 == float64(g.f1ae) {
		goto L11e0
	}
	goto L11ea
L11e0:
	g.printS("kis\x81stit.", nl)
L11ea:
	if 3 == float64(g.f1ae) {
		goto L11ff
	}
	goto L1209
L11ff:
	g.printS("vodk\xa0t.", nl)
L1209:
	g.f1b2 = float32((float64(g.f1b2) + 1))
	if 1 < float64(g.f1b2) {
		goto L1230
	}
	goto L126f
L1230:
	g.f1b2 = float32(1)
	g.color(5)
	g.printS("", nl)
	g.printS("Elejtetted az \x81veget,te \x81gyetlen!", nl)
	g.color(7)
L126f:
	g.sub_7a66()
L1274:
	goto L0fb6
L1278:
	if float64(g.f1aa) == 8 {
		goto L128e
	}
	goto L13a7
L128e:
	if (b2i((!(9 <= float64(g.f1a6)))) | b2i((!(19 >= float64(g.f1a6))))) != 0 {
		goto L12c6
	}
	goto L12d7
L12c6:
	g.sub_status()
	g.printS("Z\xa0rva van,reggel 9-kor nyit.", nl)
	goto L139e
L12d7:
	g.sub_status()
	if 14 > float64(g.f182) {
		goto L12f2
	}
	goto L12ff
L12f2:
	g.printS("Nincs r\xa0 p\x82nzed.", nl)
	goto L0fb6
L12ff:
	g.f182 = float32((float64(g.f182) + -14))
	g.f1b6 = float32((float64(g.f1b6) + 1))
	g.sub_7a66()
	g.sub_status()
	if 3 > float64(g.f1b6) {
		goto L1340
	}
	goto L1380
L1340:
	g.color(15)
	g.printS("-Egy Technokolt, meg egy nylonzacsk\xa2t", nl)
	g.printS("k\x82rek.", semi)
	g.color(7)
	g.printS("-mondod. Sz\xa2 n\x82lk\x81l kiszolg\xa0lnak.", nl)
	goto L139e
L1380:
	g.printS("Egy elad\xa2nak felt\x81nik,hogy milyen", nl)
	g.printS("sok ragaszt\xa2d van. Rohan telefon\xa0lni.", nl)
	g.f192 = float32(1)
L139e:
	g.sub_7a66()
	goto L0fb6
L13a7:
	if float64(g.f1aa) != 14 {
		goto L13bd
	}
	goto L13d2
L13bd:
	if float64(g.f1aa) == 15 {
		goto L13d2
	}
	goto L1464
L13d2:
	g.sub_status()
	if (b2i((!(9 <= float64(g.f1a6)))) | b2i((!(19 >= float64(g.f1a6))))) != 0 {
		goto L1410
	}
	goto L142d
L1410:
	g.printS("Arra m\xa0r nincs sz\x81ks\x82ged.", nl)
	g.f14a = float32((float64(g.f14a) + 1))
	goto L1460
L142d:
	g.printS("Kihaj\xa1tanak.", nl)
	g.f162 = float32((float64(g.f162) + -5))
	g.gosub = append(g.gosub, 4)
	goto L5633
R144a:
	g.sub_7a66()
	g.f14a = float32((float64(g.f14a) + 1))
L1460:
	goto L0047
L1464:
	if float64(g.f1aa) == 17 {
		goto L1479
	}
	goto L157a
L1479:
	g.sub_status()
	if (b2i((!(8 >= float64(g.f1a6)))) & b2i((!(20 <= float64(g.f1a6))))) != 0 {
		goto L14b7
	}
	goto L14c4
L14b7:
	g.printS("Nyitva van m\x82g,te idiot!", nl)
	goto L0fb6
L14c4:
	if 1 == float64(g.f1ba) {
		goto L14d9
	}
	goto L1510
L14d9:
	g.color(9)
	g.printS("Ber\xa3god a kirakatot.", nl)
	g.printS("Egy nyomoz\xa2 ugrik el\x94 az oszlop m\x94g\x81l.", nl)
	g.f14a = float32((float64(g.f14a) + 1))
	goto L5288
L1510:
	g.printS("Sikeresen v\x82grehajtottad a bet\x94r\x82st.", nl)
	g.printS("(Az ujjlenyomatod nem t\x94r\x94lted le.)", nl)
	g.f182 = float32(((math.Floor((g.rnd() * 100)) + float64(g.f182)) + 20))
	g.f192 = float32(1)
	g.f1ba = float32(1)
	g.sub_7a66()
	g.f14a = float32((float64(g.f14a) + 1))
	goto L0047
L157a:
	g.sub_status()
	g.printS("Arra m\xa0r nincs sz\x81ks\x82ged.", nl)
	goto L0fb6
L158b:
	g.sub_status()
	if 1 == float64(g.f16a) {
		goto L15a6
	}
	goto L15c3
L15a6:
	g.printS("Levitt a mozg\xa2l\x82pcs\x94.", nl)
	g.f14a = float32((float64(g.f14a) + 1))
	goto L0047
L15c3:
	g.printS("Felvitt a mozg\xa2l\x82pcs\x94.", nl)
	g.f14a = float32((float64(g.f14a) + -1))
	goto L0047
R168d:
	goto L0047
R178e:
	g.sub_7a66()
	goto L0047
L18f2:
	g.color(7, 0)
	g.f1c2 = float32((float64(g.f1c2) + 1))
	if 5 == float64(g.f1c2) {
		goto L1932
	}
	goto L194f
L1932:
	g.f1c2 = float32(0)
	g.f1a6 = float32((float64(g.f1a6) + 1))
L194f:
	g.f162 = float32((float64(g.f162) + -1))
	g.gosub = append(g.gosub, 24)
	goto L5633
R1962:
	g.sub_7a66()
	g.sub_status()
	g.f172 = float32((math.Floor((g.rnd() * 38)) + 1))
	if 17 < float64(g.f172) {
		goto L19a4
	}
	goto L1bf2
L19a4:
	g.printS("Itt nincs semmi,mehetsz tov\xa0bb.", nl)
	g.sub_6a7f()
	g.f1c6 = float32(float64(g.f17a))
	if float64(g.f1c6) != 1 {
		goto L19d5
	}
	goto L1a16
L19d5:
	if float64(g.f1c6) != 2 {
		goto L19eb
	}
	goto L1a16
L19eb:
	if float64(g.f1c6) != 3 {
		goto L1a01
	}
	goto L1a16
L1a01:
	if float64(g.f1c6) == 4 {
		goto L1a16
	}
	goto L1a1d
L1a16:
	g.gosub = append(g.gosub, 25)
	goto L56b6
R1a19:
	goto L005f
L1a1d:
	if float64(g.f1c6) == 12 {
		goto L1a32
	}
	goto L1a94
L1a32:
	g.sub_status()
	g.color(7, 0)
	g.printS("Kipihented magad.", nl)
	g.f162 = float32((float64(g.f162) + 3))
	if 100 < float64(g.f162) {
		goto L1a80
	}
	goto L1a8b
L1a80:
	g.f162 = float32(100)
L1a8b:
	g.sub_7a66()
	goto L0047
L1a94:
	if float64(g.f1c6) == 15 {
		goto L1aa9
	}
	goto L1b2a
L1aa9:
	g.sub_status()
	if 0 < float64(g.f1b2) {
		goto L1ac5
	}
	goto L1b1a
L1ac5:
	g.f1b2 = float32((float64(g.f1b2) + -1))
	g.printS("Szlop\xa0lt\xa0l egy j\xa2t,m\xa0ris jobban vagy.", nl)
	if 0 == float64(g.f1b2) {
		goto L1af4
	}
	goto L1b07
L1af4:
	g.printS("", nl)
	g.printS("(Elfogyott minden pi\xa0d.)", nl)
L1b07:
	g.f162 = float32(100)
	g.sub_7a66()
	goto L0047
L1b1a:
	g.printS("Nincs pi\xa0d.Ideje,hogy szerezz.", nl)
	goto L0047
L1b2a:
	if float64(g.f1c6) == 14 {
		goto L1b3f
	}
	goto L1bc8
L1b3f:
	g.sub_status()
	if 0 < float64(g.f1b6) {
		goto L1b5b
	}
	goto L1bb9
L1b5b:
	g.f1b6 = float32((float64(g.f1b6) + -1))
	g.printS("Szip\xa2zt\xa0l egy j\xa2t,m\xa0ris olyan gyors", nl)
	g.printS("vagy,mint a vill\xa0m.", nl)
	if 0 == float64(g.f1b6) {
		goto L1b93
	}
	goto L1ba6
L1b93:
	g.printS("", nl)
	g.printS("(Elfogyott minden ragaszt\xa2d.)", nl)
L1ba6:
	g.f18e = float32(10)
	g.sub_7a66()
	goto L0047
L1bb9:
	g.printS("Ki\x81r\x81lt a tubus.", nl)
	goto L0047
L1bc8:
	g.sub_status()
	g.color(7, 0)
	g.printS("Arra m\xa0r nincs sz\x81ks\x82ged.", nl)
	goto L0047
L1bf2:
	if (b2i((!(25 >= math.Floor((g.rnd() * 30))))) & b2i((!(30 >= float64(g.f18e))))) == 0 {
		goto L1c3c
	}
	goto L5c41
L1c3c:
	switch cint(float64(g.f172)) {
	case 1:
		goto L1c70
	case 2:
		goto L1c70
	case 3:
		goto L21b8
	case 4:
		goto L21b8
	case 5:
		goto L2815
	case 6:
		goto L2815
	case 7:
		goto L2f17
	case 8:
		goto L2f17
	case 9:
		goto L3499
	case 10:
		goto L3499
	case 11:
		goto L3c35
	case 12:
		goto L3c35
	case 13:
		goto L42d1
	case 14:
		goto L42d1
	case 15:
		goto L4952
	case 16:
		goto L4952
	case 17:
		goto L521b
	}
L1c70:
	if (b2i((!(6 <= float64(g.f1a6)))) | b2i((!(20 >= float64(g.f1a6))))) == 0 {
		goto L1ca8
	}
	goto L18f2
L1ca8:
	g.printS("Egy karcsiszem\x81veges gyerek,a h\xa0t\xa0n     iskolat\xa0sk\xa0val,riadtan n\x82z r\xa0d.", nl)
	g.sub_6a7f()
	g.f1ca = float32(float64(g.f17a))
	if float64(g.f1ca) != 1 {
		goto L1cd9
	}
	goto L1d1a
L1cd9:
	if float64(g.f1ca) != 2 {
		goto L1cef
	}
	goto L1d1a
L1cef:
	if float64(g.f1ca) != 3 {
		goto L1d05
	}
	goto L1d1a
L1d05:
	if float64(g.f1ca) == 4 {
		goto L1d1a
	}
	goto L1d21
L1d1a:
	g.gosub = append(g.gosub, 88)
	goto L56b6
R1d1d:
	goto L005f
L1d21:
	if float64(g.f1ca) != 5 {
		goto L1d37
	}
	goto L1d4d
L1d37:
	if float64(g.f1ca) == 6 {
		goto L1d4d
	}
	goto L1d79
L1d4d:
	if 6 == float64(g.f17a) {
		goto L1d62
	}
	goto L1d67
L1d62:
	g.sub_7fc2()
L1d67:
	g.sub_status()
	g.printS("S\x82rt\x94d\x94tt pof\xa0val n\x82z ut\xa0nad.", nl)
	goto L0047
L1d79:
	if float64(g.f1ca) != 7 {
		goto L1d8f
	}
	goto L1da4
L1d8f:
	if float64(g.f1ca) == 8 {
		goto L1da4
	}
	goto L1de1
L1da4:
	g.sub_status()
	g.color(15)
	g.printS("-Nincs n\xa0lam!", nl)
	g.color(7)
	g.printS("B\x94gve elrohan.", nl)
	goto L0047
L1de1:
	if float64(g.f1ca) == 9 {
		goto L1df6
	}
	goto L1e81
L1df6:
	g.sub_status()
	g.f172 = float32((math.Floor((g.rnd() * 9)) + 1))
	g.color(15)
	g.printS("-Csak", semi)
	g.printF(float64(g.f172), semi)
	g.printS("forintom van.", semi)
	g.color(7)
	g.printS("-mondja b\x94gve.", nl)
	g.f182 = float32((float64(g.f182) + float64(g.f172)))
	g.sub_7a66()
	goto L0047
L1e81:
	if float64(g.f1ca) == 10 {
		goto L1e96
	}
	goto L1f79
L1e96:
	g.f172 = float32(math.Floor((g.rnd() * 3)))
	g.sub_status()
	if 0 == float64(g.f172) {
		goto L1ece
	}
	goto L1ef4
L1ece:
	g.printS("El\x94pattan a mam\xa0ja \x82s olyan nyaklevest  kever le,hogy a f\x81led k\x82t fele \xa0ll.", nl)
	g.f162 = float32((float64(g.f162) + -5))
	g.gosub = append(g.gosub, 89)
	goto L5633
R1eeb:
	g.sub_7a66()
	goto L0047
L1ef4:
	if 1 == float64(g.f172) {
		goto L1f09
	}
	goto L1f40
L1f09:
	g.color(15)
	g.printS("-Beeee...!", semi)
	g.color(7)
	g.printS("-mondja \x82s elszalad.", nl)
	goto L0047
L1f40:
	g.printS("Riadtan b\xa0mul,te pedig kifosztod.", nl)
	g.f182 = float32(((math.Floor((g.rnd() * 9)) + float64(g.f182)) + 1))
	g.sub_7a66()
	goto L0047
L1f79:
	if float64(g.f1ca) == 11 {
		goto L1f8e
	}
	goto L1fb4
L1f8e:
	g.sub_status()
	g.printS("A felh\xa0borodott j\xa0r\xa2kel\x94k meglincselnek.", nl)
	g.f162 = float32((float64(g.f162) + -50))
	g.gosub = append(g.gosub, 90)
	goto L5633
R1fb0:
	goto L21b4
L1fb4:
	if float64(g.f1ca) != 12 {
		goto L1fca
	}
	goto L1ff5
L1fca:
	if float64(g.f1ca) != 14 {
		goto L1fe0
	}
	goto L1ff5
L1fe0:
	if float64(g.f1ca) == 15 {
		goto L1ff5
	}
	goto L2007
L1ff5:
	g.sub_status()
	g.printS("Elt\x81nik a szemed el\x94l.", nl)
	goto L0047
L2007:
	if float64(g.f1ca) == 16 {
		goto L201c
	}
	goto L209e
L201c:
	g.sub_status()
	g.printS("Elkezd sivalkodni,", semi)
	if (b2i((!(6 <= float64(g.f1a6)))) | b2i((!(22 >= float64(g.f1a6))))) != 0 {
		goto L2063
	}
	goto L206f
L2063:
	g.printS("de az alulj\xa0r\xa2 kihalt.", nl)
	goto L0047
L206f:
	g.printS("", nl)
	g.printS("a felh\xa0borodott emberek le\x81tnek.", nl)
	g.f162 = float32((float64(g.f162) + -50))
	g.gosub = append(g.gosub, 91)
	goto L5633
R2095:
	g.sub_7a66()
	goto L0047
L209e:
	if float64(g.f1ca) == 17 {
		goto L20b3
	}
	goto L20c5
L20b3:
	g.sub_status()
	g.printS("Hov\xa0?", nl)
	goto L0047
L20c5:
	if float64(g.f1ca) == 18 {
		goto L20da
	}
	goto L21b4
L20da:
	g.sub_status()
	g.f172 = float32(math.Floor((g.rnd() * 3)))
	if 0 == float64(g.f172) {
		goto L2112
	}
	goto L211f
L2112:
	g.printS("Nincs t\xa0rc\xa0ja.", nl)
	goto L0047
L211f:
	if 1 == float64(g.f172) {
		goto L2134
	}
	goto L217c
L2134:
	g.color(15)
	g.printS("Seg\xa1ts\x82g!Tolvaj!", nl)
	g.color(7)
	g.printS("Rohan rend\x94r\x82rt.", nl)
	g.f192 = float32(1)
	g.sub_7a66()
	goto L0047
L217c:
	g.printS("Elloptad \x82szrev\x82tlen\x81l.", nl)
	g.f182 = float32(((math.Floor((g.rnd() * 9)) + float64(g.f182)) + 1))
	g.sub_7a66()
	goto L0047
L21b4:
	goto L0047
L21b8:
	if (b2i((!(6 <= float64(g.f1a6)))) | b2i((!(22 >= float64(g.f1a6))))) == 0 {
		goto L21f0
	}
	goto L18f2
L21f0:
	g.sub_status()
	g.printS("Egy \x94reg fasz megy el el\x94tted.", nl)
L21fe:
	g.sub_6a7f()
	g.f1ce = float32(float64(g.f17a))
	if float64(g.f1ce) != 1 {
		goto L2226
	}
	goto L2267
L2226:
	if float64(g.f1ce) != 2 {
		goto L223c
	}
	goto L2267
L223c:
	if float64(g.f1ce) != 3 {
		goto L2252
	}
	goto L2267
L2252:
	if float64(g.f1ce) == 4 {
		goto L2267
	}
	goto L226e
L2267:
	g.gosub = append(g.gosub, 83)
	goto L56b6
R226a:
	goto L005f
L226e:
	if float64(g.f1ce) != 5 {
		goto L2284
	}
	goto L229a
L2284:
	if float64(g.f1ce) == 6 {
		goto L229a
	}
	goto L22c6
L229a:
	if 6 == float64(g.f17a) {
		goto L22af
	}
	goto L22b4
L22af:
	g.sub_7fc2()
L22b4:
	g.sub_status()
	g.printS("Betelt a nadr\xa0god?", nl)
	goto L0047
L22c6:
	if float64(g.f1ce) != 7 {
		goto L22dc
	}
	goto L22f1
L22dc:
	if float64(g.f1ce) == 8 {
		goto L22f1
	}
	goto L232e
L22f1:
	g.sub_status()
	g.color(15)
	g.printS("-Nem vagyok boltos!", semi)
	g.color(7)
	g.printS("-morogva eltotyog.", nl)
	goto L0047
L232e:
	if float64(g.f1ce) == 9 {
		goto L2343
	}
	goto L245c
L2343:
	g.sub_status()
	g.f172 = float32((math.Floor((g.rnd() * 9)) + 1))
	g.color(15)
	if ((b2i((!(1 != float64(g.f196)))) | b2i((!(3 != float64(g.f196))))) & b2i((!(0 != float64(g.f142))))) != 0 {
		goto L23cf
	}
	goto L2410
L23cf:
	g.printS("-L\xa0tom,te rendes gyerek vagy.", nl)
	g.printS("Itt van", semi)
	g.printF(float64(g.f172), semi)
	g.printS("forint.", nl)
	g.f182 = float32((float64(g.f182) + float64(g.f172)))
	g.sub_7a66()
	goto L0047
L2410:
	g.printS("-Szemtelen hulig\xa0n!", nl)
	g.color(7)
	g.printS("Fejbe \x81t a kamp\xa2sbottal.", nl)
	g.f162 = float32((float64(g.f162) + -5))
	g.gosub = append(g.gosub, 84)
	goto L5633
R2447:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L21fe
L245c:
	if float64(g.f1ce) == 10 {
		goto L2471
	}
	goto L2534
L2471:
	g.f172 = float32(math.Floor((g.rnd() * 2)))
	g.sub_status()
	if 0 == float64(g.f172) {
		goto L24a9
	}
	goto L24d8
L24a9:
	g.color(15)
	g.printS("-Tess\x82k? Rossz a hall\xa0som,hangosabban   mondja!", nl)
	g.color(7)
	goto L21fe
L24d8:
	if 1 == float64(g.f172) {
		goto L24ed
	}
	goto L2530
L24ed:
	g.color(15)
	g.printS("-\x99regany\xa0d r\x82zzel kivert vasporos       pics\xa0j\xa0t!", semi)
	g.color(7)
	g.printS("-mondja.", nl)
	g.f142 = float32(1)
	goto L21fe
L2530:
	goto L2810
L2534:
	if float64(g.f1ce) == 11 {
		goto L2549
	}
	goto L262a
L2549:
	g.sub_status()
	if 9 == float64(g.f186) {
		goto L2564
	}
	goto L25a4
L2564:
	g.color(4)
	g.printS("Led\x94fted az \x94reg buzit.", nl)
	g.f192 = float32(1)
	g.sub_7a66()
	g.color(7)
	goto L0047
L25a4:
	if (b2i((!(7 >= float64(g.f1a6)))) & b2i((!(20 <= float64(g.f1a6))))) != 0 {
		goto L25dc
	}
	goto L2601
L25dc:
	g.printS("A felh\xa0borodott j\xa0r\xa2kel\x94k megvernek.", nl)
	g.f162 = float32((float64(g.f162) + -30))
	g.gosub = append(g.gosub, 85)
	goto L5633
R25f9:
	g.sub_7a66()
	goto L0047
L2601:
	g.printS("J\xa2l megverted a nyavaly\xa0st.", nl)
	g.f162 = float32((float64(g.f162) + -5))
	g.gosub = append(g.gosub, 86)
	goto L5633
R261e:
	g.sub_7a66()
	goto L0047
L262a:
	if float64(g.f1ce) != 12 {
		goto L2640
	}
	goto L266b
L2640:
	if float64(g.f1ce) != 14 {
		goto L2656
	}
	goto L266b
L2656:
	if float64(g.f1ce) == 15 {
		goto L266b
	}
	goto L267d
L266b:
	g.sub_status()
	g.printS("Elt\x81nik.", nl)
	goto L0047
L267d:
	if float64(g.f1ce) == 16 {
		goto L2692
	}
	goto L2714
L2692:
	g.sub_status()
	g.printS("Elkezd ord\xa1tozni,", semi)
	if (b2i((!(7 <= float64(g.f1a6)))) | b2i((!(20 >= float64(g.f1a6))))) != 0 {
		goto L26d9
	}
	goto L26e5
L26d9:
	g.printS("de az alulj\xa0r\xa2 kihalt.", nl)
	goto L0047
L26e5:
	g.printS("", nl)
	g.printS("a felh\xa0borodott emberek le\x81tnek.", nl)
	g.f162 = float32((float64(g.f162) + -45))
	g.gosub = append(g.gosub, 87)
	goto L5633
R270b:
	g.sub_7a66()
	goto L0047
L2714:
	if float64(g.f1ce) == 17 {
		goto L2729
	}
	goto L273b
L2729:
	g.sub_status()
	g.printS("Hov\xa0?", nl)
	goto L0047
L273b:
	if float64(g.f1ce) == 18 {
		goto L2750
	}
	goto L2810
L2750:
	g.sub_status()
	g.f172 = float32(math.Floor((g.rnd() * 3)))
	if 0 == float64(g.f172) {
		goto L2788
	}
	goto L2795
L2788:
	g.printS("Nincs t\xa0rc\xa0ja.", nl)
	goto L0047
L2795:
	if 1 == float64(g.f172) {
		goto L27aa
	}
	goto L27d8
L27aa:
	g.color(15)
	g.printS("Seg\xa1ts\x82g!Tolvaj!Rend\x94rt!", nl)
	g.f192 = float32(1)
	g.sub_7a66()
	goto L0047
L27d8:
	g.printS("Elloptad \x82szrev\x82tlen\x81l.", nl)
	g.f182 = float32(((math.Floor((g.rnd() * 9)) + float64(g.f182)) + 1))
	g.sub_7a66()
	goto L0047
L2810:
	g.caseElseError()
L2815:
	g.sub_status()
	if (b2i((!(6 <= float64(g.f1a6)))) | b2i((!(21 >= float64(g.f1a6))))) == 0 {
		goto L2853
	}
	goto L18f2
L2853:
	g.printS("Egy v\x82n banya megy el el\x94tted.", nl)
L285c:
	g.sub_6a7f()
	g.f1d2 = float32(float64(g.f17a))
	if float64(g.f1d2) != 1 {
		goto L2884
	}
	goto L28c5
L2884:
	if float64(g.f1d2) != 2 {
		goto L289a
	}
	goto L28c5
L289a:
	if float64(g.f1d2) != 3 {
		goto L28b0
	}
	goto L28c5
L28b0:
	if float64(g.f1d2) == 4 {
		goto L28c5
	}
	goto L28cc
L28c5:
	g.gosub = append(g.gosub, 78)
	goto L56b6
R28c8:
	goto L005f
L28cc:
	if float64(g.f1d2) != 5 {
		goto L28e2
	}
	goto L28f8
L28e2:
	if float64(g.f1d2) == 6 {
		goto L28f8
	}
	goto L2950
L28f8:
	if 6 == float64(g.f17a) {
		goto L290d
	}
	goto L293b
L290d:
	g.sub_7fc2()
	g.sub_status()
	g.color(15, 0)
	g.printS("-Ezt tanulj\xa0tok az iskol\xa0ban?!?", nl)
	goto L0047
L293b:
	g.sub_status()
	g.printS("Betelt a nadr\xa0god?", nl)
	goto L0047
L2950:
	if float64(g.f1d2) != 7 {
		goto L2966
	}
	goto L297b
L2966:
	if float64(g.f1d2) == 8 {
		goto L297b
	}
	goto L29b8
L297b:
	g.sub_status()
	g.color(15)
	g.printS("-Ezek a mai fiatalok!", semi)
	g.color(7)
	g.printS("-mammogva eltotyog.", nl)
	goto L0047
L29b8:
	if float64(g.f1d2) == 9 {
		goto L29cd
	}
	goto L2b14
L29cd:
	g.sub_status()
	g.f172 = float32((math.Floor((g.rnd() * 15)) + 1))
	g.color(15)
	l0c = ((b2i((!(1 != float64(g.f196)))) | b2i((!(4 != float64(g.f196))))) & b2i((!(0 != float64(g.f142)))))
	if (l0c & b2i((!(1 >= math.Floor((g.rnd() * 3)))))) != 0 {
		goto L2a87
	}
	goto L2ac8
L2a87:
	g.printS("-L\xa0tom,te rendes kisfiu vagy.", nl)
	g.printS("Tess\x82k", semi)
	g.printF(float64(g.f172), semi)
	g.printS("forint.", nl)
	g.f182 = float32((float64(g.f182) + float64(g.f172)))
	g.sub_7a66()
	goto L0047
L2ac8:
	g.printS("-Szemtelen hulig\xa0n!", nl)
	g.color(7)
	g.printS("Fejbe \x81t az eserny\x94j\x82vel.", nl)
	g.f162 = float32((float64(g.f162) + -8))
	g.gosub = append(g.gosub, 79)
	goto L5633
R2aff:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L285c
L2b14:
	if float64(g.f1d2) == 10 {
		goto L2b29
	}
	goto L2bec
L2b29:
	g.f172 = float32(math.Floor((g.rnd() * 2)))
	g.sub_status()
	if 0 == float64(g.f172) {
		goto L2b61
	}
	goto L2b90
L2b61:
	g.color(15)
	g.printS("-Tess\x82k? Kicsit nagyothallok,aranyosk\xa0m.", nl)
	g.color(7)
	goto L285c
L2b90:
	if 1 == float64(g.f172) {
		goto L2ba5
	}
	goto L2be8
L2ba5:
	g.color(15)
	g.printS("-Az ap\xa0dat,fiam!", semi)
	g.color(7)
	g.printS("-mondja.", nl)
	g.f142 = float32(1)
	goto L285c
L2be8:
	goto L2f12
L2bec:
	if float64(g.f1d2) == 11 {
		goto L2c01
	}
	goto L2c81
L2c01:
	g.sub_status()
	if 9 == float64(g.f186) {
		goto L2c1c
	}
	goto L2c5b
L2c1c:
	g.color(4)
	g.printS("Led\x94fted a v\x82n kurv\xa0t.", nl)
	g.f192 = float32(1)
	g.sub_7a66()
	g.color(7)
	goto L0047
L2c5b:
	g.printS("A felh\xa0borodott j\xa0r\xa2kel\x94k megvernek.", nl)
	g.f162 = float32((float64(g.f162) + -30))
	g.gosub = append(g.gosub, 80)
	goto L5633
R2c78:
	g.sub_7a66()
	goto L0047
L2c81:
	if float64(g.f1d2) != 12 {
		goto L2c97
	}
	goto L2cc2
L2c97:
	if float64(g.f1d2) != 14 {
		goto L2cad
	}
	goto L2cc2
L2cad:
	if float64(g.f1d2) == 15 {
		goto L2cc2
	}
	goto L2cd4
L2cc2:
	g.sub_status()
	g.printS("Elt\x81nik.", nl)
	goto L0047
L2cd4:
	if float64(g.f1d2) == 16 {
		goto L2ce9
	}
	goto L2e16
L2ce9:
	g.sub_status()
	if (b2i((!(0 != math.Floor((g.rnd() * 3))))) & b2i((!(0 != float64(g.f142))))) != 0 {
		goto L2d38
	}
	goto L2d99
L2d38:
	g.color(15)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 81)
	goto L5633
R2d5d:
	g.printS("-Aaaah... Legk\x94zelebb a Gerbeaud-ban,   kedvesem.", semi)
	g.color(7)
	g.printS("-mondja \x82s ad 20 forintot.", nl)
	g.f182 = float32((float64(g.f182) + 20))
	g.sub_7a66()
	goto L0047
L2d99:
	g.printS("Elkezd sivalkodni,", semi)
	if (b2i((!(6 <= float64(g.f1a6)))) | b2i((!(22 >= float64(g.f1a6))))) != 0 {
		goto L2ddb
	}
	goto L2de7
L2ddb:
	g.printS("de az alulj\xa0r\xa2 kihalt.", nl)
	goto L0047
L2de7:
	g.printS("", nl)
	g.printS("a felh\xa0borodott emberek le\x81tnek.", nl)
	g.f162 = float32((float64(g.f162) + -45))
	g.gosub = append(g.gosub, 82)
	goto L5633
R2e0d:
	g.sub_7a66()
	goto L0047
L2e16:
	if float64(g.f1d2) == 17 {
		goto L2e2b
	}
	goto L2e3d
L2e2b:
	g.sub_status()
	g.printS("Hov\xa0?", nl)
	goto L0047
L2e3d:
	if float64(g.f1d2) == 18 {
		goto L2e52
	}
	goto L2f12
L2e52:
	g.sub_status()
	g.f172 = float32(math.Floor((g.rnd() * 3)))
	if 0 == float64(g.f172) {
		goto L2e8a
	}
	goto L2e97
L2e8a:
	g.printS("Nincs t\xa0rc\xa0ja.", nl)
	goto L0047
L2e97:
	if 1 == float64(g.f172) {
		goto L2eac
	}
	goto L2eda
L2eac:
	g.color(15)
	g.printS("Seg\xa1ts\x82g!Tolvaj!Ellopta a t\xa0rc\xa0m!", nl)
	g.f192 = float32(1)
	g.sub_7a66()
	goto L0047
L2eda:
	g.printS("Elloptad \x82szrev\x82tlen\x81l.", nl)
	g.f182 = float32(((math.Floor((g.rnd() * 9)) + float64(g.f182)) + 1))
	g.sub_7a66()
	goto L0047
L2f12:
	g.caseElseError()
L2f17:
	g.sub_status()
	if 0 == math.Floor((g.rnd() * 2)) {
		goto L2f43
	}
	goto L2f4f
L2f43:
	g.printS("Egy n\x94 j\x94n szembe veled.", nl)
	goto L2f58
L2f4f:
	g.printS("Egy helyes csajt l\xa0tnak szemeid.", nl)
L2f58:
	g.sub_6a7f()
	g.f1d6 = float32(float64(g.f17a))
	if float64(g.f1d6) != 1 {
		goto L2f80
	}
	goto L2fc1
L2f80:
	if float64(g.f1d6) != 2 {
		goto L2f96
	}
	goto L2fc1
L2f96:
	if float64(g.f1d6) != 3 {
		goto L2fac
	}
	goto L2fc1
L2fac:
	if float64(g.f1d6) == 4 {
		goto L2fc1
	}
	goto L2fd6
L2fc1:
	g.sub_status()
	g.printS("Kihagytad az alkalmat,te pancser.", nl)
	g.gosub = append(g.gosub, 72)
	goto L56b6
R2fd2:
	goto L0047
L2fd6:
	if float64(g.f1d6) != 5 {
		goto L2fec
	}
	goto L3002
L2fec:
	if float64(g.f1d6) == 6 {
		goto L3002
	}
	goto L302e
L3002:
	if 6 == float64(g.f17a) {
		goto L3017
	}
	goto L301c
L3017:
	g.sub_7fc2()
L301c:
	g.sub_status()
	g.printS("Megijedt\x82l?", nl)
	goto L0047
L302e:
	if float64(g.f1d6) != 7 {
		goto L3044
	}
	goto L3059
L3044:
	if float64(g.f1d6) == 8 {
		goto L3059
	}
	goto L3077
L3059:
	g.sub_status()
	g.printS("Lefitym\xa0l\xa2an n\x82z r\xa0d,de nem v\xa0laszol.", nl)
	g.f142 = float32(1)
	goto L2f58
L3077:
	if float64(g.f1d6) == 9 {
		goto L308c
	}
	goto L30bb
L308c:
	g.sub_status()
	g.color(15)
	g.printS("-Menj haza anyuk\xa0dhoz,bunk\xa2!", nl)
	g.f142 = float32(1)
	goto L2f58
L30bb:
	if float64(g.f1d6) == 10 {
		goto L30d0
	}
	goto L31b5
L30d0:
	g.f172 = float32(math.Floor((g.rnd() * 2)))
	g.sub_status()
	if 1 == float64(g.f172) {
		goto L3108
	}
	goto L3159
L3108:
	g.color(15)
	g.printS("-B\x82la! verd meg ezt a kis buzit!", nl)
	g.color(7)
	g.printS("El\x94ugrik a fiuja \x82s lebasz neked egy    nagypofont,ut\xa0na elmennek.", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 73)
	goto L5633
R3150:
	g.sub_7a66()
	goto L0047
L3159:
	if 0 == float64(g.f172) {
		goto L316e
	}
	goto L31b1
L316e:
	g.color(15)
	g.printS("-Anyuk\xa0ddal j\xa0tssz\xa0l,te kis             c\x82rnafaszu!", semi)
	g.color(7)
	g.printS("-mondja.", nl)
	g.f142 = float32(1)
	goto L2f58
L31b1:
	goto L3494
L31b5:
	if float64(g.f1d6) == 11 {
		goto L31ca
	}
	goto L3280
L31ca:
	g.sub_status()
	if 9 == float64(g.f186) {
		goto L31e5
	}
	goto L3239
L31e5:
	g.color(4)
	g.printS("Leszurtad a nyavaly\xa0s ribancot.", nl)
	g.f162 = float32((float64(g.f162) + -5))
	g.gosub = append(g.gosub, 74)
	goto L5633
R3214:
	g.f192 = float32(1)
	g.sub_7a66()
	g.color(7)
	goto L0047
L3239:
	if 1 == float64(g.f186) {
		goto L324e
	}
	goto L325a
L324e:
	g.printS("Nincs neki.", nl)
	goto L2f58
L325a:
	g.printS("A felh\xa0borodott j\xa0r\xa2kel\x94k megvernek.", nl)
	g.f162 = float32((float64(g.f162) + -30))
	g.gosub = append(g.gosub, 75)
	goto L5633
R3277:
	g.sub_7a66()
	goto L0047
L3280:
	if float64(g.f1d6) != 12 {
		goto L3296
	}
	goto L32c1
L3296:
	if float64(g.f1d6) != 14 {
		goto L32ac
	}
	goto L32c1
L32ac:
	if float64(g.f1d6) == 15 {
		goto L32c1
	}
	goto L32d3
L32c1:
	g.sub_status()
	g.printS("Segg\x82t risz\xa0lva elt\x81nik.", nl)
	goto L0047
L32d3:
	if float64(g.f1d6) == 16 {
		goto L32e8
	}
	goto L3376
L32e8:
	g.sub_status()
	g.printS("Elkezd sikoltozni,", semi)
	if (b2i((!(6 <= float64(g.f1a6)))) | b2i((!(22 >= float64(g.f1a6))))) != 0 {
		goto L332f
	}
	goto L333b
L332f:
	g.printS("de az alulj\xa0r\xa2 kihalt.", nl)
	goto L0047
L333b:
	g.printS("", nl)
	g.printS("a j\xa0r\xa2kel\x94k szaladnak rend\x94r\x82rt.", nl)
	g.f162 = float32((float64(g.f162) + -25))
	g.f192 = float32(1)
	g.gosub = append(g.gosub, 76)
	goto L5633
R336d:
	g.sub_7a66()
	goto L0047
L3376:
	if float64(g.f1d6) == 17 {
		goto L338b
	}
	goto L339d
L338b:
	g.sub_status()
	g.printS("Hov\xa0?", nl)
	goto L0047
L339d:
	if float64(g.f1d6) == 18 {
		goto L33b2
	}
	goto L3494
L33b2:
	g.sub_status()
	g.f172 = float32(math.Floor((g.rnd() * 3)))
	if 0 == float64(g.f172) {
		goto L33ea
	}
	goto L33f7
L33ea:
	g.printS("Nincs t\xa0rc\xa0ja.", nl)
	goto L0047
L33f7:
	if 1 == float64(g.f172) {
		goto L340c
	}
	goto L345c
L340c:
	g.color(15)
	g.printS("-Seg\xa1ts\x82g!Tolvaj!", nl)
	g.color(7)
	g.printS("Egy buzg\xa2m\xa2csing h\xa0tulr\xa2l le\x81t.", nl)
	g.f162 = float32((float64(g.f162) + -15))
	g.gosub = append(g.gosub, 77)
	goto L5633
R3454:
	g.sub_7a66()
	goto L0047
L345c:
	g.printS("Elloptad \x82szrev\x82tlen\x81l.", nl)
	g.f182 = float32(((math.Floor((g.rnd() * 19)) + float64(g.f182)) + 1))
	g.sub_7a66()
	goto L0047
L3494:
	g.caseElseError()
L3499:
	g.sub_status()
	if 0 == math.Floor((g.rnd() * 2)) {
		goto L34c5
	}
	goto L34d1
L34c5:
	g.printS("Egy muks\xa2 j\x94n szembe veled.", nl)
	goto L34da
L34d1:
	g.printS("Egy k\x94z\x82pkoru embert l\xa0tnak szemeid.", nl)
L34da:
	g.sub_6a7f()
	g.f1da = float32(float64(g.f17a))
	if float64(g.f1da) != 1 {
		goto L3502
	}
	goto L3543
L3502:
	if float64(g.f1da) != 2 {
		goto L3518
	}
	goto L3543
L3518:
	if float64(g.f1da) != 3 {
		goto L352e
	}
	goto L3543
L352e:
	if float64(g.f1da) == 4 {
		goto L3543
	}
	goto L354f
L3543:
	g.sub_status()
	g.gosub = append(g.gosub, 61)
	goto L56b6
R354b:
	goto L005f
L354f:
	if float64(g.f1da) != 5 {
		goto L3565
	}
	goto L357b
L3565:
	if float64(g.f1da) == 6 {
		goto L357b
	}
	goto L35a7
L357b:
	if 6 == float64(g.f17a) {
		goto L3590
	}
	goto L3595
L3590:
	g.sub_7fc2()
L3595:
	g.sub_status()
	g.printS("Megijedt\x82l t\x94le?", nl)
	goto L0047
L35a7:
	if float64(g.f1da) != 7 {
		goto L35bd
	}
	goto L35d2
L35bd:
	if float64(g.f1da) == 8 {
		goto L35d2
	}
	goto L35f0
L35d2:
	g.sub_status()
	g.printS("Ett\x94l ugyan aligha.", nl)
	g.f142 = float32(1)
	goto L34da
L35f0:
	if float64(g.f1da) == 9 {
		goto L3605
	}
	goto L36e6
L3605:
	g.sub_status()
	g.color(15)
	l0c = (b2i((!(0 != float64(g.f142)))) & b2i((!(4 <= float64(g.f196)))))
	if (l0c & b2i((!(1 != math.Floor((g.rnd() * 3)))))) != 0 {
		goto L3682
	}
	goto L36ca
L3682:
	g.printS("-Itt van 10 forint.", nl)
	g.color(7)
	g.printS("-mondja gyanakodva.", nl)
	g.f182 = float32((float64(g.f182) + 10))
	g.sub_7a66()
	g.f142 = float32(1)
	goto L34da
L36ca:
	g.printS("-T\x81nj el,am\xa1g sz\x82pen besz\x82lek!", nl)
	g.f142 = float32(1)
	goto L34da
L36e6:
	if float64(g.f1da) == 11 {
		goto L36fb
	}
	goto L38d9
L36fb:
	g.sub_status()
	if 5 > float64(g.f186) {
		goto L3717
	}
	goto L379a
L3717:
	if (b2i((!(60 >= float64(g.f18e)))) | b2i((!(1 != float64(g.f142))))) != 0 {
		goto L374f
	}
	goto L3774
L374f:
	g.printS("H\xa0r\xa1tja a t\xa0mad\xa0st \x82s ad egy nagypofont.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.gosub = append(g.gosub, 62)
	goto L5633
R376c:
	g.sub_7a66()
	goto L34da
L3774:
	g.printS("-Leverted a nyavaly\xa0st.", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 63)
	goto L5633
R3791:
	g.sub_7a66()
	goto L0047
L379a:
	if 9 == float64(g.f186) {
		goto L37b0
	}
	goto L386d
L37b0:
	if (b2i((!(1 != float64(g.f142)))) | b2i((!(0.5 >= float64(g.f18e))))) != 0 {
		goto L37e8
	}
	goto L3819
L37e8:
	g.printS("Kicsavarja a k\x82st a kezedb\x94l \x82s j\xa2l     bok\xa0n rug.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.gosub = append(g.gosub, 64)
	goto L5633
R3805:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L34da
L3819:
	g.color(4)
	g.printS("Pontosan sziven szurtad a gazembert.", nl)
	g.f192 = float32(1)
	g.f162 = float32((float64(g.f162) + -5))
	g.gosub = append(g.gosub, 65)
	goto L5633
R3853:
	g.sub_7a66()
	g.color(7)
	goto L0047
L386d:
	if 30 < float64(g.f18e) {
		goto L3882
	}
	goto L38b3
L3882:
	g.printS("F\x82lreugrik \x82s j\xa2l sz\xa0jba ver.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.gosub = append(g.gosub, 66)
	goto L5633
R389f:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L34da
L38b3:
	g.printS("Lesulytottad a nyavaly\xa0st.", nl)
	g.f162 = float32((float64(g.f162) + -2))
	g.gosub = append(g.gosub, 67)
	goto L5633
R38d0:
	g.sub_7a66()
	goto L0047
L38d9:
	if float64(g.f1da) == 10 {
		goto L38ee
	}
	goto L39c6
L38ee:
	g.sub_status()
	if (b2i((!(1 != float64(g.f19e)))) & b2i((!(0 != float64(g.f142))))) != 0 {
		goto L392c
	}
	goto L3966
L392c:
	g.color(15)
	g.printS("-Menj a francba,\x94csi!", nl)
	g.color(7)
	g.f142 = float32(1)
	goto L34da
L3966:
	g.color(15)
	g.printS("-Az any\xa0d pics\xa0j\xa0t,kis paraszt!", nl)
	g.color(7)
	g.printS("J\xa2l gyomron v\xa0g.", nl)
	g.f162 = float32((float64(g.f162) + -15))
	g.gosub = append(g.gosub, 68)
	goto L5633
R39ae:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L34da
L39c6:
	if float64(g.f1da) != 12 {
		goto L39dc
	}
	goto L3a07
L39dc:
	if float64(g.f1da) != 14 {
		goto L39f2
	}
	goto L3a07
L39f2:
	if float64(g.f1da) == 15 {
		goto L3a07
	}
	goto L3a19
L3a07:
	g.sub_status()
	g.printS("Aktat\xa0sk\xa0j\xa0t l\xa2b\xa0lva elt\x81nik.", nl)
	goto L0047
L3a19:
	if float64(g.f1da) == 16 {
		goto L3a2e
	}
	goto L3aee
L3a2e:
	g.sub_status()
	g.color(15)
	g.printS("-Seg\xa1ts\x82g!Buzi!", nl)
	g.color(7)
	if (b2i((!(6 <= float64(g.f1a6)))) | b2i((!(22 >= float64(g.f1a6))))) != 0 {
		goto L3a97
	}
	goto L3abc
L3a97:
	g.printS("Senki se hallja meg.", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 69)
	goto L5633
R3ab4:
	g.sub_7a66()
	goto L0047
L3abc:
	g.printS("A j\xa0r\xa2kel\x94k szaladnak rend\x94r\x82rt.", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.f192 = float32(1)
	g.gosub = append(g.gosub, 70)
	goto L5633
R3ae5:
	g.sub_7a66()
	goto L0047
L3aee:
	if float64(g.f1da) == 17 {
		goto L3b03
	}
	goto L3b15
L3b03:
	g.sub_status()
	g.printS("Hov\xa0?", nl)
	goto L0047
L3b15:
	if float64(g.f1da) == 18 {
		goto L3b2a
	}
	goto L3c2f
L3b2a:
	g.sub_status()
	g.f172 = float32(math.Floor((g.rnd() * 3)))
	if 0 == float64(g.f172) {
		goto L3b62
	}
	goto L3b6f
L3b62:
	g.printS("Nincs t\xa0rc\xa0ja.", nl)
	goto L0047
L3b6f:
	if (b2i((!(1 != float64(g.f172)))) | b2i((!(1 != float64(g.f142))))) != 0 {
		goto L3ba7
	}
	goto L3bf7
L3ba7:
	g.color(15)
	g.printS("-Seg\xa1ts\x82g!Tolvaj!", nl)
	g.color(7)
	g.printS("Egy faszparaszt h\xa0tulr\xa2l le\x81t.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.gosub = append(g.gosub, 71)
	goto L5633
R3bef:
	g.sub_7a66()
	goto L0047
L3bf7:
	g.printS("Elloptad \x82szrev\x82tlen\x81l.", nl)
	g.f182 = float32(((math.Floor((g.rnd() * 39)) + float64(g.f182)) + 1))
	g.sub_7a66()
	goto L0047
L3c2f:
	g.caseElseError()
L3c35:
	if (b2i((!(8 >= float64(g.f1a6)))) & b2i((!(18 <= float64(g.f1a6))))) == 0 {
		goto L3c6d
	}
	goto L18f2
L3c6d:
	g.sub_status()
	g.f172 = float32(math.Floor((g.rnd() * 4)))
	if 0 == float64(g.f172) {
		goto L3ca5
	}
	goto L3cc9
L3ca5:
	g.printS("Egy szakadt cs\x94ves ugrik el\x94.", nl)
	g.color(15)
	g.printS("-A kurva any\xa0d,te szem\x82t Nyugtatis dig\xa2!", nl)
L3cc9:
	if (b2i((!(1 != float64(g.f172)))) | b2i((!(2 != float64(g.f172))))) != 0 {
		goto L3d01
	}
	goto L3d2e
L3d01:
	g.printS("Egy Heavy Metal-os ugrik el\x94.", nl)
	g.color(15)
	g.printS("-A kurva any\xa0d,te rohadt", nl)
	g.printS("Modern Talkingos!", nl)
L3d2e:
	if 3 == float64(g.f172) {
		goto L3d43
	}
	goto L3d66
L3d43:
	g.printS("Egy skin-head ugrik el\x94.", nl)
	g.color(15)
	g.printS("-A kurva any\xa0d!", nl)
L3d66:
	g.sub_6a7f()
	g.f1de = float32(float64(g.f17a))
	if float64(g.f1de) != 1 {
		goto L3d8e
	}
	goto L3dcf
L3d8e:
	if float64(g.f1de) != 2 {
		goto L3da4
	}
	goto L3dcf
L3da4:
	if float64(g.f1de) != 3 {
		goto L3dba
	}
	goto L3dcf
L3dba:
	if float64(g.f1de) == 4 {
		goto L3dcf
	}
	goto L3e14
L3dcf:
	g.sub_status()
	g.printS("Ut\xa0nad rohan \x82s seggberug.", nl)
	g.color(15)
	g.printS("-Nem hallod? A kurva any\xa0d!", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 48)
	goto L5633
R3e0b:
	g.sub_7a66()
	goto L3d66
L3e14:
	if float64(g.f1de) != 5 {
		goto L3e2a
	}
	goto L3e40
L3e2a:
	if float64(g.f1de) == 6 {
		goto L3e40
	}
	goto L3f4b
L3e40:
	if (b2i((!(30 <= float64(g.f162)))) | b2i((!(1 >= float64(g.f18e))))) != 0 {
		goto L3e78
	}
	goto L3ea3
L3e78:
	g.sub_status()
	g.printS("Ut\xa0nad ered,utol\x82r \x82s j\xa2l t\x94k\x94nver.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.gosub = append(g.gosub, 49)
	goto L5633
R3e9a:
	g.sub_7a66()
	goto L3d66
L3ea3:
	if 6 == float64(g.f17a) {
		goto L3eb8
	}
	goto L3ebe
L3eb8:
	g.sub_7fc2()
L3ebe:
	if (b2i((!(30 <= float64(g.f162)))) | b2i((!(1 >= float64(g.f18e))))) != 0 {
		goto L3ef6
	}
	goto L3f20
L3ef6:
	g.sub_status()
	g.printS("Ut\xa0nad ered,utol\x82r \x82s j\xa2l t\x94k\x94nver.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.gosub = append(g.gosub, 50)
	goto L5633
R3f18:
	g.sub_7a66()
	goto L3d66
L3f20:
	g.sub_status()
	g.printS("Megijedt\x82l t\x94le?", nl)
	g.f162 = float32((float64(g.f162) + -2))
	g.gosub = append(g.gosub, 51)
	goto L5633
R3f42:
	g.sub_7a66()
	goto L0047
L3f4b:
	if float64(g.f1de) != 7 {
		goto L3f61
	}
	goto L3f76
L3f61:
	if float64(g.f1de) == 8 {
		goto L3f76
	}
	goto L3fcc
L3f76:
	g.sub_status()
	g.color(15)
	g.printS("-Mi van,baszd meg,fosol?", nl)
	g.color(7)
	g.printS("Lefejel.", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 52)
	goto L5633
R3fc3:
	g.sub_7a66()
	goto L3d66
L3fcc:
	if float64(g.f1de) == 9 {
		goto L3fe1
	}
	goto L4037
L3fe1:
	g.sub_status()
	g.color(15)
	g.printS("-Mi van,baszd meg,t\x94lem akarsz lejmolni?", nl)
	g.color(7)
	g.printS("Gyomorsz\xa0jon v\xa0g.", nl)
	g.f162 = float32((float64(g.f162) + -15))
	g.gosub = append(g.gosub, 53)
	goto L5633
R402e:
	g.sub_7a66()
	goto L3d66
L4037:
	if float64(g.f1de) == 11 {
		goto L404c
	}
	goto L41b3
L404c:
	g.sub_status()
	if 5 > float64(g.f186) {
		goto L4068
	}
	goto L40c8
L4068:
	if 60 < float64(g.f18e) {
		goto L407d
	}
	goto L40a2
L407d:
	g.printS("H\xa0r\xa1tja a t\xa0mad\xa0st \x82s orrba ver.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.gosub = append(g.gosub, 54)
	goto L5633
R409a:
	g.sub_7a66()
	goto L0047
L40a2:
	g.printS("Leverted a nyavaly\xa0st.", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 55)
	goto L5633
R40bf:
	g.sub_7a66()
	goto L0047
L40c8:
	if 9 == float64(g.f186) {
		goto L40de
	}
	goto L416c
L40de:
	if 80 < float64(g.f18e) {
		goto L40f3
	}
	goto L4124
L40f3:
	g.printS("Kicsavarja a k\x82st a kezedb\x94l \x82s j\xa2l     \xa0llcsucson \x81t.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.gosub = append(g.gosub, 56)
	goto L5633
R4110:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L3d66
L4124:
	g.color(4)
	g.printS("Pontosan sziven szurtad a diszn\xa2t.", nl)
	g.f162 = float32((float64(g.f162) + -2))
	g.gosub = append(g.gosub, 57)
	goto L5633
R4152:
	g.sub_7a66()
	g.color(7)
	goto L0047
L416c:
	if 30 < float64(g.f18e) {
		goto L4181
	}
	goto L41a6
L4181:
	g.printS("F\x82lreugrik \x82s j\xa2l sz\xa0jba ver.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.gosub = append(g.gosub, 58)
	goto L5633
R419e:
	g.sub_7a66()
	goto L3d66
L41a6:
	g.printS("Lesulytottad a nyavaly\xa0st.", nl)
	goto L0047
L41b3:
	if float64(g.f1de) == 10 {
		goto L41c8
	}
	goto L421e
L41c8:
	g.sub_status()
	g.color(15)
	g.printS("-Te rohadt patk\xa0ny!", nl)
	g.color(7)
	g.printS("J\xa2l gyomron rug.", nl)
	g.f162 = float32((float64(g.f162) + -15))
	g.gosub = append(g.gosub, 59)
	goto L5633
R4215:
	g.sub_7a66()
	goto L3d66
L421e:
	if float64(g.f1de) != 12 {
		goto L4234
	}
	goto L42a1
L4234:
	if float64(g.f1de) != 14 {
		goto L424a
	}
	goto L42a1
L424a:
	if float64(g.f1de) != 15 {
		goto L4260
	}
	goto L42a1
L4260:
	if float64(g.f1de) != 16 {
		goto L4276
	}
	goto L42a1
L4276:
	if float64(g.f1de) != 17 {
		goto L428c
	}
	goto L42a1
L428c:
	if float64(g.f1de) == 18 {
		goto L42a1
	}
	goto L42cb
L42a1:
	g.sub_status()
	g.printS("Sz\x94rny\x81 haragra gerjed,nyakon v\xa0g \x82s    lel\x82p", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 60)
	goto L5633
R42c3:
	g.sub_7a66()
	goto L0047
L42cb:
	g.caseElseError()
L42d1:
	if (b2i((!(6 <= float64(g.f1a6)))) | b2i((!(21 >= float64(g.f1a6))))) == 0 {
		goto L4309
	}
	goto L18f2
L4309:
	g.sub_status()
	g.printS("Egy paraszt j\x94n szembe veled.", nl)
L4317:
	g.sub_6a7f()
	g.f1e2 = float32(float64(g.f17a))
	if float64(g.f1e2) != 1 {
		goto L433f
	}
	goto L4380
L433f:
	if float64(g.f1e2) != 2 {
		goto L4355
	}
	goto L4380
L4355:
	if float64(g.f1e2) != 3 {
		goto L436b
	}
	goto L4380
L436b:
	if float64(g.f1e2) == 4 {
		goto L4380
	}
	goto L4392
L4380:
	g.sub_status()
	g.printS("Kiker\x81lted.", nl)
	goto L0047
L4392:
	if float64(g.f1e2) != 5 {
		goto L43a8
	}
	goto L43be
L43a8:
	if float64(g.f1e2) == 6 {
		goto L43be
	}
	goto L4476
L43be:
	if 6 == float64(g.f17a) {
		goto L43d3
	}
	goto L4464
L43d3:
	g.sub_7fc2()
	g.sub_status()
	g.f142 = float32(1)
	g.color(15)
	g.printS("-Hijjnye,hogy az a f\x81zf\xa0n f\x81ty\x81l\x94s      mind\x94ns\x82git!!!", nl)
	g.color(7)
	g.printS("Ut\xa0nad haj\xa1tja a kamp\xa2sbotj\xa0t.", nl)
	if 30 < float64(g.f18e) {
		goto L4433
	}
	goto L4458
L4433:
	g.printS("El is tal\xa0l vele.", nl)
	g.f162 = float32((float64(g.f162) + -25))
	g.gosub = append(g.gosub, 40)
	goto L5633
R4450:
	g.sub_7a66()
	goto L4461
L4458:
	g.printS("Nem tal\xa0l el.", nl)
L4461:
	goto L0047
L4464:
	g.sub_status()
	g.printS("Megl\x82pt\x82l el\x94le.", nl)
	goto L0047
L4476:
	if float64(g.f1e2) == 7 {
		goto L448b
	}
	goto L4570
L448b:
	g.sub_status()
	g.color(15)
	if (b2i((!(1 != float64(g.f142)))) | b2i((!(30 <= float64(g.f182))))) != 0 {
		goto L44da
	}
	goto L44e6
L44da:
	g.printS("-Nem adok \x82n,Isten uccse',semmi p\xa1z\x82r'!", nl)
	goto L4317
L44e6:
	g.printS("-Itt van ni a butykosba remek kis\x81sti,  vigye csak kend!", nl)
	g.f1b2 = float32((float64(g.f1b2) + 1))
	if 1 < float64(g.f1b2) {
		goto L4517
	}
	goto L4557
L4517:
	g.f1b2 = float32(1)
	g.color(5)
	g.printS("", nl)
	g.printS("Elejtetted az \x81veget,te \x81gyetlen!", nl)
	g.color(7)
L4557:
	g.f182 = float32((float64(g.f182) + -30))
	g.sub_7a66()
	goto L4317
L4570:
	if float64(g.f1e2) == 8 {
		goto L4585
	}
	goto L45c5
L4585:
	g.sub_status()
	g.color(15)
	g.printS("-Mit k\xa0romolsz itt,muramista?!?", nl)
	g.color(7)
	g.f142 = float32(1)
	goto L4317
L45c5:
	if float64(g.f1e2) == 9 {
		goto L45da
	}
	goto L466d
L45da:
	g.sub_status()
	g.color(15)
	g.f172 = float32(math.Floor((g.rnd() * 5)))
	if 1 > float64(g.f172) {
		goto L4623
	}
	goto L4651
L4623:
	g.printS("-Nesze,itt van e 20 f\xa2rint.", nl)
	g.f182 = float32((float64(g.f182) + 20))
	g.f142 = float32(1)
	g.sub_7a66()
	goto L4317
L4651:
	g.printS("-Huzz\xa0l,mint a vadludak,te fattyu!", nl)
	g.f142 = float32(1)
	goto L4317
L466d:
	if float64(g.f1e2) == 11 {
		goto L4682
	}
	goto L47f5
L4682:
	g.sub_status()
	if 5 > float64(g.f186) {
		goto L469e
	}
	goto L470a
L469e:
	if 60 < float64(g.f18e) {
		goto L46b3
	}
	goto L46e4
L46b3:
	g.printS("Meg se kottyan neki,f\x81lt\x94v\x94n v\xa0g.", nl)
	g.f162 = float32((float64(g.f162) + -40))
	g.gosub = append(g.gosub, 41)
	goto L5633
R46d0:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L4317
L46e4:
	g.printS("Leverted a p\xa2r\x82t.", nl)
	g.f162 = float32((float64(g.f162) + -10))
	g.gosub = append(g.gosub, 42)
	goto L5633
R4701:
	g.sub_7a66()
	goto L0047
L470a:
	if 9 == float64(g.f186) {
		goto L4720
	}
	goto L47ae
L4720:
	if 80 < float64(g.f18e) {
		goto L4735
	}
	goto L4766
L4735:
	g.printS("Ki\x81ti a k\x82st a kezedb\x94l \x82s t\x94k\x94n rug.", nl)
	g.f162 = float32((float64(g.f162) + -45))
	g.gosub = append(g.gosub, 43)
	goto L5633
R4752:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L4317
L4766:
	g.color(4)
	g.printS("Hasba szurtad a diszn\xa2t.", nl)
	g.f162 = float32((float64(g.f162) + -2))
	g.gosub = append(g.gosub, 44)
	goto L5633
R4794:
	g.sub_7a66()
	g.color(7)
	goto L0047
L47ae:
	if 30 < float64(g.f18e) {
		goto L47c3
	}
	goto L47e8
L47c3:
	g.printS("F\x82lreugrik \x82s j\xa2l orrba ver.", nl)
	g.f162 = float32((float64(g.f162) + -40))
	g.gosub = append(g.gosub, 45)
	goto L5633
R47e0:
	g.sub_7a66()
	goto L4317
L47e8:
	g.printS("Lesulytottad a nyavaly\xa0s p\xa2r\x82t.", nl)
	goto L0047
L47f5:
	if float64(g.f1e2) == 10 {
		goto L480a
	}
	goto L486c
L480a:
	g.sub_status()
	g.color(15)
	g.printS("-H\xa1jj,azt a s\xa0rga l\xa0bu istenit az       any\xa0dnak!", nl)
	g.color(7)
	g.printS("J\xa2l gyomron rug.", nl)
	g.f162 = float32((float64(g.f162) + -30))
	g.gosub = append(g.gosub, 46)
	goto L5633
R4857:
	g.sub_7a66()
	g.f142 = float32(1)
	goto L4317
L486c:
	if float64(g.f1e2) == 12 {
		goto L4881
	}
	goto L4893
L4881:
	g.sub_status()
	g.printS("Amig v\xa0rsz,elmegy.", nl)
	goto L494f
L4893:
	if float64(g.f1e2) != 16 {
		goto L48a9
	}
	goto L48be
L48a9:
	if float64(g.f1e2) == 18 {
		goto L48be
	}
	goto L48e9
L48be:
	g.sub_status()
	g.printS("Sz\x94rny\x81 haragra gerjed,nyakon v\xa0g \x82s    lel\x82p", nl)
	g.f162 = float32((float64(g.f162) + -30))
	g.gosub = append(g.gosub, 47)
	goto L5633
R48e0:
	g.sub_7a66()
	goto L0047
L48e9:
	if float64(g.f1e2) != 14 {
		goto L48ff
	}
	goto L4914
L48ff:
	if float64(g.f1e2) == 15 {
		goto L4914
	}
	goto L4926
L4914:
	g.sub_status()
	g.printS("Elt\x81nik a szemed el\x94l.", nl)
	goto L494f
L4926:
	if float64(g.f1e2) == 17 {
		goto L493b
	}
	goto L494c
L493b:
	g.sub_status()
	g.printS("Hov\xa0?", nl)
	goto L4317
L494c:
	goto L0047
L494f:
	goto L0047
L4952:
	g.sub_status()
	if 0 == float64(g.f1e6) {
		goto L496d
	}
	goto L4991
L496d:
	g.printS("Egy cig\xa0nypurd\x82 ugrik el\x94.", nl)
	g.color(15)
	g.printS("-Cs\xa0v\xa2k\xa0m,adj egy b\x82l\xa0st,gy\xa2gyszerre!", nl)
L4991:
	if 1 == float64(g.f1e6) {
		goto L49a6
	}
	goto L49c9
L49a6:
	g.printS("Tal\xa0lkozt\xa0l h\xa0rom gy\xa0li cig\xa0nnyal.", nl)
	g.color(15)
	g.printS("-Te G\xa0zsi,ezs va\xa2t azs,\xa0ki a mutko      gecizett a Lajcsival!", nl)
L49c9:
	g.sub_6a7f()
	g.f1ea = float32(float64(g.f17a))
	if float64(g.f1ea) != 1 {
		goto L49f1
	}
	goto L4a32
L49f1:
	if float64(g.f1ea) != 2 {
		goto L4a07
	}
	goto L4a32
L4a07:
	if float64(g.f1ea) != 3 {
		goto L4a1d
	}
	goto L4a32
L4a1d:
	if float64(g.f1ea) == 4 {
		goto L4a32
	}
	goto L4ac2
L4a32:
	g.sub_status()
	if 0 == float64(g.f1e6) {
		goto L4a4d
	}
	goto L4a84
L4a4d:
	g.printS("El\x82d \xa0ll.", nl)
	g.color(15)
	g.printS("-Na aggy\xa0l m\xa0,mer le\x81tlek!", nl)
	g.color(7)
	goto L49c9
L4a84:
	g.printS("K\x94r\x81lvesznek \x82s j\xa2l le\x81tnek.", nl)
	g.f162 = float32((float64(g.f162) + -50))
	g.gosub = append(g.gosub, 26)
	goto L5633
R4aa2:
	g.f1e6 = float32(0)
	g.f182 = float32(0)
	g.sub_7a66()
	goto L0047
L4ac2:
	if float64(g.f1ea) != 5 {
		goto L4ad8
	}
	goto L4aee
L4ad8:
	if float64(g.f1ea) == 6 {
		goto L4aee
	}
	goto L4ba9
L4aee:
	if 6 == float64(g.f17a) {
		goto L4b03
	}
	goto L4b08
L4b03:
	g.sub_7fc2()
L4b08:
	g.sub_status()
	if 0 == float64(g.f1e6) {
		goto L4b23
	}
	goto L4b6b
L4b23:
	g.color(15)
	g.printS("-Mingy\xa0 gy\x81n a 12 b\xa0ty\xa0m,oszt kiny\xa1rnak!", semi)
	g.color(7)
	g.printS("ord\xa1tja ut\xa0nad.", nl)
	g.f1e6 = float32(1)
	g.sub_7a66()
	goto L0047
L4b6b:
	g.printS("Ut\xa0nad erednek,utol\x82rnek \x82s j\xa2l         megrugdosnak.", nl)
	g.f162 = float32((float64(g.f162) + -55))
	g.gosub = append(g.gosub, 27)
	goto L5633
R4b89:
	g.f182 = float32(0)
	g.f1e6 = float32(0)
	g.sub_7a66()
	goto L0047
L4ba9:
	if float64(g.f1ea) != 7 {
		goto L4bbf
	}
	goto L4bd4
L4bbf:
	if float64(g.f1ea) == 8 {
		goto L4bd4
	}
	goto L4c70
L4bd4:
	g.sub_status()
	g.color(15)
	if 0 == float64(g.f1e6) {
		goto L4c00
	}
	goto L4c0c
L4c00:
	g.printS("Na ne gecizz\x82 m\xa0,cs\xa0v\xa2k\xa0m!", nl)
	goto L49c9
L4c0c:
	g.printS("-Ne,te bib\xa0s,szopacsaki mak\xa0!", nl)
	g.color(7)
	g.printS("Megtaposs\xa0k a fejedet.", nl)
	g.f162 = float32((float64(g.f162) + -50))
	g.f182 = float32(0)
	g.f1e6 = float32(0)
	g.gosub = append(g.gosub, 28)
	goto L5633
R4c5c:
	g.f182 = float32(0)
	g.sub_7a66()
	goto L0047
L4c70:
	if float64(g.f1ea) == 9 {
		goto L4c85
	}
	goto L4d0c
L4c85:
	g.sub_status()
	if 0 == float64(g.f1e6) {
		goto L4ca0
	}
	goto L4cce
L4ca0:
	g.color(15)
	g.printS("-Azt a j\xa2 kurva any\xa0dat,te paraszt!", nl)
	g.f1e6 = float32(1)
	g.sub_7a66()
	goto L0047
L4cce:
	g.printS("Nem \x82rt\x82kelik a humorodat.Lecsapnak.", nl)
	g.f1e6 = float32(0)
	g.f162 = float32((float64(g.f162) + -50))
	g.gosub = append(g.gosub, 29)
	goto L5633
R4cf8:
	g.f182 = float32(0)
	g.sub_7a66()
	goto L0047
L4d0c:
	if float64(g.f1ea) == 11 {
		goto L4d21
	}
	goto L4f74
L4d21:
	g.sub_status()
	if (b2i((!(0 != float64(g.f1e6)))) & b2i((!(9 == float64(g.f186))))) != 0 {
		goto L4d5f
	}
	goto L4d78
L4d5f:
	g.printS("Ugy megverted,hogy s\xa1rva szalad haza.", nl)
	g.f1e6 = float32(1)
	goto L0047
L4d78:
	if 5 > float64(g.f186) {
		goto L4d8e
	}
	goto L4e12
L4d8e:
	if 60 < float64(g.f18e) {
		goto L4da3
	}
	goto L4de0
L4da3:
	g.printS("Meg se tudsz moccanni \x82s m\xa0ris levernek.", nl)
	g.f162 = float32((float64(g.f162) + -55))
	g.f1e6 = float32(0)
	g.f182 = float32(0)
	g.gosub = append(g.gosub, 30)
	goto L5633
R4dd8:
	g.sub_7a66()
	goto L0047
L4de0:
	g.printS("Egyet levert\x82l,a t\x94bbi elszalad.", nl)
	g.f162 = float32((float64(g.f162) + -30))
	g.gosub = append(g.gosub, 31)
	goto L5633
R4dfe:
	g.f1e6 = float32(0)
	g.sub_7a66()
	goto L0047
L4e12:
	if 9 == float64(g.f186) {
		goto L4e28
	}
	goto L4ef0
L4e28:
	if 0 == float64(g.f1e6) {
		goto L4e3d
	}
	goto L4e6c
L4e3d:
	g.color(4)
	g.printS("Led\x94fted a roma cs\xa0v\xa2t.", nl)
	g.color(7)
	goto L0047
L4e6c:
	if 50 < float64(g.f18e) {
		goto L4e81
	}
	goto L4e9c
L4e81:
	g.printS("Egyet leszurt\xa0l,de a m\xa0sik kett\x94        \xa0tv\xa0gja a torkodat.", nl)
	g.f162 = float32(0)
	g.gosub = append(g.gosub, 32)
	goto L5633
R4e99:
	goto L4ef0
L4e9c:
	g.color(4)
	g.printS("Pontosan sziven szurtad a rajk\xa2t,a      m\xa0sik kett\x94 lel\x82p.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.gosub = append(g.gosub, 35)
	goto L5633
R4ecb:
	g.f1e6 = float32(0)
	g.sub_7a66()
	g.color(7)
	goto L0047
L4ef0:
	if 30 < float64(g.f18e) {
		goto L4f05
	}
	goto L4f42
L4f05:
	g.printS("Lassu vagy,j\xa2l elp\xa0holnak.", nl)
	g.f162 = float32((float64(g.f162) + -45))
	g.f1e6 = float32(0)
	g.f182 = float32(0)
	g.gosub = append(g.gosub, 33)
	goto L5633
R4f3a:
	g.sub_7a66()
	goto L0047
L4f42:
	g.printS("Lesulytott\xa0l egyet,a t\x94bbi lel\x82cel.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.gosub = append(g.gosub, 34)
	goto L5633
R4f60:
	g.f1e6 = float32(0)
	g.sub_7a66()
	goto L0047
L4f74:
	if float64(g.f1ea) == 10 {
		goto L4f89
	}
	goto L503f
L4f89:
	g.sub_status()
	g.color(15)
	if 0 == float64(g.f1e6) {
		goto L4fb5
	}
	goto L4fe7
L4fb5:
	g.printS("-A tetves geci fejedet,te paraszt!", nl)
	g.color(7)
	g.printS("orditja,es elfut.", nl)
	g.f1e6 = float32(1)
	goto L0047
L4fe7:
	g.printS("-Te rohadt geci paraszt!", nl)
	g.color(7)
	g.printS("J\xa2l hasba rugdalnak.", nl)
	g.f162 = float32((float64(g.f162) + -50))
	g.f1e6 = float32(0)
	g.f182 = float32(0)
	g.gosub = append(g.gosub, 36)
	goto L5633
R5036:
	g.sub_7a66()
	goto L0047
L503f:
	if float64(g.f1ea) == 12 {
		goto L5054
	}
	goto L5122
L5054:
	g.sub_status()
	if 0 == float64(g.f1e6) {
		goto L506f
	}
	goto L50cc
L506f:
	g.printS("Belenyul a zsebedbe.", nl)
	g.f182 = float32((float64(g.f182) - (math.Floor((g.rnd() * 10)) + 10)))
	if 0 > float64(g.f182) {
		goto L50b9
	}
	goto L50c4
L50b9:
	g.f182 = float32(0)
L50c4:
	g.sub_7a66()
	goto L0047
L50cc:
	g.printS("Piszkosul megvernek.", nl)
	g.f162 = float32((float64(g.f162) + -45))
	g.f182 = float32(0)
	g.f1b2 = float32(0)
	g.f1b6 = float32(0)
	g.gosub = append(g.gosub, 37)
	goto L5633
R510e:
	g.f1e6 = float32(0)
	g.sub_7a66()
	goto L0047
L5122:
	if float64(g.f1ea) != 14 {
		goto L5138
	}
	goto L518f
L5138:
	if float64(g.f1ea) != 15 {
		goto L514e
	}
	goto L518f
L514e:
	if float64(g.f1ea) != 16 {
		goto L5164
	}
	goto L518f
L5164:
	if float64(g.f1ea) != 17 {
		goto L517a
	}
	goto L518f
L517a:
	if float64(g.f1ea) == 18 {
		goto L518f
	}
	goto L5218
L518f:
	g.sub_status()
	if 0 == float64(g.f1e6) {
		goto L51aa
	}
	goto L51db
L51aa:
	g.printS("J\xa2l megver \x82s t\xa0vozik.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.f1e6 = float32(1)
	g.gosub = append(g.gosub, 38)
	goto L5633
R51d3:
	g.sub_7a66()
	goto L0047
L51db:
	g.printS("J\xa2l megvernek \x82s t\xa0voznak.", nl)
	g.f162 = float32((float64(g.f162) + -45))
	g.f1e6 = float32(0)
	g.f182 = float32(0)
	g.gosub = append(g.gosub, 39)
	goto L5633
R5210:
	g.sub_7a66()
	goto L0047
L5218:
	goto L0047
L521b:
	g.sub_status()
	g.color(9)
	if 0 == float64(g.f192) {
		goto L5247
	}
	goto L527f
L5247:
	g.printS("Egy rend\x94r gyanakodva n\x82z,de azt\xa0n      od\x82bb\xa0ll. Ezt meg\xa3sztad.", nl)
L5250:
	g.s13e = g.inputS(1)
	if chr(13) != g.s13e {
		goto L5250
	}
	goto L005f
L527f:
	g.printS("Elkapott egy rend\x94r.", nl)
L5288:
	g.sub_6a7f()
	g.sub_status()
	if (b2i((!(5 != float64(g.f17a)))) | b2i((!(6 != float64(g.f17a))))) != 0 {
		goto L52cc
	}
	goto L5345
L52cc:
	if 6 == float64(g.f17a) {
		goto L52e1
	}
	goto L52fc
L52e1:
	g.sub_7fc2()
	g.sub_status()
	g.color(7)
L52fc:
	g.printS("Ut\xa0nad l\x94,", semi)
	if 20 < float64(g.f18e) {
		goto L531b
	}
	goto L5338
L531b:
	g.printS("minden els\x94t\x82t\x81l el\x94tted.", nl)
	g.f162 = float32(0)
	g.gosub = append(g.gosub, 5)
	goto L5633
R5333:
	g.end()
L5338:
	g.printS("de nem tal\xa0l el.", nl)
	goto L0047
L5345:
	if 9 == float64(g.f186) {
		goto L535b
	}
	goto L53b3
L535b:
	if 20 < float64(g.f18e) {
		goto L5370
	}
	goto L538d
L5370:
	g.printS("Lass\xa3 vagy;lepuffant,mint egy kuty\xa0t.", nl)
	g.f162 = float32(0)
	g.gosub = append(g.gosub, 6)
	goto L5633
R5388:
	g.end()
L538d:
	g.color(4)
	g.printS("Csak a tork\xa0t v\xa0gtad el,m\x82gis kinyiffant ", nl)
	g.printS("Te \xa0rtatlan vagy,v\x82letlen baleset volt.", nl)
	goto L0047
L53b3:
	g.printS("J\xa2l fejbe ver a gumibottal \x82s bevisz.", nl)
	g.s13e = g.inputS(1)
	g.cls(-1)
	g.color(7)
	g.printS("Egy \x94rj\x94ng\x94 fel\x81gyel\x94 ugrik el\x94,", nl)
	g.printS("\x82s j\xa2l megpuh\xa1t egy gumibottal.", nl)
	g.printS("", nl)
	g.printS("", nl)
	fcarry = 0
	goto L5425
L541b:
	fcarry = (float64(g.f15e) + 1)
L5425:
	g.f15e = float32(fcarry)
	if 8000 >= float64(g.f15e) {
		goto L541b
	}
	g.printS("Egy m\xa0sik fel\x81gyel\x94 j\x94n be.", nl)
	g.printS("", nl)
	g.color(15)
	g.printS("-Mit k\x82pzel,maga vad\xa0llat,hogy b\xa0nik", nl)
	g.printS("ezzel a szerencs\x82tlen gyermekkel ?!?", nl)
	g.printS("Hiszen l\xa0tszik,hogy nem volt gyerek-", nl)
	g.printS("szob\xa0ja!", semi)
	g.color(7)
	g.printS("-kizavarja az els\x94 zsarut.", nl)
	g.printS("", nl)
	fcarry = 0
	goto L54bb
L54b1:
	fcarry = (float64(g.f15e) + 1)
L54bb:
	g.f15e = float32(fcarry)
	if 10000 >= float64(g.f15e) {
		goto L54b1
	}
	g.printS("(Valami azt s\xa3gja neked,hogy ez csak", nl)
	g.printS(" sz\xa1nj\xa0t\x82k)", nl)
	fcarry = 0
	goto L54f9
L54ef:
	fcarry = (float64(g.f15e) + 1)
L54f9:
	g.f15e = float32(fcarry)
	if 15000 >= float64(g.f15e) {
		goto L54ef
	}
	g.cls(-1)
	g.color(15)
	g.printS("-N\x82zd fiam,\x82n csak a javadat akarom,", nl)
	g.printS("tudom,hogy t\x82ged csak a t\x94bbi gazember", nl)
	g.printS("k\x82nyszer\xa1tett r\xa0,\x82n seg\xa1tek neked,", nl)
	g.printS("csak valld be,a t\xa0rsaid m\xa0r \xa3gyis ", nl)
	g.printS("mindent bevallottak,ne akard,hogy", nl)
	g.printS("durva legyek,blab-blab,...stb.stb.", nl)
	fcarry = 0
	goto L5575
L556b:
	fcarry = (float64(g.f15e) + 1)
L5575:
	g.f15e = float32(fcarry)
	if 20000 >= float64(g.f15e) {
		goto L556b
	}
	g.printS("", nl)
	g.printS("", nl)
	g.printS("-Rohadt zsaru!", semi)
	g.color(7)
	g.printS("-mondod viccesen.", nl)
	g.printS("", nl)
	g.printS("De \x94 nem \x82rti a tr\x82f\xa0t.", nl)
	g.printS("", nl)
	g.printS("(Addig \x81lsz,m\xa1g belez\x94ld\x81lsz.)", nl)
	g.color(31)
	g.printS("", nl)
	g.printS("", nl)
	g.printS("V\x82ged!", nl)
	g.s13e = g.inputS(1)
	g.clear()
	goto L0030
L5633:
	if 0 < float64(g.f162) {
		goto L5648
	}
	goto L564a
L5648:
	goto gosubReturn
L564a:
	g.f162 = float32(0)
	g.sub_7a66()
	g.locate(25, 1)
	g.color(31, 0)
	g.printS("V\x82ged!", semi)
	g.sub_6ee3()
	g.s13e = g.inputS(1)
	g.clear()
	goto L0030
L56b6:
	g.f1ee = float32(float64(g.f17a))
	if float64(g.f1ee) == 1 {
		goto L56d8
	}
	goto L56f3
L56d8:
	g.f16a = float32(1)
	g.f16e = float32(0)
	goto L5786
L56f3:
	if float64(g.f1ee) == 2 {
		goto L5709
	}
	goto L5724
L5709:
	g.f16a = float32(-1)
	g.f16e = float32(0)
	goto L5786
L5724:
	if float64(g.f1ee) == 3 {
		goto L573a
	}
	goto L5755
L573a:
	g.f16a = float32(0)
	g.f16e = float32(1)
	goto L5786
L5755:
	if float64(g.f1ee) == 4 {
		goto L576b
	}
	goto L5786
L576b:
	g.f16a = float32(0)
	g.f16e = float32(-1)
	goto L5786
L5786:
	g.f14a = float32((float64(g.f14a) + float64(g.f16a)))
	g.f146 = float32((float64(g.f146) + float64(g.f16e)))
	goto gosubReturn
L57a8:
	g.color(7, 0, 0)
	g.cls(-1)
	g.printS("Ugy d\x94ntessz,hogy elhagyod az alulj\xa0r\xa2t", nl)
	g.printS("\x82s j\xa2 \xa3tra t\x82rsz.", nl)
	g.s13e = g.inputS(1)
	g.cls(-1)
	g.printS("Tal\xa0lkozt\xa0l a r\x82gi band\xa0d 3 tagj\xa0val.", nl)
	g.printS("", nl)
	g.printS("Nem ismernek fel.", nl)
	g.printS("", nl)
	g.color(15)
	g.printS("-Ide a stexet,te rohadt dig\xa2!", nl)
	g.s13e = g.inputS(1)
	g.printS("", nl)
	if 0 == float64(g.f182) {
		goto L5876
	}
	goto L598c
L5876:
	g.printS("-Nincs p\x82nzem!", semi)
	g.color(7)
	g.printS("-ord\xa1tod riadtan.", nl)
	g.printS("", nl)
	g.color(31)
	g.printS("Akkor kiny\xa1runk,hapsik\xa0m!", nl)
	fcarry = 1
	goto L58cf
L58c5:
	fcarry = (float64(g.f15e) + 1)
L58cf:
	g.f15e = float32(fcarry)
	if 1000 >= float64(g.f15e) {
		goto L58c5
	}
	g.color(omit, omit, 2)
	fcarry = 1
	goto L5911
L5907:
	fcarry = (float64(g.f15e) + 1)
L5911:
	g.f15e = float32(fcarry)
	if 100 >= float64(g.f15e) {
		goto L5907
	}
	g.color(omit, omit, 0)
	g.color(7)
	g.printS("", nl)
	g.printS("", nl)
	g.printS("V\x82ged!", nl)
	g.s13e = g.inputS(1)
	g.clear()
	goto L0030
L598c:
	g.printS("-Csak", semi)
	g.printF(float64(g.f182), semi)
	g.printS(" forint van n\xa0lam!", semi)
	g.color(7)
	g.printS("-ord\xa1tod riadtan.", nl)
	g.printS("", nl)
	g.color(15)
	g.printS("Kev\x82s.Kiny\xa1runk,hapsik\xa0m!", nl)
	g.color(omit, omit, 2)
	fcarry = 1
	goto L5a11
L5a07:
	fcarry = (float64(g.f15e) + 1)
L5a11:
	g.f15e = float32(fcarry)
	if 1000 >= float64(g.f15e) {
		goto L5a07
	}
	fcarry = 1
	goto L5a3d
L5a33:
	fcarry = (float64(g.f15e) + 1)
L5a3d:
	g.f15e = float32(fcarry)
	if 100 >= float64(g.f15e) {
		goto L5a33
	}
	g.color(omit, omit, 0)
	g.color(7)
	g.printS("", nl)
	g.printS("", nl)
	g.printS("V\x82ged!", nl)
	g.s13e = g.inputS(1)
	g.clear()
	goto L0030
L5ab5:
	g.sub_status()
	g.color(9)
	g.printS("Amilyen farokgyerek vagy,h\xa0t nem", nl)
	g.printS("beugott\xa0l a metr\xa2sinekre?!?", nl)
	g.printS("(Term\x82szetesen GAME OVER)", nl)
	g.s13e = g.inputS(1)
	g.clear()
	goto L0030
L5b05:
	g.sub_status()
	g.color(15)
	g.printS("-J\xa0hh\xa0hh\xa0\xa0\xa0\xa0...!", semi)
	g.color(7)
	g.printS("-ord\xa1tod,", nl)
	g.printS("\x82s p\x94r\x94gve zuhansz le egy csatorn\xa0ba.", nl)
	g.f162 = float32((float64(g.f162) + -20))
	g.gosub = append(g.gosub, 2)
	goto L5633
R5b5b:
	g.printS("K\xa0romkodva kim\xa0szol \x82s visszarakod a ", nl)
	g.printS("fedel\x82t.", nl)
	g.a0da[cint(float64(g.f146))] = ((left(g.a0da[cint(float64(g.f146))], cint((float64(g.f14a)+-1))) + "\x94") + right(g.a0da[cint(float64(g.f146))], cint((float64(len(g.a0da[cint(float64(g.f146))]))-float64(g.f14a)))))
	g.sub_7a66()
	g.f16a = float32(0)
	g.f16e = float32(0)
	goto L0047
L5c1f:
	g.sub_status()
	g.printS("Egy csatornany\xa1l\xa0s felett \xa0llsz.", nl)
	g.printS("M\xa0kod van, \x82ppen z\xa0rva a fedele.", nl)
	g.sub_6a7f()
	g.gosub = append(g.gosub, 1)
	goto L56b6
R5c3e:
	goto L005f
L5c41:
	g.sub_status()
	g.printS("Elbambult\xa0l,egy cig\xa0ny kotor\xa0szik a", nl)
	g.printS("zsebedben. ", semi)
	if 0 < float64(g.f182) {
		goto L5c6f
	}
	goto L5d07
L5c6f:
	g.f172 = float32(math.Floor(((float64(g.f182) / 2) + 0.5)))
	if 50 < float64(g.f172) {
		goto L5ca2
	}
	goto L5cad
L5ca2:
	g.f172 = float32(50)
L5cad:
	g.color(8)
	g.printS("(Megf\xa3jta", semi)
	g.printF(float64(g.f172), semi)
	g.printS("Ft-odat.)", nl)
	g.color(7)
	g.f182 = float32((float64(g.f182) - float64(g.f172)))
	g.sub_7a66()
	goto L5d32
L5d07:
	g.printS("L\xa0tja,hogy nincs p\x82nzed,d\x81h\x82ben j\xa2l", nl)
	g.printS("megrugdos \x82s kereket old.", nl)
	g.f162 = float32((float64(g.f162) + -12))
	g.gosub = append(g.gosub, 92)
	goto L5633
R5d2d:
	g.sub_7a66()
L5d32:
	goto L0047
}

// sub_title (1010:5d38)
func (g *Game) sub_title() {
	var l0e string
	_ = l0e
	var l12 float32
	_ = l12
	var l16 float32
	_ = l16
	var l1a string
	_ = l1a
	var l1e string
	_ = l1e
	var l20 int
	_ = l20
	var l22 int
	_ = l22
	var l26 string
	_ = l26
	var l28 int
	_ = l28
	var l2c float32
	_ = l2c
	var l30 string
	_ = l30
	var l34 float32
	_ = l34
	var fcarry float64
	g.screen(1, 0)
	g.color(0, 1)
	g.cls(-1)
	l0e = "r30 d25 l20 d15 l10 u39 bm+10,7 r10 d10 l10 u10"
	l12 = float32(0)
L5d89:
	l12 = float32((float64(l12) + 1))
	g.draw(("s=" + varptrF(l12)))
	g.draw("bm 105,0 c3")
	g.draw(("x" + varptrS(l0e)))
	g.draw("bm 130,20 c1")
	g.draw(("x" + varptrS(l0e)))
	g.draw("c0 bm 125,19 l1")
	g.draw("bm 130,20 c3")
	g.draw(("x" + varptrS(l0e)))
	fcarry = 0
	goto L5e42
L5e39:
	fcarry = (float64(l16) + 1)
L5e42:
	l16 = float32(fcarry)
	if 200 >= float64(l16) {
		goto L5e39
	}
	g.draw(("s=" + varptrF(l12)))
	g.draw("bm 105,0 c0")
	g.draw(("x" + varptrS(l0e)))
	g.draw("bm 130,20 c0")
	g.draw(("x" + varptrS(l0e)))
	g.draw("c0 bm 125,19 l1")
	g.draw("bm 130,20 c0")
	g.draw(("x" + varptrS(l0e)))
	if 7 <= float64(l12) {
		goto L5f06
	}
	goto L5d89
L5f06:
	g.draw(("s=" + varptrF(l12)))
	g.draw("bm 105,0 c3")
	g.draw(("x" + varptrS(l0e)))
	g.setPoint(106, 1)
	g.paint(2, 3)
	g.draw("bm 130,20 c1")
	g.draw(("x" + varptrS(l0e)))
	g.draw("c0 bm 125,19 l1")
	g.setPoint(131, 21)
	g.paint(1, 1)
	g.draw("bm 130,20 c3")
	g.draw(("x" + varptrS(l0e)))
	fcarry = 0
	goto L5feb
L5fd8:
	g.printS("", nl)
	fcarry = (float64(l16) + 1)
L5feb:
	l16 = float32(fcarry)
	if 12 >= float64(l16) {
		goto L5fd8
	}
	l1a = "          Szakadt csoves '88."
	g.printS("      Pato Pal Software presents:", nl)
	g.printS("", nl)
	g.printS(l1a, nl)
	g.printS("", nl)
	l1e = "by SZY & KAMARAS"
	l20 = b2i((!(mid(l1e, 4, 1) == "S")))
	l22 = (l20 | b2i((!(mid(l1e, 5, 1) == "Z"))))
	if (l22 | b2i((!(mid(l1e, 6, 1) == "Y")))) != 0 {
		goto L60ba
	}
	goto L6137
L60ba:
	g.screen(0)
	g.cls(-1)
	g.printS("\x8ft\xa1rtad a nevemet,te kutya!", nl)
	g.printS("", nl)
	g.printS("Igy nem tudsz j\xa0tszani.", nl)
	g.printS("", nl)
	g.printS("HA-HA-HA!", nl)
	g.printS("", nl)
	g.printS("", nl)
	g.printS("                 SZY ZOLT\x8fN", nl)
	l26 = g.inputS(2)
	g.end()
L6137:
	g.printS("      (C) 1989.", semi)
	g.printS(l1e, nl)
L6149:
	l26 = g.inkey()
	if l26 == "" {
		goto L6149
	}
	if (asc(l26)) != 13 {
		goto L6149
	}
	g.screen(0, 0)
	g.cls(-1)
	g.locate(omit, omit, 0)
	g.printS("Kell inform\xa0ci\xa2 a j\xa0t\x82kr\xa2l (I/N)? ", semi)
	l26 = g.inputS(1)
	l28 = b2i((!(l26 != "i")))
	if (l28 | b2i((!(l26 != "I")))) != 0 {
		goto L61fe
	}
	goto L62ed
L61fe:
	g.cls(-1)
	g.printS("Te egy szakadt cs\x94ves vagy a Marx t\x82ri", nl)
	g.printS("", nl)
	g.printS("alulj\xa0r\xa2ban.  Az a c\x82lod , hogy  min\x82l", nl)
	g.printS("", nl)
	g.printS("t\x94bb l\xa2v\x82t szerezz. A Sk\xa0l\xa0ban vehetsz", nl)
	g.printS("", nl)
	g.printS("pi\xa0t  \x82s  Technokolt.  J\xa2, ha vigy\xa0zol", nl)
	g.printS("", nl)
	g.printS("magadra, mert olyan h\xa1rek terjengenek,", nl)
	g.printS("", nl)
	g.printS("hogy gy\xa0li cig\xa0nyok j\x94ttek fel Pestre.", nl)
	g.printS("", nl)
	g.printS("De az\x82rt  nem kell mindj\xa0rt berezelni,", nl)
	g.printS("", nl)
	g.printS("a n\xa0lad gyeng\x82bbeket  \x81sd csak le nyu-", nl)
	g.printS("", nl)
	g.printS("godtan. A reflexid\x94d befoly\xa0solhatja a", nl)
	g.printS("", nl)
	g.printS("menek\x81l\x82s  \x82s  a t\xa0mad\xa0s hat\xa0soss\xa0g\xa0t.", nl)
	g.printS("", nl)
	g.printS("Ha gyenge vagy, pi\xa0lj egyet, ha lass\xa3-", nl)
	g.printS("", nl)
	g.printS("nak  \x82rzed  magad , szip\xa2zz  egy  j\xa2t.", nl)
	l26 = g.inputS(1)
L62ed:
	g.cls(-1)
	g.a12e[1] = "Karcsi,a takar\x82kos     (Beginner)"
	g.a12e[2] = "Jancsi,a p\xa2r\x82          (Intermediate)"
	g.a12e[3] = "MacFly,a bandavez\x82r    (Expert)"
	g.color(7, 4)
	g.printS(g.a12e[1], nl)
	g.color(7, 0)
	fcarry = 2
	goto L637d
L635d:
	g.printS(g.a12e[cint(float64(l16))], nl)
	fcarry = (float64(l16) + 1)
L637d:
	l16 = float32(fcarry)
	if 3 >= float64(l16) {
		goto L635d
	}
	g.printS("---------------------------------------", nl)
	l2c = float32(1)
	g.printS("", nl)
	g.printS("", nl)
	g.color(0, 7)
	g.printS(" - ", semi)
	g.color(7, 0)
	g.printS(" .......fel", nl)
	g.printS("", nl)
	g.color(0, 7)
	g.printS(" + ", semi)
	g.color(7, 0)
	g.printS(" .......le", nl)
	g.printS("", nl)
	g.color(0, 7)
	g.printS(" Enter ", semi)
	g.color(7, 0)
	g.printS(" ...k\x82sz", nl)
	g.printS("", nl)
	g.printS("", nl)
	g.printS("A j\xa0t\x82k sor\xa0n mindig ezek a billenty\x81k", nl)
	g.printS("", nl)
	g.printS("funkcion\xa0lnak.A k\xa0romkod\xa0s megjelen\xa1t-", nl)
	g.printS("", nl)
	g.printS("het\x94 a ", semi)
	g.color(0, 7)
	g.printS(" Space ", semi)
	g.color(7, 0)
	g.printS(" billenty\x81vel.", nl)
	g.printS("", nl)
	g.printS("", nl)
	g.printS("V\xa0lassz,melyikkel leszel!", nl)
L652e:
	l26 = g.inputS(1)
	l30 = l26
	if chr(13) == l30 {
		goto L656a
	}
	goto L656d
L656a:
	goto L6725
L656d:
	if l30 == "+" {
		goto L6580
	}
	goto L6647
L6580:
	if 3 > float64(l2c) {
		goto L6594
	}
	goto L6644
L6594:
	g.color(7, 0)
	g.locate(cint(float64(l2c)), 1)
	g.printS(g.a12e[cint(float64(l2c))], nl)
	g.color(7, 4)
	g.locate(cint((float64(l2c) + 1)), 1)
	g.printS(g.a12e[cint((float64(l2c)+1))], nl)
	l2c = float32((float64(l2c) + 1))
	goto L652e
L6644:
	goto L6721
L6647:
	if l30 == "-" {
		goto L665a
	}
	goto L6721
L665a:
	if 1 < float64(l2c) {
		goto L666e
	}
	goto L671e
L666e:
	g.color(7, 0)
	g.locate(cint(float64(l2c)), 1)
	g.printS(g.a12e[cint(float64(l2c))], nl)
	g.color(7, 4)
	g.locate(cint((float64(l2c) + -1)), 1)
	g.printS(g.a12e[cint((float64(l2c)+-1))], nl)
	l2c = float32((float64(l2c) + -1))
	goto L652e
L671e:
	goto L6721
L6721:
	goto L652e
L6725:
	l34 = float32(float64(l2c))
	if float64(l34) == 1 {
		goto L6744
	}
	goto L67a7
L6744:
	g.f18e = float32(20)
	g.f162 = float32(100)
	g.f1b2 = float32(0)
	g.f1b6 = float32(1)
	g.f1e6 = float32(0)
	g.f192 = float32(0)
	g.f1a6 = float32(15)
	g.f182 = float32(g.tune("startMoney", 55))
	goto L6896
L67a7:
	if float64(l34) == 2 {
		goto L67bc
	}
	goto L681f
L67bc:
	g.f18e = float32(20)
	g.f162 = float32(80)
	g.f1b2 = float32(1)
	g.f1b6 = float32(0)
	g.f1e6 = float32(0)
	g.f192 = float32(0)
	g.f1a6 = float32(18)
	g.f182 = float32(0)
	goto L6896
L681f:
	if float64(l34) == 3 {
		goto L6834
	}
	goto L6896
L6834:
	g.f18e = float32(20)
	g.f162 = float32(60)
	g.f1b2 = float32(0)
	g.f1b6 = float32(0)
	g.f1e6 = float32(1)
	g.f192 = float32(1)
	g.f1a6 = float32(22)
	g.f182 = float32(0)
	goto L6896
L6896:
	g.color(7, 0)
	g.cls(-1)
	g.f1ae = float32(float64(l2c))
	return
}

// sub_initart (1010:68fa)
func (g *Game) sub_initart() {
	g.a0da[1] = "fffffflllf      ff  lll   ffll"
	g.a0da[2] = "l    f   fffffff  ff   fff   l"
	g.a0da[3] = "l                            f"
	g.a0da[4] = "ffff    oo   oo   oo        ff"
	g.a0da[5] = "    ff  oo   oo   oo      ffii"
	g.a0da[6] = "     f                  u m   f"
	g.a0da[7] = "     f       u            m   f"
	g.a0da[8] = "     f    c               f   f"
	g.a0da[9] = "     f               c    ffii"
	g.a0da[10] = " ffff u oo   oo   oo      uff "
	g.a0da[11] = "fs c    oo   oo   oo        cf"
	g.a0da[12] = "fs         c         c oo    f"
	g.a0da[13] = "fs                     oo   \x94f"
	g.a0da[14] = "fs      oo   oo   oo         f"
	g.a0da[15] = "fs         c          \x94   ffff"
	g.a0da[16] = " f u              \x94        f  "
	g.a0da[17] = "  f   fff    fff  ffff      f "
	g.a0da[18] = "  f cf   f   f f  f   ff     f"
	g.a0da[19] = "  f  f   f   f f  f     f    f"
	g.a0da[20] = "  fllf   flllf fllf      llll "
	g.randomize(g.timer())
	g.f1f2 = float32(10)
	g.f14a = float32(g.tune("startCol", 4))
	g.f146 = float32(g.tune("startRow", 19))
	g.f16a = float32(0)
	g.f16e = float32(0)
	g.s156 = " fousmilc\x94"
	g.f1c2 = float32(0)
	return
}

// sub_6a7f (1010:6a7f)
func (g *Game) sub_6a7f() {
	var l0e float32
	_ = l0e
	var l12 float32
	_ = l12
	var l16 string
	_ = l16
	var l1a string
	_ = l1a
	var l1e float32
	_ = l1e
	var fcarry float64
	g.f19e = float32(0)
	g.f186 = float32(0)
	g.f196 = float32(0)
	g.f17a = float32(0)
	g.sub_6ee3()
	l0e = float32(1)
	g.a036[1] = "Megyek \x82szakra"
	g.a036[2] = "Megyek d\x82lre"
	g.a036[3] = "Megyek keletre"
	g.a036[4] = "Megyek nyugatra"
	g.a036[5] = "Menek\x81l\x94k"
	g.a036[6] = "Menek\x81l\x94k,de m\x82g visszaki\xa0ltok"
	g.a036[7] = "Veszek pi\xa0t"
	g.a036[8] = "Veszek Technokolt"
	g.a036[9] = "Lejmolok"
	g.a036[10] = "K\xa0romkodok"
	g.a036[11] = "Megt\xa0madom"
	g.a036[12] = "V\xa0rok"
	g.a036[13] = "Egy\x82b"
	g.color(7, 4)
	g.printS(g.a036[1], nl)
	g.color(7, 0)
	fcarry = 2
	goto L6bd1
L6bb1:
	g.printS(g.a036[cint(float64(l12))], nl)
	fcarry = (float64(l12) + 1)
L6bd1:
	l12 = float32(fcarry)
	if 13 >= float64(l12) {
		goto L6bb1
	}
	g.locate(omit, omit, 0)
L6bfc:
	l1a = g.inputS(1)
	l16 = l1a
	if chr(13) == l16 {
		goto L6c38
	}
	goto L6c91
L6c38:
	g.locate(cint(float64(l0e)), 1)
	g.color(7, 0)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.f17a = float32(float64(l0e))
	goto L6e49
L6c91:
	if l16 == "+" {
		goto L6ca4
	}
	goto L6d6b
L6ca4:
	if 13 > float64(l0e) {
		goto L6cb8
	}
	goto L6d68
L6cb8:
	g.color(7, 0)
	g.locate(cint(float64(l0e)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 4)
	g.locate(cint((float64(l0e) + 1)), 1)
	g.printS(g.a036[cint((float64(l0e)+1))], nl)
	l0e = float32((float64(l0e) + 1))
	goto L6bfc
L6d68:
	goto L6e45
L6d6b:
	if l16 == "-" {
		goto L6d7e
	}
	goto L6e45
L6d7e:
	if 1 < float64(l0e) {
		goto L6d92
	}
	goto L6e42
L6d92:
	g.color(7, 0)
	g.locate(cint(float64(l0e)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 4)
	g.locate(cint((float64(l0e) + -1)), 1)
	g.printS(g.a036[cint((float64(l0e)+-1))], nl)
	l0e = float32((float64(l0e) + -1))
	goto L6bfc
L6e42:
	goto L6e45
L6e45:
	goto L6bfc
L6e49:
	l1e = float32(float64(l0e))
	if float64(l1e) == 13 {
		goto L6e67
	}
	goto L6e70
L6e67:
	g.sub_6f69()
	goto L6ec6
L6e70:
	if float64(l1e) == 9 {
		goto L6e84
	}
	goto L6e8d
L6e84:
	g.sub_735d()
	goto L6ec6
L6e8d:
	if float64(l1e) == 10 {
		goto L6ea1
	}
	goto L6eaa
L6ea1:
	g.sub_7fc2()
	goto L6ec6
L6eaa:
	if float64(l1e) == 11 {
		goto L6ebe
	}
	goto L6ec6
L6ebe:
	g.sub_76e2()
	goto L6ec6
L6ec6:
	return
}

// sub_6ee3 (1010:6ee3)
func (g *Game) sub_6ee3() {
	var l0e float32
	_ = l0e
	g.color(0, 0, 0)
	g.locate(1, 1)
	l0e = float32(0)
L6f24:
	g.printS("                                        ", semi)
	l0e = float32((float64(l0e) + 1))
	if 13 > float64(l0e) {
		goto L6f24
	}
	g.locate(1, 1)
	return
}

// sub_6f69 (1010:6f69)
func (g *Game) sub_6f69() {
	var l0e float32
	_ = l0e
	var l12 float32
	_ = l12
	var l16 string
	_ = l16
	var l1a string
	_ = l1a
	var fcarry float64
	g.sub_6ee3()
	l0e = float32(1)
	g.a036[1] = "Szip\xa2zok egy j\xa2t"
	g.a036[2] = "Pi\xa0lok egy j\xa2t"
	g.a036[3] = "Meger\x94szakolom"
	g.a036[4] = "Bet\x94r\x94k"
	g.a036[5] = "H\xa0t\xa0ba ker\x81l\x94k \x82s ellopom a t\xa0rc\xa0j\xa0t"
	g.a036[6] = "\x99ngyilkos leszek"
	g.color(7, 4)
	g.printS(g.a036[1], nl)
	g.color(7, 0)
	fcarry = 2
	goto L702f
L700f:
	g.printS(g.a036[cint(float64(l12))], nl)
	fcarry = (float64(l12) + 1)
L702f:
	l12 = float32(fcarry)
	if 6 >= float64(l12) {
		goto L700f
	}
	g.locate(omit, omit, 0)
L705a:
	l1a = g.inputS(1)
	l16 = l1a
	if chr(13) == l16 {
		goto L7096
	}
	goto L7189
L7096:
	g.locate(cint(float64(l0e)), 1)
	g.color(7, 0)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.f17a = float32((float64(l0e) + 13))
	if 6 == float64(l0e) {
		goto L7106
	}
	goto L7183
L7106:
	g.sub_status()
	g.color(12)
	g.printS("Ugy d\x94ntesz,te ezt nem b\xa1rod tov\xa0bb.", nl)
	g.printS("", nl)
	g.printS("Beveszel egy doboz Kalopi\xa2t.", nl)
	g.locate(25, 1)
	g.color(31)
	g.printS("V\x82ged!", semi)
	l1a = g.inputS(1)
	g.end()
L7183:
	goto L7340
L7189:
	if l16 == "+" {
		goto L719c
	}
	goto L7263
L719c:
	if 6 > float64(l0e) {
		goto L71b0
	}
	goto L7260
L71b0:
	g.color(7, 0)
	g.locate(cint(float64(l0e)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 4)
	g.locate(cint((float64(l0e) + 1)), 1)
	g.printS(g.a036[cint((float64(l0e)+1))], nl)
	l0e = float32((float64(l0e) + 1))
	goto L705a
L7260:
	goto L733d
L7263:
	if l16 == "-" {
		goto L7276
	}
	goto L733d
L7276:
	if 1 < float64(l0e) {
		goto L728a
	}
	goto L733a
L728a:
	g.color(7, 0)
	g.locate(cint(float64(l0e)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 4)
	g.locate(cint((float64(l0e) + -1)), 1)
	g.printS(g.a036[cint((float64(l0e)+-1))], nl)
	l0e = float32((float64(l0e) + -1))
	goto L705a
L733a:
	goto L733d
L733d:
	goto L705a
L7340:
	return
}

// sub_735d (1010:735d)
func (g *Game) sub_735d() {
	var l0e float32
	_ = l0e
	var l12 float32
	_ = l12
	var l16 string
	_ = l16
	var l1a string
	_ = l1a
	var fcarry float64
	g.sub_6ee3()
	l0e = float32(1)
	g.a036[1] = "K\x82rek sz\x82pen 10 forintot,gy\xa2gyszerre."
	g.a036[2] = "Ki tudn\xa0l p\xa0r forinttal seg\xa1teni?"
	g.a036[3] = "B\xa0csi,vonatjegyre kellene 15 Ft."
	g.a036[4] = "N\x82ni,tudna p\xa0r Ft-ot adni,\x82hes vagyok."
	g.a036[5] = "Ide a l\xa2v\x82t,hapsik\xa0m,mert k\x81l\x94nben ..."
	g.a036[6] = "B\xa0csi,az ap\xa0d fasz\xa0t,ide a p\x82nzt!"
	g.a036[7] = "Ide a stexet,te v\x82n loty\xa2!"
	g.a036[8] = "Ide a doh\xa0nyt,te k\x94cs\x94g!"
	g.a036[9] = "Ide a gub\xa0t,te ribanc!"
	g.color(7, 4)
	g.printS(g.a036[1], nl)
	g.color(7, 0)
	fcarry = 2
	goto L744b
L742b:
	g.printS(g.a036[cint(float64(l12))], nl)
	fcarry = (float64(l12) + 1)
L744b:
	l12 = float32(fcarry)
	if 9 >= float64(l12) {
		goto L742b
	}
	g.locate(omit, omit, 0)
L7476:
	l1a = g.inputS(1)
	l16 = l1a
	if chr(13) == l16 {
		goto L74b2
	}
	goto L750e
L74b2:
	g.locate(cint(float64(l0e)), 1)
	g.color(7, 0)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.f196 = float32(float64(l0e))
	goto L76c5
L750e:
	if l16 == "+" {
		goto L7521
	}
	goto L75e8
L7521:
	if 9 > float64(l0e) {
		goto L7535
	}
	goto L75e5
L7535:
	g.color(7, 0)
	g.locate(cint(float64(l0e)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 4)
	g.locate(cint((float64(l0e) + 1)), 1)
	g.printS(g.a036[cint((float64(l0e)+1))], nl)
	l0e = float32((float64(l0e) + 1))
	goto L7476
L75e5:
	goto L76c2
L75e8:
	if l16 == "-" {
		goto L75fb
	}
	goto L76c2
L75fb:
	if 1 < float64(l0e) {
		goto L760f
	}
	goto L76bf
L760f:
	g.color(7, 0)
	g.locate(cint(float64(l0e)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 4)
	g.locate(cint((float64(l0e) + -1)), 1)
	g.printS(g.a036[cint((float64(l0e)+-1))], nl)
	l0e = float32((float64(l0e) + -1))
	goto L7476
L76bf:
	goto L76c2
L76c2:
	goto L7476
L76c5:
	return
}

// sub_76e2 (1010:76e2)
func (g *Game) sub_76e2() {
	var l0e float32
	_ = l0e
	var l12 float32
	_ = l12
	var l16 string
	_ = l16
	var l1a string
	_ = l1a
	var fcarry float64
	g.sub_6ee3()
	l0e = float32(1)
	g.a036[1] = "T\x94k\x94nverem"
	g.a036[2] = "Lefejelem"
	g.a036[3] = "Nyakonv\xa0gom"
	g.a036[4] = "Orrbaverem"
	g.a036[5] = "Hasbarugom"
	g.a036[6] = "T\x94k\x94nrugom"
	g.a036[7] = "Leteperem \x82s megt\x82p\xa0zom"
	g.a036[8] = "L\xa0nccal r\xa0suhintok"
	g.a036[9] = "Megk\x82selem"
	g.color(7, 4)
	g.printS(g.a036[1], nl)
	g.color(7, 0)
	fcarry = 2
	goto L77cf
L77af:
	g.printS(g.a036[cint(float64(l12))], nl)
	fcarry = (float64(l12) + 1)
L77cf:
	l12 = float32(fcarry)
	if 9 >= float64(l12) {
		goto L77af
	}
	g.locate(omit, omit, 0)
L77fa:
	l1a = g.inputS(1)
	l16 = l1a
	if chr(13) == l16 {
		goto L7836
	}
	goto L7892
L7836:
	g.locate(cint(float64(l0e)), 1)
	g.color(7, 0)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.f186 = float32(float64(l0e))
	goto L7a49
L7892:
	if l16 == "+" {
		goto L78a5
	}
	goto L796c
L78a5:
	if 9 > float64(l0e) {
		goto L78b9
	}
	goto L7969
L78b9:
	g.color(7, 0)
	g.locate(cint(float64(l0e)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 4)
	g.locate(cint((float64(l0e) + 1)), 1)
	g.printS(g.a036[cint((float64(l0e)+1))], nl)
	l0e = float32((float64(l0e) + 1))
	goto L77fa
L7969:
	goto L7a46
L796c:
	if l16 == "-" {
		goto L797f
	}
	goto L7a46
L797f:
	if 1 < float64(l0e) {
		goto L7993
	}
	goto L7a43
L7993:
	g.color(7, 0)
	g.locate(cint(float64(l0e)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 4)
	g.locate(cint((float64(l0e) + -1)), 1)
	g.printS(g.a036[cint((float64(l0e)+-1))], nl)
	l0e = float32((float64(l0e) + -1))
	goto L77fa
L7a43:
	goto L7a46
L7a46:
	goto L77fa
L7a49:
	return
}

// sub_7a66 (1010:7a66)
func (g *Game) sub_7a66() {
	var l0e string
	_ = l0e
	if 24 == float64(g.f1a6) {
		goto L7a85
	}
	goto L7a91
L7a85:
	g.f1a6 = float32(0)
L7a91:
	if 100 == float64(g.f162) {
		goto L7aa7
	}
	goto L7ab3
L7aa7:
	g.f18e = float32(20)
L7ab3:
	if 10 > float64(g.f162) {
		goto L7ac9
	}
	goto L7add
L7ac9:
	g.f18e = float32((float64(g.f18e) + 25))
	goto L7b5a
L7add:
	if 20 > float64(g.f162) {
		goto L7af3
	}
	goto L7b07
L7af3:
	g.f18e = float32((float64(g.f18e) + 10))
	goto L7b5a
L7b07:
	if 30 > float64(g.f162) {
		goto L7b1d
	}
	goto L7b31
L7b1d:
	g.f18e = float32((float64(g.f18e) + 5))
	goto L7b5a
L7b31:
	if 50 > float64(g.f162) {
		goto L7b47
	}
	goto L7b5a
L7b47:
	g.f18e = float32((float64(g.f18e) + 1))
	goto L7b5a
L7b5a:
	g.locate(14, 1)
	g.color(15)
	g.printS(stringFill(40, 196), nl)
	g.color(7)
	g.locate(15, 1)
	g.color(7, 0)
	if 50 > float64(g.f162) {
		goto L7be9
	}
	goto L7bfb
L7be9:
	g.color(4)
L7bfb:
	if 20 > float64(g.f162) {
		goto L7c10
	}
	goto L7c21
L7c10:
	g.color(20)
L7c21:
	g.printS("Energia  :", semi)
	g.printF(float64(g.f162), semi)
	g.printS("%  ", tab)
	g.color(7, 0)
	if 20 < float64(g.f18e) {
		goto L7c6e
	}
	goto L7c7f
L7c6e:
	g.color(5)
L7c7f:
	g.printS("Reflexid\x94:", semi)
	g.printF((float64(g.f18e) / 100), semi)
	g.printS("        ", nl)
	g.color(7)
	g.printS("P\x82nzed   :", semi)
	g.printF(float64(g.f182), semi)
	g.printS("Ft    ", nl)
	g.printS("Technokol:", semi)
	g.printF(float64(g.f1b6), semi)
	g.printS("        Pia:", semi)
	g.printF(float64(g.f1b2), semi)
	g.printS("  ", nl)
	g.color(15)
	g.printS(stringFill(40, 196), nl)
	g.color(7)
	g.locate(15, 22)
	g.printS("Pontos id\x94:", semi)
	g.printF(float64(g.f1a6), semi)
	g.printS("\xa2ra ", nl)
	if 0 < float64(g.f192) {
		goto L7d92
	}
	goto L7dd7
L7d92:
	g.locate(16, 22)
	g.color(20)
	g.printS("K\x94r\x94z a rend\x94rs\x82g!", nl)
	g.color(7)
L7dd7:
	if 1 == float64(g.f1e6) {
		goto L7dec
	}
	goto L7e33
L7dec:
	g.locate(17, 22)
	g.color(8)
	g.printS("Cig\xa0nyvesz\x82ly!", nl)
	g.color(7)
	goto L7e56
L7e33:
	g.locate(17, 22)
	g.printS("              ", nl)
L7e56:
	if 250 <= float64(g.f182) {
		goto L7e6b
	}
	goto L7fae
L7e6b:
	l0e = g.inputS(1)
	g.cls(-1)
	g.color(12)
	g.printS("CONGRATULATIONS!", nl)
	g.printS("", nl)
	g.color(7)
	g.printF(float64(g.f182), semi)
	g.printS(" Ft-ot gy\x81jt\x94tt\x82l be.", nl)
	g.printS("", nl)
	g.printS("Ezzel  az  \x94sszeggel  m\xa0r  elmehetsz", nl)
	g.printS("", nl)
	g.printS("j\xa0t\x82ktermezni \x82s j\xa2l kikup\xa0l\xa2dhatsz.", nl)
	g.printS("", nl)
	g.printS("De ez m\xa0r egy m\xa0sik t\x94rt\x82net.", nl)
	g.printS("", nl)
	g.printS("", nl)
	g.printS("A Pat\xa2 P\xa0l Software k\x94sz\x94ni t\x81relmedet.", nl)
	g.printS("", nl)
	g.printS("", nl)
	g.printS("", nl)
	g.color(0, 7)
	g.printS(" Space ", semi)
	g.color(7, 0)
	g.printS(" ... v\x82ge a j\xa0t\x82knak.", nl)
L7f83:
	l0e = g.inputS(1)
	if l0e != " " {
		goto L7f83
	}
	g.end()
L7fae:
	return
}

// sub_7fc2 (1010:7fc2)
func (g *Game) sub_7fc2() {
	var l0e float32
	_ = l0e
	var l12 float32
	_ = l12
	var l16 string
	_ = l16
	var l1a string
	_ = l1a
	var l1e float32
	_ = l1e
	var l22 float32
	_ = l22
	var l26 float32
	_ = l26
	var l2a float32
	_ = l2a
	var l2e string
	_ = l2e
	var l32 float32
	_ = l32
	var l34 int
	_ = l34
	var l36 int
	_ = l36
	var fcarry float64
	goto L7fc2
gosubReturn:
	switch g.gosubPop() {
	case 93:
		goto R83ed
	}
	panic("bad gosub return")
L7fc2:
	g.f19e = float32(0)
	g.s1f6 = ""
L7fe3:
	g.sub_6ee3()
	l0e = float32(1)
	l12 = float32(1)
	g.a036[1] = "(Megsz\xa2l\xa1t\xa0sok)"
	g.a036[2] = "a terembur\xa0j\xa0t "
	g.a036[3] = "a mindens\x82git "
	g.a036[4] = "te "
	g.a036[5] = "azt a 7 meg a 8-\xa0t "
	g.a036[6] = "ejnye-bejnye "
	g.a036[7] = "a kisk\x82sit neki "
	g.a036[8] = "1/2 esz\x81 bolond "
	g.a036[9] = "nyavaly\xa0s "
	g.a036[10] = "rohadt "
	g.a036[11] = "\xa0llat "
	g.a036[12] = "barom "
	g.a036[13] = "diszn\xa2 "
	g.a036[14] = "kurva "
	g.a036[15] = "r\x82szeg "
	g.a036[16] = "szem\x82t "
	g.a036[17] = "a kurva any\xa0d "
	g.a036[18] = "fasz "
	g.a036[19] = "l\xa2fasz a seggedbe "
	g.a036[20] = "nyald ki "
	g.a036[21] = "kapd be "
	g.a036[22] = "sz\xa1vj meg "
	g.a036[23] = "h\x81lye "
	g.a036[24] = "az ap\xa0d fasz\xa0t "
	g.a036[25] = "gan\x82j "
	g.a036[26] = "b\x81d\x94s "
	g.a036[27] = "patk\xa0ny "
	goto L866c
L8174:
	l1a = g.inputS(1)
	l16 = l1a
	if l16 == "," {
		goto L81aa
	}
	goto L81e8
L81aa:
	if right(g.s1f6, 1) == " " {
		goto L81c6
	}
	goto L81e5
L81c6:
	g.s1f6 = midStmt(g.s1f6, len(g.s1f6), 1, ",")
L81e5:
	goto L8668
L81e8:
	if chr(13) == l16 {
		goto L8201
	}
	goto L8348
L8201:
	l1e = float32(0)
	if 1 == float64(l0e) {
		goto L8220
	}
	goto L8229
L8220:
	g.sub_8841()
	goto L7fe3
L8229:
	if (b2i((!(8 <= float64(l0e)))) & b2i((!(2 <= float64(g.f19e))))) != 0 {
		goto L8261
	}
	goto L8270
L8261:
	g.f19e = float32(1)
	goto L827b
L8270:
	g.f19e = float32(2)
L827b:
	g.color(1, 4)
	g.locate(cint(float64(l12)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	if g.s1f6 == "" {
		goto L82d1
	}
	goto L82db
L82d1:
	l1e = float32(1)
L82db:
	g.s1f6 = (g.s1f6 + g.a036[cint(float64(l0e))])
	if 1 == float64(l1e) {
		goto L8315
	}
	goto L8345
L8315:
	g.s1f6 = midStmt(g.s1f6, 1, 1, chr((asc(mid(g.s1f6, 1, 1)) + -32)))
L8345:
	goto L8668
L8348:
	if l16 == " " {
		goto L835a
	}
	goto L840a
L835a:
	if g.s1f6 != "" {
		goto L836c
	}
	goto L7fe3
L836c:
	g.color(3, 0)
	g.locate(20, 1)
	g.printS("                                        ", semi)
	g.printS("                                        ", semi)
	g.printS("                                        ", semi)
	g.locate(20, 1)
	g.s1f6 = midStmt(g.s1f6, len(g.s1f6), 1, "!")
	g.gosub = append(g.gosub, 93)
	goto L8720
R83ed:
	l1a = g.inputS(1)
	goto L881b
L840a:
	if l16 == "+" {
		goto L841d
	}
	goto L8539
L841d:
	if 9 == float64(l0e) {
		goto L8432
	}
	goto L8445
L8432:
	l0e = float32(10)
	g.sub_6ee3()
	goto L866c
L8445:
	if 18 == float64(l0e) {
		goto L845a
	}
	goto L846d
L845a:
	l0e = float32(19)
	g.sub_6ee3()
	goto L866c
L846d:
	if 27 > float64(l0e) {
		goto L8481
	}
	goto L8536
L8481:
	g.color(7, 0)
	g.locate(cint(float64(l12)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 4)
	l0e = float32((float64(l0e) + 1))
	l12 = float32((float64(l12) + 1))
	g.locate(cint(float64(l12)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	goto L8174
L8536:
	goto L8668
L8539:
	if l16 == "-" {
		goto L854c
	}
	goto L8668
L854c:
	if 10 == float64(l0e) {
		goto L8561
	}
	goto L8574
L8561:
	l0e = float32(1)
	g.sub_6ee3()
	goto L866c
L8574:
	if 19 == float64(l0e) {
		goto L8589
	}
	goto L859c
L8589:
	l0e = float32(10)
	g.sub_6ee3()
	goto L866c
L859c:
	if 1 < float64(l0e) {
		goto L85b0
	}
	goto L8665
L85b0:
	g.color(7, 0)
	g.locate(cint(float64(l12)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 4)
	l0e = float32((float64(l0e) + -1))
	l12 = float32((float64(l12) + -1))
	g.locate(cint(float64(l12)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	goto L8174
L8665:
	goto L8668
L8668:
	goto L8174
L866c:
	l12 = float32(1)
	g.locate(1, 1)
	g.color(7, 4)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 0)
	l22 = float32((float64(l0e) + 8))
	fcarry = (float64(l0e) + 1)
	goto L8707
L86e7:
	g.printS(g.a036[cint(float64(l26))], nl)
	fcarry = (float64(l26) + 1)
L8707:
	l26 = float32(fcarry)
	if float64(l22) >= float64(l26) {
		goto L86e7
	}
	goto L8174
L8720:
	l2e = ""
	l32 = float32(0)
	l2a = float32(float64(len(g.s1f6)))
	fcarry = 1
	goto L8801
L8755:
	l32 = float32((float64(l32) + 1))
	l1a = mid(g.s1f6, cint(float64(l26)), 1)
	l34 = b2i((!(35 >= float64(l32))))
	l36 = b2i((!(l1a != " ")))
	if (l34 & (l36 | b2i((!(l1a != ","))))) != 0 {
		goto L87d7
	}
	goto L87ee
L87d7:
	g.printS(l1a, nl)
	l32 = float32(0)
	goto L87f8
L87ee:
	g.printS(l1a, semi)
L87f8:
	fcarry = (float64(l26) + 1)
L8801:
	l26 = float32(fcarry)
	if float64(l2a) < float64(l26) {
		goto L881a
	}
	goto L8755
L881a:
	goto gosubReturn
L881b:
	return
}

// sub_8841 (1010:8841)
func (g *Game) sub_8841() {
	var l0e float32
	_ = l0e
	var l12 float32
	_ = l12
	var l16 string
	_ = l16
	var l1a string
	_ = l1a
	var l1e float32
	_ = l1e
	var fcarry float64
	g.sub_6ee3()
	l0e = float32(1)
	g.a036[1] = "B\xa0csi "
	g.a036[2] = "N\x82ni "
	g.a036[3] = "Bar\xa0tom "
	g.a036[4] = "Haver "
	g.a036[5] = "Nigger "
	g.a036[6] = "Geci "
	g.a036[7] = "Cig\xa0ny "
	g.a036[8] = "Paraszt "
	g.a036[9] = "Zsid\xa2 "
	g.a036[10] = "Cs\xa0v\xa2k\xa0m "
	g.a036[11] = "Hugyos J\xa2zsi "
	g.a036[12] = "Fakab\xa0t "
	g.color(7, 4)
	g.printS(g.a036[1], nl)
	g.color(7, 0)
	fcarry = 2
	goto L8955
L8935:
	g.printS(g.a036[cint(float64(l12))], nl)
	fcarry = (float64(l12) + 1)
L8955:
	l12 = float32(fcarry)
	if 12 >= float64(l12) {
		goto L8935
	}
	g.locate(omit, omit, 0)
L8980:
	l1a = g.inputS(1)
	l16 = l1a
	if chr(13) == l16 {
		goto L89bd
	}
	goto L8a90
L89bd:
	if (b2i((!(5 <= float64(l0e)))) & b2i((!(2 <= float64(l1e))))) != 0 {
		goto L89f4
	}
	goto L8a02
L89f4:
	l1e = float32(1)
	goto L8a0c
L8a02:
	l1e = float32(2)
L8a0c:
	if g.s1f6 == "" {
		goto L8a1f
	}
	goto L8a65
L8a1f:
	g.s1f6 = (left(g.a036[cint(float64(l0e))], (len(g.a036[cint(float64(l0e))])-1)) + ",")
	goto L8c47
L8a65:
	g.s1f6 = (g.s1f6 + g.a036[cint(float64(l0e))])
	goto L8c47
L8a90:
	if l16 == "+" {
		goto L8aa3
	}
	goto L8b6a
L8aa3:
	if 12 > float64(l0e) {
		goto L8ab7
	}
	goto L8b67
L8ab7:
	g.color(7, 0)
	g.locate(cint(float64(l0e)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 4)
	g.locate(cint((float64(l0e) + 1)), 1)
	g.printS(g.a036[cint((float64(l0e)+1))], nl)
	l0e = float32((float64(l0e) + 1))
	goto L8980
L8b67:
	goto L8c44
L8b6a:
	if l16 == "-" {
		goto L8b7d
	}
	goto L8c44
L8b7d:
	if 1 < float64(l0e) {
		goto L8b91
	}
	goto L8c41
L8b91:
	g.color(7, 0)
	g.locate(cint(float64(l0e)), 1)
	g.printS(g.a036[cint(float64(l0e))], nl)
	g.color(7, 4)
	g.locate(cint((float64(l0e) + -1)), 1)
	g.printS(g.a036[cint((float64(l0e)+-1))], nl)
	l0e = float32((float64(l0e) + -1))
	goto L8980
L8c41:
	goto L8c44
L8c44:
	goto L8980
L8c47:
	return
}

// sub_status (1010:8c64)
func (g *Game) sub_status() {
	g.color(7, 0)
	g.locate(20, 1)
	g.printS("                                        ", semi)
	g.printS("                                        ", semi)
	g.printS("                                        ", semi)
	g.printS("                                        ", semi)
	g.locate(20, 1)
	return
}

// sub_8cdf (1010:8cdf)
func (g *Game) sub_8cdf() {
	g.a0da[cint(float64(g.f146))] = midStmt(g.a0da[cint(float64(g.f146))], cint(float64(g.f14a)), 1, " ")
	return
}
