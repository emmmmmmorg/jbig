jbig lib for go

[c lib](https://github.com/qyot27/jbigkit)

## Install

```
go get github.com/emmmmmmorg/jbig
```


## demo

```
package main

import (
        "encoding/base64"
        "fmt"
        "io/ioutil"

        "github.com/emmmmmmorg/jbig"
)

func main() {
        data := "AAABAAAAAPAAAABQAAAABwgAAxz/AtMEpAeH58oV99IAW1/g/wIY/OPZmcC/IfgLo2aeNY2r/wLKAY8dcDGcqe58dRFZllJkgP8CQ63bWc8QQvnn5OmaI8enxdXesP8CCKvsdfZUkeaF0pFvfKrv2rTVyNvY5wljvz+fNeoT/wK42Qs2GP5kXS09M8kfEK8nNVJWruK/x+eMFD/iUhEiKMj/An7hqN/j9mEsebmSOzB/Ty6lv9JdvFErbjRP++Prgi6f7bpa/wIBIDeSnbEtCaDQ/aNBzh5ebaQa4totMpqFRFgTjrVxzbxUcZvA/wLDW1QzXdtoAecI1PUZo8H5LJhux/SndhysIP8CF2gTB4wF4rLI/wLfgP8C"
        jbigsrc, _ := base64.StdEncoding.DecodeString(data)

        fmt.Printf("%X\n", jbigsrc)
        pbmbuf, err := jbig.JBIGConvToPBM([]byte(jbigsrc))
        if err != nil {
                fmt.Printf("转换非法[%s]", err)
                return
        }

        ioutil.WriteFile("r.pbm", pbmbuf, 0644)
        return

}

```
