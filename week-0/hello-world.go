package main

import (
	"fmt"
	testpackage "learninggolang/week-0/test-package"
)

func main() {
	fmt.Println("Hello world")
	values()
	testpackage.TestPkg()
	varibles()
	loop()

}

// To run the program, put the code in hello-world.go and use go run.

// $ go run hello-world.go
// hello world

// Sometimes we’ll want to build our programs into binaries. We can do this using go build.

// $ go build hello-world.go
// $ ls

/*
GO LEARNING NOTES / MISTAKES I MADE (IMPORTANT)

1) Go does NOT import files — it imports PACKAGES
   - All .go files in the same folder with the same `package` name
     are compiled together automatically.
   - You never import another .go file directly.

2) Only ONE `main()` function is allowed
   - `main()` exists only in `package main`
   - Other logic should be in separate functions or packages.

3) Go requires a MODULE (go.mod) for cross-folder imports
   - Without `go mod init`, Go searches only GOROOT (stdlib)
   - Always run: `go mod init <module-name>` at project root

4) Import paths are MODULE-BASED, not relative paths
   - ❌ import "week-0/test-package"
   - ✅ import "learninggolang/week-0/testpackage"

5) Folder name ≠ package name
   - Import path uses folder name
   - Code uses package name
   - They can be different (but usually shouldn’t be)

6) Hyphens (-) are NOT allowed in Go identifiers
   - Folder name like `test-package` forces an import alias
   - VS Code auto-adds:
       testpackage "learninggolang/week-0/test-package"
   - This is correct behavior, not a bug

7) Exported functions MUST start with a capital letter
   - func TestPkg() → accessible from other packages
   - func testPkg() → package-private

8) Best practice: avoid import aliases unless necessary
   - Rename folders to valid identifiers instead
   - Example:
       test-package ❌
       testpackage  ✅

9) Always run Go programs from the module root
   - go run .
   - go run ./week-0
   - NOT from random subfolders

MENTAL MODEL:
- go.mod defines the project root
- Imports start from module name
- Packages are directories
- Files are implementation details
*/
