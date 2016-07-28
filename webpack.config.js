var path = require('path');
var HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  context: __dirname,
  entry: "./javascript/src/index.js",
  output: {
    path: "out/public/js",
    filename: "app.js",
    hash: true
  },
  module: {
    loaders: [
      {
        test: /\.jsx?$/,
        exclude: /(node_modules|bower_components)/,
        loader: 'babel',
        query: {
            presets: ['react', 'es2015']
        }
      },

      {
        test: /\.css$/,
        loader: 'style-loader!css-loader'
      },

      {
        test: /\.(png|woff|woff2|eot|ttf|svg)$/,
        loader: 'url-loader?limit=100000'
      }
    ]
  },
  plugins: [new HtmlWebpackPlugin({
        template: 'index.html'
    })
  ],
  devServer: {
    hot: true,
    port: 3000
  }
};