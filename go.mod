module github.com/switfs/switfs-block

go 1.20

require (
	github.com/filecoin-project/lotus v1.23.2
	github.com/ipfs/go-log/v2 v2.5.1
	github.com/urfave/cli v1.22.10
	github.com/urfave/cli/v2 v2.16.3
)
replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/goleak v1.1.12 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
)
