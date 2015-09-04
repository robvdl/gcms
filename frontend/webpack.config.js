var ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
  entry: './app/main',
  output: {
    path: '../static',
    filename: 'js/bundle.js'
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
    new ExtractTextPlugin('css/app.css')
  ],
  cssnext: {
    browsers: 'last 2 versions'
  }
}
