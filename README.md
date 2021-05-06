# MicroServerGoAuth

Bootstraps a Go-based service that demonstrates server-side auth support for your growing Mifedom.

**NOTE:**  Built to work seamlessly with [bin-go](https://github.com/wejafoo/bin-go) lightweight mife build/test/deploy utility

----

## Developer Install

Clone the git repository and link project root to your path.

$  `git clone git@github.com:micro-cosm/micro-server-go-auth.git`

$  `go mod init weja.us/micro/micro-server-go-auth`     # if needed 

$  `go mod tidy`


### Deploy Local Docker

$   `bingo --local`

### Deploy Remote Docker/Cloud

$   `bingo --remote --alias stage`
