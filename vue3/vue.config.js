module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8096/api/v1',
        changeOrigin: true,
        pathRewrite: { '^/api': '' },
      },
    },
    host: 'localhost',
    hotOnly: true,
  },
  css: {
    requireModuleExtension: true,
  },
}
