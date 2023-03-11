
## golbuild

> a package to allow me prepare build info file to be embedded

### Cmd

> Prepare a binary for you as `go build -o golbuild cmd/golbuild.go`.

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

* also if want to manage yourselves, can utilize `(build *GolBuild) FetchDetails()` and `(build *GolBuild) Print()`

---
