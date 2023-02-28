# HTTP JSON-RPC Test
# hjkim, 2022.03.23

# pip3 install urllib2

import urllib.request
import json


def rpc_call(url, method, args):
    data = json.dumps({
        'id': 1,
        'method': method,
        'params': [args]
        }).encode('utf8')
    req = urllib.request.Request( url, data, {'Content-Type': 'application/json'})

    #req = urllib.request.Request(url)
    #req.add_header('Authorization', 'Bearer ' + token)
    #req.add_header("Content-Type", "application/json")

    response = urllib.request.urlopen(req)
    #response = urllib.request.urlopen(req, data)

    _res = response.read()
    res = _res.decode("utf8")
    return json.loads(res)

url = 'http://192.168.0.6:1235/rpc'
args = {'Dummy': 1}
result = rpc_call(url, "LocalDB.JSONRPC_get_txns_all_mixed", args)
print( "result =" )
print( result )
