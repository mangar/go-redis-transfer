# Go Redis Transfer

_Motivation_
Transfer data from one Redis to another



## Development Setup

1. Start the first Redis server
```
docker run -d -p 6379:6379 -i -t redis
```
2. Populate the first Redis server (only for development)
3. Start the second Redis server
```
docker run -d -p 6380:6379 -i -t redis
```
4. Transfer keys from the first to the second Redis server






## Configuration


In the root directory (not the src dir):
```
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
export GOPATH=`pwd`
```
