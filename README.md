# GoUp
[Go](https://github.com/golang/go) installer tool like [rustup](https://github.com/rust-lang/rustup)
## Installation
### Get binary
For Windows
```bash
go get -u github.com/sadesakaswl/goup & move %USERPROFILE%\go\bin\goup.exe %USERPROFILE%
```
For Unix
```bash
go get -u github.com/sadesakaswl/goup & mv ~/go/bin/goup ~
```
### Before Installation
Uninstall current Go installation
### Install with GoUp
For Windows
```bash
%USERPROFILE%\goup install
```
For Unix
```bash
~/goup install
```
## Install Go
This command downloads Go and updates it to the latest version, if available.
```bash
goup install
```
## Upgrade Go
This command upgrades Go to latest version.
```bash
goup upgrade
```
