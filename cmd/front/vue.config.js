module.exports = {
    outputDir: process.env.outputDir,
    assetsDir: 'static',
    publicPath: '/',
    devServer: {
        open: false,
        host: '0.0.0.0',
        port: 8999,
        https: false,
        hotOnly: false,
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:8102',
                changeOrigin: true,
                secure: false,
                pathRewrite: {
                    '^/api': '',
                }
            },
        }
    }
}
