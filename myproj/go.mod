module sandbox/myproj

go 1.15

require golang.org/x/xerrors 5ec99f83aff1

require sandbox/mylib v0.0.0-00010101000000-000000000000

replace sandbox/mylib => ../mylib
