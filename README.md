### nextfetch

> Simple cross platform fetch program, written in Go

#### Prerequisites
- True color (24-bit) or 256-color (8-bit) compatible terminals
- `go >= 17` (build only)

#### Building
```sh
git clone https://github.com/AlphaNecron/nextfetch.git
cd nextfetch
# build the program
go build
./nextfetch
# or run it directly
go run main.go
```

#### Config
> Copy `nextfetch.default` to `$HOME/.nextfetch` on Linux or `%USERPROFILE%/.nextfetch` on Windows.
```json
{
    "noAscii": false, // Whether to print ASCII art
    "primaryColor": "blue", // Default color, can be either hex without hash, "r,g,b", color name or number (256-color)
    "colorBlock": "▁▁", // Character to be used as color block
    "clearScreen": false // Whether to clear screen on startup
}
```

#### Screenshot
![Nextfetch](https://user-images.githubusercontent.com/57827456/153422522-7c9b7452-ee04-4e75-90c9-dce7ef03432e.png)
