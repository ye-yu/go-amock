package main

import (
	"fmt"
	"github.com/mariomac/gostream/stream"
	"io"
	"os"
	"strings"
)

const sample = `
func ItReturns0x1[
	Out1 any,
](
	fnPt *func() Out1,
	out1 Out1,
) (e error) {
	e = ItReturns(fnPt, out1)
	return
}
`

const template = `
func ItReturns%s[%s](fnPt *%s,%s) (e error) {
	e = ItReturns(fnPt, %s)
	return
}
`

func main() {
	funcs := make([]string, 0)
	for ins := 0; ins <= 9; ins++ {
		insArr := make(
			[]string,
			ins,
		)
		for i := range insArr {
			insArr[i] = fmt.Sprintf("In%d", i+1)
		}

		for outs := 1; outs <= 9; outs++ {
			outsArr := make(
				[]string,
				outs,
			)
			for i := range outsArr {
				outsArr[i] = fmt.Sprintf("Out%d", i+1)
			}

			gParams := append(insArr, outsArr...)

			matrix := fmt.Sprintf("%dx%d", ins, outs)
			genericsDecl := strings.Join(
				stream.
					OfSlice(gParams).
					Map(func(s string) string {
						return fmt.Sprintf("%s any,", s)
					}).
					ToSlice(),
				"\n",
			)
			fnInParams := strings.Join(insArr, ",")

			fnOutParams := strings.Join(outsArr, ",") + ","

			fnParam := fmt.Sprintf(
				"func(%s) (%s)",
				fnInParams,
				fnOutParams,
			)
			fmt.Fprintf(io.Discard, "%s%s%s", matrix, genericsDecl, fnParam)
			outsParams := strings.Join(
				stream.
					OfSlice(outsArr).
					Map(func(s string) string {
						return fmt.Sprintf("%s %s", strings.ToLower(s), s)
					}).
					ToSlice(),
				",")
			itReturnsParams := strings.Join(
				stream.
					OfSlice(outsArr).
					Map(strings.ToLower).
					ToSlice(),
				",")

			res := fmt.Sprintf(template,
				matrix,
				genericsDecl,
				fnParam,
				outsParams,
				itReturnsParams,
			)

			funcs = append(funcs, res)

		}
	}

	content := fmt.Sprintf("package amock\n\n%s", strings.Join(funcs, "\n\n"))
	os.WriteFile("amock/it_returns_generics.go", []byte(content), 0644)
}
