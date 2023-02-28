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
    Go: go1.15.5 linux/amd64
    Network: Ethereum Private Network
    Node.js: node-v16.13.2
    MySQL: v8.0.32


Backend
----------
> HTTP RPC Server: </br>
> Fetcher Service: Supports ETH, ERC-20, ERC-1155 (NFT)


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
Golang
---------------------------------
$ wget https://go.dev/dl/go1.15.5.linux-amd64.tar.gz



---------------------------------
Node.js
---------------------------------
$ wget https://nodejs.org/dist/v16.13.2/node-v16.13.2-linux-x64.tar.xz
$ tar xJvf node-v16.13.2-linux-x64.tar.xz
$ export PATH=`pwd`/node-v16.13.2-linux-x64/bin:$PATH
$ echo export PATH=`pwd`/node-v16.13.2-linux-x64/bin:$PATH >> $HOME/.profile
$ echo export PATH=`pwd`/node-v16.13.2-linux-x64/bin:$PATH >> $HOME/.bashrc



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

```


Screenshots
----------

> Ethereum Block Explorer </br>
<img src="https://github.com/godmode2k/eth_block_explorer/raw/main/screenshot.png" width="50%" height="50%">

