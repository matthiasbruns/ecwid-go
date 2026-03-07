module github.com/matthiasbruns/ecwid-go/e2e

go 1.26.0

require (
	github.com/matthiasbruns/ecwid-go/config v0.0.0
	github.com/matthiasbruns/ecwid-go/ecwid v0.0.0
)

replace (
	github.com/matthiasbruns/ecwid-go/config => ../config
	github.com/matthiasbruns/ecwid-go/ecwid => ../ecwid
)
