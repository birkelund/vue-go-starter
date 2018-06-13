module.exports = {
    assetsDir : 'static',

    devServer: {
        proxy: {
            '/api': {
                target: 'http://localhost:4000'
            }
        }
    }
}
