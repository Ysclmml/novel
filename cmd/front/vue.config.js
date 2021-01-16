const path = require('path')

function resolve(dir) {
    return path.join(__dirname, '.', dir)
}

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
    },
    // svg配置
    chainWebpack: config => {
        config.module.rules.delete("svg") //重点:删除默认配置中处理svg,
        // const svgRule = config.module.rule('svg')
        // svgRule.uses.clear()
        config.module
            .rule('svg-sprite-loader')
            .test(/\.svg$/)
            .include
            .add(resolve('src/assets')) //处理svg目录
            .end()
            .use('svg-sprite-loader')
            .loader('svg-sprite-loader')
            .options({
                symbolId: 'icon-[name]'
            })
    },
}
