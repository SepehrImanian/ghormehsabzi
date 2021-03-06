## Ghormehsabzi
Generate **Concurrent Request** for load testing with **Ghormehsabzi** :-0

### Installation

download latest release file and move it in in **/usr/local/bin** for linux users

```bash
wget https://github.com/SepehrImanian/ghormehsabzi/releases/download/v0.0.2/gs
chmod +x ./gs
mv gs /usr/local/bin/gs
```

### CLI Interface

use help command
```bash
gs -h or --help
```
```bash
NAME:
   gs - Load test with ghormehsabzi

USAGE:
   gs [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --url value, -u value         http url for load test
   --requests value, -r value    Number of requests (default: 1)
   --concurrent value, -c value  Number of concurrent requests (default: 1)
   --help, -h                    show help (default: false)
   --version, -v                 print the version (default: false)
```

### Example
Generate load test with 100 concurrent process and 100 requests each of them
```bash
gs -u "https://example.com" -c 100 -r 100
```
**Output**:
```bash
------------ Starting DDOS ---------- url :  https://example.com  Concurrents :  100   Requests :  100
count :  1   Open Concurrent Process   url/host :  https://example.com
count :  2   Open Concurrent Process   url/host :  https://example.com
count :  3   Open Concurrent Process   url/host :  https://example.com
count :  4   Open Concurrent Process   url/host :  https://example.com
count :  5   Open Concurrent Process   url/host :  https://example.com
count :  6   Open Concurrent Process   url/host :  https://example.com
count :  7   Open Concurrent Process   url/host :  https://example.com
count :  8   Open Concurrent Process   url/host :  https://example.com
count :  9   Open Concurrent Process   url/host :  https://example.com
...
```
