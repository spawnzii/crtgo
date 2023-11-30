# crtgo

This tool is used to retrieve the list of subdomains available on [crt.sh](https://crt.sh/).

## Install
- With go install :
`go install github.com/spawnzii/crtgo@latest`
- With git clone :
```
git clone https://github.com/spawnzii/crtgo.git
cd crtgo/
go build
```
## Usage
```
cat targets.txt | crtgo

crtgo -domain foo.com
```
