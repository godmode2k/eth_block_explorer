/* --------------------------------------------------------------
Project:    NFT Meta APIs Server
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since September 3, 2024
Filename:   nft_meta_apis_main.go

License:

*
* Copyright (C) 2024 Ho-Jung Kim (godmode2k@hotmail.com)
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
-----------------------------------------------------------------
Note:
-----------------------------------------------------------------
Reference:
 - https://github.com/labstack/echo
 - https://echo.labstack.com/docs
 - https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
 - https://github.com/mattn/go-sqlite3/issues/204
 - https://pkg.go.dev/database/sql
 - https://pkg.go.dev/github.com/mattn/go-sqlite3

Dependencies:
// go get github.com/labstack/echo/{version}
$ go get github.com/labstack/echo/v4
$ go get github.com/labstack/echo/v4/middleware
$ go get github.com/mattn/go-sqlite3


Run:
// http
$ go run nft_meta_apis_main.go
// https: SSL/TLS
$ sudo go run nft_meta_apis_main.go

// Set HTTP, HTTPS
nft_meta_apis_main.go: main() {
    var CERT = "cert.pem"
    var CERT_KEY = "cert-key.pem"
    //var HOST = ":443"
    var HOST = ":8888"

    // http
    e.Logger.Fatal( e.Start(HOST) )

    // https: SSL/TLS
    e.Logger.Fatal( e.StartTLS(HOST, CERT, CERT_KEY) )
}


NFT Token creation:
$ yes "" | bash new_token_erc1155.sh NFT1 A,B 0,1 1,2 https://127.0.0.1/apis/tokens/nft1 privatenet 11112

NFT Metadata APIs: {
(resource path)
nft_meta_apis/nft_meta_api_golang/assets/resources/

Directory:    <NFT-Name>/: nft1/
Metadata :    <NFT-Name>_<id>_meta.json: nft1_0_meta.json
Image    :    <NFT-Name>_<id>.jpg: nft1_0.jpg

nft1/nft1_0_meta.json
nft1/nft1_0.jpg


(metadata)
$ curl http://172.17.0.2:8888/apis/tokens/nft1/0
$ curl http://172.17.0.2:8888/apis/tokens/nft1/1
{
  "description": "Test NFT1 #1 image: warty-final-ubuntu", 
  "external_url": "http://172.17.0.2:8888/apis/tokens/nft1/1", 
  "image": "http://172.17.0.2:8888/resources/tokens/nft1/1", 
  "name": "Test NFT1 #1",
  "attributes": []
}


(image)
$ wget http://172.17.0.2:8888/resources/tokens/nft1/0
$ wget http://172.17.0.2:8888/resources/tokens/nft1/1
}
-------------------------------------------------------------- */
package main



//! Header
// ---------------------------------------------------------------

import (
    "fmt"

    //"log"
    //"database/sql"

    //_ "github.com/mattn/go-sqlite3"

    //"net"
    "net/http"
    //"net/rpc"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "nft_meta_apis/include"
)



//! Definition
// --------------------------------------------------------------------
var CERT = "/etc/ssl/example.com+4.pem"
var CERT_KEY = "/etc/ssl/example.com+4-key.pem"
//var HOST = ":443"
var HOST = ":8888"



//! Implementation
// --------------------------------------------------------------------

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    //e.Use(middleware.Recover())

    e.GET("/request", func(c echo.Context) error {
        req := c.Request()
        format := `
            <code>
            Protocol: %s<br>
            Host: %s<br>
            Remote Address: %s<br>
            Method: %s<br>
            Path: %s<br>
            </code>
        `
        return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
    })

    // Assets

    // http://127.0.0.1:8888/resources/tokens/nft1/1
    // http://172.17.0.2:8888/resources/nft1/ubuntu-numpy-numbat-on-dell-xps-13.jpg
    //e.Static( "/", "./assets" )
    e.Static( "/", nft_meta_route.DEF_ASSETS_DIR )


    // /v1/apis/

    // /v1/apis/tokens
    e.GET( "/apis/tokens/:name/:id", nft_meta_route.Endpoint__tokens)
    e.POST( "/apis/tokens/:name/:id", nft_meta_route.Endpoint__tokens)

    // /v1/resources
    e.GET( "/resources/tokens/:name/:id", nft_meta_route.Endpoint__resources)
    e.POST( "/resources/tokens/:name/:id", nft_meta_route.Endpoint__resources)


    // http
    e.Logger.Fatal( e.Start(HOST) )

    // https: SSL/TLS
    //e.Logger.Fatal( e.StartTLS(HOST, CERT, CERT_KEY) )
}



