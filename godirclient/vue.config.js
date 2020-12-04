// vue.config.js
module.exports = {
    configureWebpack: {
        devServer: {
            proxy: {
                "/dirserv": {
                    target: process.env.VUE_APP_URL,
                    changeOrigin: true, // 是否改变域名
                    ws: true,
                    pathRewrite: {
                        "/dirserv": ""
                    }
                }
                
            }
        }
    }
}