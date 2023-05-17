# go-embedfs-http

![License](https://img.shields.io/badge/license-MIT-blue)

go-embedfs-http is a super small http.FileSystem that serves files from an embedded filesystem
(designed for go 1.16' embed feature), while also disabling directory listings and strips the path prefix.

I outsourced this into a separate module, because I use it in multiple projects.

## Installation

Add it to your go.mod like every other module:
```bash
go get github.com/d0x7/go-embedfs-http
```

## License

[MIT](LICENSE) © 2023 Dorian Heinrichs
