# Go Sandbox

## Resources

- [Go by Example](https://gobyexample.com/): hands-on introduction to Go using annotated example programs.
- [Gophercises](https://gophercises.com/): course where you build 20 mini-applications

## Run Programs

`go run <filename>` or
`go build <filename> && ./<filename_without_go_suffix>` (creates a binary that you can execute directly)

## Type Hints

Install the Go VSCode extension

## Auto formatting

Add this to settings.json

```json
"[go]": {
  "editor.formatOnSave": true,
  "editor.defaultFormatter": "golang.go"
},
"go.formatTool": "goimports"
```
