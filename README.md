Influxseed
==

Build
--

```bash
$ go build -ldflags "-X main.version=$(git rev-parse --abbrev-ref HEAD) -X main.built=$(date +'%Y%m%d-%H%M%S')"
```
