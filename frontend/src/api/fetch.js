import axios from "axios";

//axios.defaults.headers.post["Content-Type"] = "application/json;charset=utf-8";
//axios.defaults.headers.post["Access-Control-Allow-Origin"] = "*";
//axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';

export function fetch(options) {
/*
  return new Promise((resolve, reject) => {
    const instance = axios.create({
      baseURL: process.env.VUE_APP_BASE_URL,
      //baseURL: "https://filscan.io/:8700/v0/filscan",
      timeout: 100000
    });
    instance.interceptors.response.use(
      response => {
        let data;
        if (response.data === undefined) {
          data = response.request.responseText;
        } else {
          data = response.data;
        }
        return data.data || data;
      },
      err => {
        return Promise.reject(err);
      }
    );
    instance(options)
      .then(res => {
        resolve(res);
      })
      .catch(error => {
        reject(error);
      });
  });
*/
    return options;
}

export function fetch2(options) {
  return new Promise((resolve, reject) => {
    const instance = axios.create({
      //! CORS Error: DO NOT USE baseURL
      //baseURL: process.env.VUE_APP_BASE_URL,
      timeout: 100000
    });
    instance.interceptors.response.use(
      response => {
        let data;
        if (response.data === undefined) {
          data = response.request.responseText;
        } else {
          data = response.data;
        }

        //console.log( data )
        return data.data || data;
      },
      err => {
        return Promise.reject(err);
      }
    );
    instance(options)
      .then(res => {
        resolve(res);
      })
      .catch(error => {
        reject(error);
      });
  });
}

/*
// JSON-RPC POST Request
export function fetch3() {
    axios.post (
        "/rpc",
        //! CORS Error: DO NOT USE "Full URL", USE URL WITHOUT DOMAIN
        //process.env.VUE_APP_BASE_URL,
        {
            jsonrpc: "2.0",
            id: "1",
            method: "LocalDB.JSONRPC_get_txns_all_mixed",
            params: [{ "Dummy": 0 }]
        }
    ).then(
        response => {
            console.log( response.result );
        }).catch (error => {
            console.log( "ERROR = " );
            console.log( error );
        });
}
*/
