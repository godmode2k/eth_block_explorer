# Ethereum Block Explorer


Summary
----------
> Ethereum Block Explorer </br>
>
> WORK IN-PROGRESS (NOTE: NOT READY for Production) </br>
> This is a test version. so, USE THIS AT YOUR OWN RISK.


Environment
----------
> build all and tested on GNU/Linux

    GNU/Linux: Ubuntu 20.04_x64 LTS
    Ethereum: geth vx.x.x
    Python: v3.8.10 (pip 20.0.2)
    Go: go1.22.5 linux/amd64
    Network: Ethereum Private Network
    Node.js: node-v20.15.0
    MySQL: v8.0.32


Backend
----------
> HTTP RPC Server: </br>
> Fetcher Service: Supports ETH, ERC-20, ERC-1155 (NFT) </br>
> NFT Metadata APIs


Frontend
----------
> Source-based from Filscan-frontend (Vue.js): https://github.com/Filscan-Team/filscan-frontend


Run
----------
```sh
---------------------------------
Python
---------------------------------
$ sudo apt-get install wget python3 python3-pip

$ sudo apt-get install mysql-server
or
$ sudo apt-get install mariadb-server



---------------------------------
Ethereum Private Network
---------------------------------
SEE:
 - https://github.com/godmode2k/blockchain/tree/master/build_guide
 - (Docker) https://github.com/godmode2k/blockchain/tree/master/build_guide/ethereum



---------------------------------
Golang
---------------------------------
$ wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
$ tar xzvf go1.22.5.linux-amd64.tar.gz -C /usr/local/
$ echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.profile
$ echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.bashrc



---------------------------------
Node.js
---------------------------------
$ wget https://nodejs.org/dist/v20.15.0/node-v20.15.0-linux-x64.tar.xz
$ tar xJvf node-v20.15.0-linux-x64.tar.xz
$ echo 'export PATH=$PATH:$HOME/node-v20.15.0-linux-x64/bin' >> $HOME/.profile
$ echo 'export PATH=$PATH:$HOME/node-v20.15.0-linux-x64/bin' >> $HOME/.bashrc



---------------------------------
Backend
---------------------------------
1. Ethereum
 - run 'geth' with port "8544"

2. RPC Server: backend/rpc_server_main.go
(Edit)
 - // Connect to an Ethereum RPC Server
 - var SERVER_ADDRESS = "127.0.0.1"
 - var SERVER_PORT = "8544"
 -
 - // Connect to Database
 - // SEE: mysql.sql
 - var DB_SERVER_ADDRESS = "127.0.0.1:3306"
 - var DB_NAME = "ethereum_block_explorer"
 - var DB_LOGIN_USERNAME = "root"
 - var DB_LOGIN_PASSWORD = "mysql"
 -
 - // HTTP RPC Server
 - var HTTP_RPC_SERVER_HOST_PORT = ":1234" // Internal
 - var HTTP_JSONRPC_SERVER_HOST_PORT = ":1235" // External

3. Fetcher Service: server/eth_txns_fetcher_main_func.go
(Edit)
 - // Connect to RPC Server
 - var SERVER_ADDRESS = "127.0.0.1"
 - var SERVER_PORT = "8544"

 - // Connect to Database
 - var DB_SERVER_ADDRESS = "127.0.0.1:3306"
 - var DB_NAME = "ethereum_block_explorer"
 - var DB_LOGIN_USERNAME = "root"
 - var DB_LOGIN_PASSWORD = "mysql"

// Dependencies
$ go get -u github.com/go-sql-driver/mysql
$ go get github.com/mattn/go-sqlite3

4. Database
$ mysql -u root -p < mysql.sql

5. NFT Metadata APIs: nft_meta_apis/nft_meta_api_golang/nft_meta_apis_main.go
(Edit)
 - var CERT = "/etc/ssl/example.com+4.pem"
 - var CERT_KEY = "/etc/ssl/example.com+4-key.pem"
 - //var HOST = ":443" // https
 - var HOST = ":8888" // http



---------------------------------
Frontend
---------------------------------
.env.development
.env.pre
.env.production
.env.test

(Edit)
 - // HTTP RPC Server (backend)
 - VUE_APP_BASE_URL=http://x.x.x.x:1235



vue.config.js

(Edit)
 - // Service port
 - port: "4396"


// dependency
$ npm install



---------------------------------
Run
---------------------------------
// Backend
$ cd backend
$ go run rpc_server_main.go


// Frontend
$ cd frontend
$ npm run serve


(Web Browser)
http://127.0.0.1:4396


// Backend: NFT Meta APIs {
$ cd nft_meta_apis/nft_meta_api_golang
$ sh run_ntf_meta_apis.sh

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
// Backend: NFT Meta APIs }



---------------------------------
NodeJS: node-v20.15.0
Error:
---------------------------------
// Frontend

$ cd frontend


// ERROR #1
$ npm install

(ERROR)
npm error gyp ERR! UNCAUGHT EXCEPTION
npm error gyp ERR! stack TypeError: Cannot assign to read only property 'cflags' of object '#<Object>'
...
npm error gyp ERR! node -v v20.15.0
npm error gyp ERR! node-gyp -v v7.1.2
npm error gyp ERR! Node-gyp failed to build your package.

(FIX)
// node-sass
// https://github.com/sass/node-sass
//
// NodeJS 20: 9.0+ (supported)
// NodeJS 19: 8.0+ (supported)
// ...

(EDIT)
$ vim package.json
"devDependencies": {
...
"node-sass": "^9.0.0",
...
}
$ npm install



// ERROR #2
$ npm run serve

(ERROR)
Error: error:0308010C:digital envelope routines::unsupported

(FIX)
$ export NODE_OPTIONS=--openssl-legacy-provider
$ npm run serve
or
$ NODE_OPTIONS=--openssl-legacy-provider npm run serve



// ERROR #3
$ NODE_OPTIONS=--openssl-legacy-provider npm run serve

(ERROR)
Node Sass version 9.0.0 is incompatible with ^4.0.0 || ^5.0.0 || ^6.0.0.

(FIX)
$ rm -f package-lock.json
$ rm -fr node_modules
$ npm install
$ npm run serve

```


Screenshots
----------

> Ethereum Block Explorer </br>
<img src="https://github.com/godmode2k/eth_block_explorer/raw/main/screenshots/screenshot.png" width="50%" height="50%">
<br>
<img src="https://github.com/godmode2k/eth_block_explorer/raw/main/screenshots/screenshot_1.jpg" width="50%" height="50%">
<br>
<img src="https://github.com/godmode2k/eth_block_explorer/raw/main/screenshots/screenshot_2.jpg" width="50%" height="50%">
<br>
<img src="https://github.com/godmode2k/eth_block_explorer/raw/main/screenshots/screenshot_3.jpg" width="50%" height="50%">

