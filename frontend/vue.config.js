module.exports = {
  devServer: {
    host: "0.0.0.0",
    port: "4396",
    //proxy: "http://192.168.0.6:1235"
    proxy: {
        "/rpc": {
            target: process.env.VUE_APP_BASE_URL,
            changeOrigin: true
        }
    }
  }
};
