module sandbox/myproj

go 1.15

require golang.org/x/xerrors v0.0.0-20240903120638-7835f813f4da

require sandbox/mylib v0.0.0-00010101000000-000000000000

replace sandbox/mylib => ../mylib
