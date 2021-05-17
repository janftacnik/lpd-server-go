# lpd-server-go

LPD server written in Go. Uses port 515, accepts print job (just like printer) and saves print job into file
with name prjob#number.prn, where #number is sequentially growing integer starting from 1. On top of that
another file is created called prjob#number.cfg that contains values used in so called LPR header in communication
between lpr client and lpd server. It is sometimes instructive to see what given client is sending.

LPD server has been tested on all possible combinations of current desktop operating systems (Windows, Linux, MacOS) and
processor architectures (intel, arm64).
