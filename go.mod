module example.com/hello

go 1.26.3

require (
	com.practice/tools v0.0.0-00010101000000-000000000000
	com.practice/transport v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.19.0
)

require (
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.44.0 // indirect
)

replace com.practice/tools => ./tools

replace com.practice/transport => ./transport

replace check-utils => ./check-utils

replace heartbeat => ./heartbeat
