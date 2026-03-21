# internal-package-demo

Given this directory structure:

```
├── cow
│   └── bull.go
├── go.mod
├── internal
│   ├── bar
│   │   └── bar.go
│   ├── foo
│   │   └── foo.go
│   └── moo.go
├── main.go
```

Please read the files, including go.mod to understand the use of internal packages (foo and bar) versus exportable (public) package such as cow.

