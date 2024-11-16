# Description

Making Ctrl-V work for the game Freelancer(2003).
Providing versions for Linux and Windows.

# Download

- Insert here link to download

# How to use?

- Turn it on. (By mouse clickig executor, or from console)
- Copy some text into your clipboard (Ctrl+C)
- press Ctrl+V in the game.
- Release Ctrl-V as soon as possible. You have 1 second to release it. (If u will not release it, text will be written as you had pressed Ctrl)

Limitations:
- You try to print `*`, it will print `8` instead.
- It looks to be able to handle Capitalizing letters on its own, but watch out for any not ordinary english symbols.
- Brackets i noticed it is not printing too

# Development.

- Build only from OS to which u need it. Cross compiling between OSes leads to weird bugs.

# Building from Ubuntu

```
# Copy paste from https://github.com/go-vgo/robotgo?tab=readme-ov-file#requirements
# gcc
sudo apt install gcc libc6-dev

# x11
sudo apt install libx11-dev xorg-dev libxtst-dev

# Clipboard
sudo apt install xsel xclip

#
# Bitmap
sudo apt install libpng++-dev

# GoHook
sudo apt install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev libxkbcommon-dev
```

Install the stuff and run task build:linux (if u have installed taskfile.dev), or just check commands inside Taskfile.yml

# Building from windows

- Download already built mingw from https://winlibs.com/
- Unzip to for C:\mingw64
- and add C:\mingw64\bin folder into `Path` (search for Environment variables in taskbar)
- task build:windows (or just check commands inside)
    - builds greats if running the command from `Git Bash` console

# Other useful links

- https://github.com/robotn/gohook?tab=readme-ov-file
- https://github.com/go-vgo/robotgo?tab=readme-ov-file#requirements
- https://pkg.go.dev/golang.design/x/clipboard#section-readme

P.S. you probably don't need separate cllipboard program, as it is present in robotgo.
