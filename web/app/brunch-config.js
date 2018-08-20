// See http://brunch.io for documentation.
module.exports = {
  notifications: false,
  hot: true,

  plugins: {
    // /////////////
    // JS plugins //
    // /////////////

    // Transpile ES201* features that aren't supported by the latest WebKit
    babel: {
      presets: [
        [
          'env',
          {
            targets: {
              // This seems to map closest to WebKit...
              browsers: ['last 1 safari versions']
            }
          }
        ]
      ]
    },
    // Minify JS
    uglify: {
      toplevel: true,
      mangle: true,
      compress: {
        global_defs: { DEBUG: false }
      }
    },
    // Obfuscate our code for protection against peeping toms
    obfuscator: {
      debugProtection: true,
      disableConsoleOutput: true,
      renameGlobals: true,
      seed: new Date().getTime(),
      transformObjectKeys: true
    },

    // //////////////
    // CSS plugins //
    // //////////////
    sassLint: {
      warnOnly: process.env.NODE_ENV !== 'production'
    },
    csscomb: {
      encoding: 'zen'
    },
    sass: {
      modules: {
        ignore: [/\/index\.scss/]
      }
    },
    postcss: {
      processors: [require('autoprefixer')(['last 1 safari versions'])]
    }
  },

  files: {
    javascripts: { joinTo: 'app.js' },
    stylesheets: { joinTo: 'app.css' }
  }
};
