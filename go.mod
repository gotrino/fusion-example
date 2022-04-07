module github.com/gotrino/fusion-example

go 1.18

require (
	github.com/gotrino/fusion v0.0.0-20220406164020-8c11dfc69c0d
	github.com/gotrino/fusion-rt-wasmjs v0.0.0-20220406164243-245388362c48
)

require honnef.co/go/js/dom/v2 v2.0.0-20210725211120-f030747120f2 // indirect

replace github.com/gotrino/fusion => ../fusion.git

replace github.com/gotrino/fusion-rt-wasmjs => ../fusion-rt-wasmjs.git
