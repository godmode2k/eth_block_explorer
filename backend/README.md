# Ethereum Explorer


Summary
----------
> Backend JSON-RPC Server for Ethereum Explorer </br>


Environment
----------
> build all and tested on GNU/Linux

    GNU/Linux: Ubuntu 20.04_x64 LTS
    Ethereum: geth vx.x.x
    Python: v3.8.10
    Go: go1.15.5 linux/amd64
    MySQL: 8.0.30
    Network: Ethereum Private Network


Run
----------
```sh
// Database
$ mysql -u xxx -p < mysql.sql
mysql> CREATE USER 'ethereum_block_explorer'@'localhost' IDENTIFIED BY 'test';
mysql> REVOKE ALL ON ethereum_block_explorer.* FROM 'ethereum_block_explorer'@'localhost';
mysql> GRANT ALL ON ethereum_block_explorer.* TO 'ethereum_block_explorer'@'localhost';


// Dependencies
$ go get -u github.com/go-sql-driver/mysql
$ go get github.com/mattn/go-sqlite3


// Fetcher (standalone)
//$ go run ./eth_txns_fetcher_main.go


// JSON-RPC Server + Fetcher
$ go run ./rpc_server_main.go


// JSON-RPC Client (test, Go version)
$ go run ./rpc_client_main.go


// JSON-RPC Client (test, Python version)
$ pip3 install urllib2
$ python3 ./rpc_client_main.py

```

