package amock

func ItReturns0x1[Out1 any](fnPt *func() Out1, out1 Out1) (e error) {
	e = ItReturns(fnPt, out1)
	return
}

func ItReturns0x2[Out1 any,
	Out2 any](fnPt *func() (Out1, Out2), out1 Out1, out2 Out2) (e error) {
	e = ItReturns(fnPt, out1, out2)
	return
}

func ItReturns0x3[Out1 any,
	Out2 any,
	Out3 any](fnPt *func() (Out1, Out2, Out3), out1 Out1, out2 Out2, out3 Out3) (e error) {
	e = ItReturns(fnPt, out1, out2, out3)
	return
}

func ItReturns0x4[Out1 any,
	Out2 any,
	Out3 any,
	Out4 any](fnPt *func() (Out1, Out2, Out3, Out4), out1 Out1, out2 Out2, out3 Out3, out4 Out4) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4)
	return
}

func ItReturns0x5[Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any](fnPt *func() (Out1, Out2, Out3, Out4, Out5), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5)
	return
}

func ItReturns0x6[Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any](fnPt *func() (Out1, Out2, Out3, Out4, Out5, Out6), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6)
	return
}

func ItReturns0x7[Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any](fnPt *func() (Out1, Out2, Out3, Out4, Out5, Out6, Out7), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7)
	return
}

func ItReturns0x8[Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any](fnPt *func() (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8)
	return
}

func ItReturns0x9[Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any,
	Out9 any](fnPt *func() (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8, out9 Out9) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8, out9)
	return
}

func ItReturns1x1[In1 any,
	Out1 any](fnPt *func(In1) Out1, out1 Out1) (e error) {
	e = ItReturns(fnPt, out1)
	return
}

func ItReturns1x2[In1 any,
	Out1 any,
	Out2 any](fnPt *func(In1) (Out1, Out2), out1 Out1, out2 Out2) (e error) {
	e = ItReturns(fnPt, out1, out2)
	return
}

func ItReturns1x3[In1 any,
	Out1 any,
	Out2 any,
	Out3 any](fnPt *func(In1) (Out1, Out2, Out3), out1 Out1, out2 Out2, out3 Out3) (e error) {
	e = ItReturns(fnPt, out1, out2, out3)
	return
}

func ItReturns1x4[In1 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any](fnPt *func(In1) (Out1, Out2, Out3, Out4), out1 Out1, out2 Out2, out3 Out3, out4 Out4) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4)
	return
}

func ItReturns1x5[In1 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any](fnPt *func(In1) (Out1, Out2, Out3, Out4, Out5), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5)
	return
}

func ItReturns1x6[In1 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any](fnPt *func(In1) (Out1, Out2, Out3, Out4, Out5, Out6), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6)
	return
}

func ItReturns1x7[In1 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any](fnPt *func(In1) (Out1, Out2, Out3, Out4, Out5, Out6, Out7), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7)
	return
}

func ItReturns1x8[In1 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any](fnPt *func(In1) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8)
	return
}

func ItReturns1x9[In1 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any,
	Out9 any](fnPt *func(In1) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8, out9 Out9) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8, out9)
	return
}

func ItReturns2x1[In1 any,
	In2 any,
	Out1 any](fnPt *func(In1, In2) Out1, out1 Out1) (e error) {
	e = ItReturns(fnPt, out1)
	return
}

func ItReturns2x2[In1 any,
	In2 any,
	Out1 any,
	Out2 any](fnPt *func(In1, In2) (Out1, Out2), out1 Out1, out2 Out2) (e error) {
	e = ItReturns(fnPt, out1, out2)
	return
}

func ItReturns2x3[In1 any,
	In2 any,
	Out1 any,
	Out2 any,
	Out3 any](fnPt *func(In1, In2) (Out1, Out2, Out3), out1 Out1, out2 Out2, out3 Out3) (e error) {
	e = ItReturns(fnPt, out1, out2, out3)
	return
}

func ItReturns2x4[In1 any,
	In2 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any](fnPt *func(In1, In2) (Out1, Out2, Out3, Out4), out1 Out1, out2 Out2, out3 Out3, out4 Out4) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4)
	return
}

func ItReturns2x5[In1 any,
	In2 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any](fnPt *func(In1, In2) (Out1, Out2, Out3, Out4, Out5), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5)
	return
}

func ItReturns2x6[In1 any,
	In2 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any](fnPt *func(In1, In2) (Out1, Out2, Out3, Out4, Out5, Out6), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6)
	return
}

func ItReturns2x7[In1 any,
	In2 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any](fnPt *func(In1, In2) (Out1, Out2, Out3, Out4, Out5, Out6, Out7), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7)
	return
}

func ItReturns2x8[In1 any,
	In2 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any](fnPt *func(In1, In2) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8)
	return
}

func ItReturns2x9[In1 any,
	In2 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any,
	Out9 any](fnPt *func(In1, In2) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8, out9 Out9) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8, out9)
	return
}

func ItReturns3x1[In1 any,
	In2 any,
	In3 any,
	Out1 any](fnPt *func(In1, In2, In3) Out1, out1 Out1) (e error) {
	e = ItReturns(fnPt, out1)
	return
}

func ItReturns3x2[In1 any,
	In2 any,
	In3 any,
	Out1 any,
	Out2 any](fnPt *func(In1, In2, In3) (Out1, Out2), out1 Out1, out2 Out2) (e error) {
	e = ItReturns(fnPt, out1, out2)
	return
}

func ItReturns3x3[In1 any,
	In2 any,
	In3 any,
	Out1 any,
	Out2 any,
	Out3 any](fnPt *func(In1, In2, In3) (Out1, Out2, Out3), out1 Out1, out2 Out2, out3 Out3) (e error) {
	e = ItReturns(fnPt, out1, out2, out3)
	return
}

func ItReturns3x4[In1 any,
	In2 any,
	In3 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any](fnPt *func(In1, In2, In3) (Out1, Out2, Out3, Out4), out1 Out1, out2 Out2, out3 Out3, out4 Out4) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4)
	return
}

func ItReturns3x5[In1 any,
	In2 any,
	In3 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any](fnPt *func(In1, In2, In3) (Out1, Out2, Out3, Out4, Out5), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5)
	return
}

func ItReturns3x6[In1 any,
	In2 any,
	In3 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any](fnPt *func(In1, In2, In3) (Out1, Out2, Out3, Out4, Out5, Out6), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6)
	return
}

func ItReturns3x7[In1 any,
	In2 any,
	In3 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any](fnPt *func(In1, In2, In3) (Out1, Out2, Out3, Out4, Out5, Out6, Out7), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7)
	return
}

func ItReturns3x8[In1 any,
	In2 any,
	In3 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any](fnPt *func(In1, In2, In3) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8)
	return
}

func ItReturns3x9[In1 any,
	In2 any,
	In3 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any,
	Out9 any](fnPt *func(In1, In2, In3) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8, out9 Out9) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8, out9)
	return
}

func ItReturns4x1[In1 any,
	In2 any,
	In3 any,
	In4 any,
	Out1 any](fnPt *func(In1, In2, In3, In4) Out1, out1 Out1) (e error) {
	e = ItReturns(fnPt, out1)
	return
}

func ItReturns4x2[In1 any,
	In2 any,
	In3 any,
	In4 any,
	Out1 any,
	Out2 any](fnPt *func(In1, In2, In3, In4) (Out1, Out2), out1 Out1, out2 Out2) (e error) {
	e = ItReturns(fnPt, out1, out2)
	return
}

func ItReturns4x3[In1 any,
	In2 any,
	In3 any,
	In4 any,
	Out1 any,
	Out2 any,
	Out3 any](fnPt *func(In1, In2, In3, In4) (Out1, Out2, Out3), out1 Out1, out2 Out2, out3 Out3) (e error) {
	e = ItReturns(fnPt, out1, out2, out3)
	return
}

func ItReturns4x4[In1 any,
	In2 any,
	In3 any,
	In4 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any](fnPt *func(In1, In2, In3, In4) (Out1, Out2, Out3, Out4), out1 Out1, out2 Out2, out3 Out3, out4 Out4) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4)
	return
}

func ItReturns4x5[In1 any,
	In2 any,
	In3 any,
	In4 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any](fnPt *func(In1, In2, In3, In4) (Out1, Out2, Out3, Out4, Out5), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5)
	return
}

func ItReturns4x6[In1 any,
	In2 any,
	In3 any,
	In4 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any](fnPt *func(In1, In2, In3, In4) (Out1, Out2, Out3, Out4, Out5, Out6), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6)
	return
}

func ItReturns4x7[In1 any,
	In2 any,
	In3 any,
	In4 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any](fnPt *func(In1, In2, In3, In4) (Out1, Out2, Out3, Out4, Out5, Out6, Out7), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7)
	return
}

func ItReturns4x8[In1 any,
	In2 any,
	In3 any,
	In4 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any](fnPt *func(In1, In2, In3, In4) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8)
	return
}

func ItReturns4x9[In1 any,
	In2 any,
	In3 any,
	In4 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any,
	Out9 any](fnPt *func(In1, In2, In3, In4) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8, out9 Out9) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8, out9)
	return
}

func ItReturns5x1[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	Out1 any](fnPt *func(In1, In2, In3, In4, In5) Out1, out1 Out1) (e error) {
	e = ItReturns(fnPt, out1)
	return
}

func ItReturns5x2[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	Out1 any,
	Out2 any](fnPt *func(In1, In2, In3, In4, In5) (Out1, Out2), out1 Out1, out2 Out2) (e error) {
	e = ItReturns(fnPt, out1, out2)
	return
}

func ItReturns5x3[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	Out1 any,
	Out2 any,
	Out3 any](fnPt *func(In1, In2, In3, In4, In5) (Out1, Out2, Out3), out1 Out1, out2 Out2, out3 Out3) (e error) {
	e = ItReturns(fnPt, out1, out2, out3)
	return
}

func ItReturns5x4[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any](fnPt *func(In1, In2, In3, In4, In5) (Out1, Out2, Out3, Out4), out1 Out1, out2 Out2, out3 Out3, out4 Out4) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4)
	return
}

func ItReturns5x5[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any](fnPt *func(In1, In2, In3, In4, In5) (Out1, Out2, Out3, Out4, Out5), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5)
	return
}

func ItReturns5x6[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any](fnPt *func(In1, In2, In3, In4, In5) (Out1, Out2, Out3, Out4, Out5, Out6), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6)
	return
}

func ItReturns5x7[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any](fnPt *func(In1, In2, In3, In4, In5) (Out1, Out2, Out3, Out4, Out5, Out6, Out7), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7)
	return
}

func ItReturns5x8[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any](fnPt *func(In1, In2, In3, In4, In5) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8)
	return
}

func ItReturns5x9[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any,
	Out9 any](fnPt *func(In1, In2, In3, In4, In5) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8, out9 Out9) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8, out9)
	return
}

func ItReturns6x1[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	Out1 any](fnPt *func(In1, In2, In3, In4, In5, In6) Out1, out1 Out1) (e error) {
	e = ItReturns(fnPt, out1)
	return
}

func ItReturns6x2[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	Out1 any,
	Out2 any](fnPt *func(In1, In2, In3, In4, In5, In6) (Out1, Out2), out1 Out1, out2 Out2) (e error) {
	e = ItReturns(fnPt, out1, out2)
	return
}

func ItReturns6x3[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	Out1 any,
	Out2 any,
	Out3 any](fnPt *func(In1, In2, In3, In4, In5, In6) (Out1, Out2, Out3), out1 Out1, out2 Out2, out3 Out3) (e error) {
	e = ItReturns(fnPt, out1, out2, out3)
	return
}

func ItReturns6x4[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any](fnPt *func(In1, In2, In3, In4, In5, In6) (Out1, Out2, Out3, Out4), out1 Out1, out2 Out2, out3 Out3, out4 Out4) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4)
	return
}

func ItReturns6x5[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any](fnPt *func(In1, In2, In3, In4, In5, In6) (Out1, Out2, Out3, Out4, Out5), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5)
	return
}

func ItReturns6x6[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any](fnPt *func(In1, In2, In3, In4, In5, In6) (Out1, Out2, Out3, Out4, Out5, Out6), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6)
	return
}

func ItReturns6x7[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any](fnPt *func(In1, In2, In3, In4, In5, In6) (Out1, Out2, Out3, Out4, Out5, Out6, Out7), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7)
	return
}

func ItReturns6x8[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any](fnPt *func(In1, In2, In3, In4, In5, In6) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8)
	return
}

func ItReturns6x9[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any,
	Out9 any](fnPt *func(In1, In2, In3, In4, In5, In6) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8, out9 Out9) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8, out9)
	return
}

func ItReturns7x1[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	Out1 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7) Out1, out1 Out1) (e error) {
	e = ItReturns(fnPt, out1)
	return
}

func ItReturns7x2[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	Out1 any,
	Out2 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7) (Out1, Out2), out1 Out1, out2 Out2) (e error) {
	e = ItReturns(fnPt, out1, out2)
	return
}

func ItReturns7x3[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	Out1 any,
	Out2 any,
	Out3 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7) (Out1, Out2, Out3), out1 Out1, out2 Out2, out3 Out3) (e error) {
	e = ItReturns(fnPt, out1, out2, out3)
	return
}

func ItReturns7x4[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7) (Out1, Out2, Out3, Out4), out1 Out1, out2 Out2, out3 Out3, out4 Out4) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4)
	return
}

func ItReturns7x5[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7) (Out1, Out2, Out3, Out4, Out5), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5)
	return
}

func ItReturns7x6[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7) (Out1, Out2, Out3, Out4, Out5, Out6), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6)
	return
}

func ItReturns7x7[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7) (Out1, Out2, Out3, Out4, Out5, Out6, Out7), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7)
	return
}

func ItReturns7x8[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8)
	return
}

func ItReturns7x9[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any,
	Out9 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8, out9 Out9) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8, out9)
	return
}

func ItReturns8x1[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	Out1 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8) Out1, out1 Out1) (e error) {
	e = ItReturns(fnPt, out1)
	return
}

func ItReturns8x2[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	Out1 any,
	Out2 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8) (Out1, Out2), out1 Out1, out2 Out2) (e error) {
	e = ItReturns(fnPt, out1, out2)
	return
}

func ItReturns8x3[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	Out1 any,
	Out2 any,
	Out3 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8) (Out1, Out2, Out3), out1 Out1, out2 Out2, out3 Out3) (e error) {
	e = ItReturns(fnPt, out1, out2, out3)
	return
}

func ItReturns8x4[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8) (Out1, Out2, Out3, Out4), out1 Out1, out2 Out2, out3 Out3, out4 Out4) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4)
	return
}

func ItReturns8x5[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8) (Out1, Out2, Out3, Out4, Out5), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5)
	return
}

func ItReturns8x6[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8) (Out1, Out2, Out3, Out4, Out5, Out6), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6)
	return
}

func ItReturns8x7[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8) (Out1, Out2, Out3, Out4, Out5, Out6, Out7), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7)
	return
}

func ItReturns8x8[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8)
	return
}

func ItReturns8x9[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any,
	Out9 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8, out9 Out9) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8, out9)
	return
}

func ItReturns9x1[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	In9 any,
	Out1 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8, In9) Out1, out1 Out1) (e error) {
	e = ItReturns(fnPt, out1)
	return
}

func ItReturns9x2[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	In9 any,
	Out1 any,
	Out2 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8, In9) (Out1, Out2), out1 Out1, out2 Out2) (e error) {
	e = ItReturns(fnPt, out1, out2)
	return
}

func ItReturns9x3[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	In9 any,
	Out1 any,
	Out2 any,
	Out3 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8, In9) (Out1, Out2, Out3), out1 Out1, out2 Out2, out3 Out3) (e error) {
	e = ItReturns(fnPt, out1, out2, out3)
	return
}

func ItReturns9x4[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	In9 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8, In9) (Out1, Out2, Out3, Out4), out1 Out1, out2 Out2, out3 Out3, out4 Out4) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4)
	return
}

func ItReturns9x5[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	In9 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8, In9) (Out1, Out2, Out3, Out4, Out5), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5)
	return
}

func ItReturns9x6[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	In9 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8, In9) (Out1, Out2, Out3, Out4, Out5, Out6), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6)
	return
}

func ItReturns9x7[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	In9 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8, In9) (Out1, Out2, Out3, Out4, Out5, Out6, Out7), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7)
	return
}

func ItReturns9x8[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	In9 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8, In9) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8)
	return
}

func ItReturns9x9[In1 any,
	In2 any,
	In3 any,
	In4 any,
	In5 any,
	In6 any,
	In7 any,
	In8 any,
	In9 any,
	Out1 any,
	Out2 any,
	Out3 any,
	Out4 any,
	Out5 any,
	Out6 any,
	Out7 any,
	Out8 any,
	Out9 any](fnPt *func(In1, In2, In3, In4, In5, In6, In7, In8, In9) (Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9), out1 Out1, out2 Out2, out3 Out3, out4 Out4, out5 Out5, out6 Out6, out7 Out7, out8 Out8, out9 Out9) (e error) {
	e = ItReturns(fnPt, out1, out2, out3, out4, out5, out6, out7, out8, out9)
	return
}
