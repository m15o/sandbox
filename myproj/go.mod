module sandbox/myproj

go 1.15

require golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
require sandbox/mylib v0.0.0-00010101000000-000000000000

replace sandbox/mylib => ../mylib

