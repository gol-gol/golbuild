
## golbuild

> a package to allow me prepare build info file to be embedded

### Cmd

> Can install with `go get -u github.com/gol-gol/golbuild/cmd`.

* Running `golbuild -f .golbuild.json` would create file for current build.

* Following code blocks would let you use it

> adding public declaration for it to be embedded as bytes

```
//go:embed .golbuild.json                                                   
BuildBytes []byte
```

### Public Functions

* to get the struct `GolBuild` from the `BuildBytes` above

```
var build golbuild.GolBuild
err := golbuild.Unmarshal(BuildBytes, &build)
```

* to directly display details from `BuildBytes`, can use `golbuild.Print(BuildBytes)`

---
