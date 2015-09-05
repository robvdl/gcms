var webpack = require('webpack'),
    ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
  entry: {
    login: './app/login',
    admin: './app/admin'
  },
  output: {
    path: '../static',
    filename: 'js/[name].js'
  },
  module: {
    loaders: [
      {
        test: /\.css$/,
        loader: ExtractTextPlugin.extract('style-loader', 'css-loader!cssnext-loader')
      },
      {
        test: /\.js$/,
        loader: 'babel-loader'
      }
    ]
  },
  plugins: [
    new webpack.optimize.CommonsChunkPlugin('common', 'js/common.js'),
    new ExtractTextPlugin('css/[name].css')
  ],
  cssnext: {
    browsers: 'last 2 versions'
  }
}
