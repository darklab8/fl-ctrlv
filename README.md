# Description

Making Ctrl-V work for the game Freelancer(2003).
Providing versions for
- Windows
- Linux (X11) (should work Ubuntu at least as it was compiled on it.)
    - If it will not work for other linuxes... you need to compile it from this specific Linux u have

# Download

- [Download from releases](<https://github.com/darklab8/fl-darkctrlv/releases>)

# How to use?

- Turn it on. (By mouse clicking executor, or from console)
- Copy some text into your clipboard (Ctrl+C)
- press Ctrl+V in the game.
- Release Ctrl-V as soon as possible. You have 1 second to release it. (If u will not release it, text will be written as you if u had pressed Ctrl)

Window Title Targeting:
- for convenience excluded its working in not Freelancer (or not Wine) windows.
- it seeks for title to contain "Freelancer" or "Wine".
- env var TITLE_CONTAINS can override behavior for targeting any other window

Limitations:
- You try to print `*`, it could print `8` instead potentially (works from Linux correctly)
- It looks to be able to handle Capitalizing letters on its own, but watch out for any not ordinary english symbols.
- Not printing at least brackets

# For other Freelancer dark tools

- [link](<https://darklab8.github.io/blog/community_freelancer.html>)

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

Install the stuff and run `task build:linux` (if u have installed taskfile.dev), or just check commands inside Taskfile.yml


# Building from other linuxes

See here https://github.com/go-vgo/robotgo?tab=readme-ov-file#requirements what needs to be installed
then run same command go build command, the same as for ubuntu

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
