/** @type {import('tailwindcss').Config} */
// const plugin = require('tailwindcss/plugin')

module.exports = {
  content: [
    './index.html',
    './assets/css/input.css',
    './templates/**/*.html',
    'node_modules/preline/dist/*.js',
    // ... add any other file types as needed
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('preline/plugin')
  ],
}

