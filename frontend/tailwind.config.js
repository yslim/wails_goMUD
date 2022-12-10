const colors = require ('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,ts}",
  ],
  theme: {
    extend: {
      colors: {
        'mdb-blue': '#0d6efd',
        'mdb-indigo': '#6610f2',
        'mdb-purple': '#6f42c1',
        'mdb-pink': '#d63384',
        'mdb-red': '#dc3545',
        'mdb-orange': '#fd7e14',
        'mdb-yellow': '#ffc107',
        'mdb-green': '#198754',
        'mdb-teal': '#20c997',
        'mdb-cyan': '#0dcaf0',
        'mdb-white': '#fff',
        'mdb-gray': '#757575',
        'mdb-gray-dark': '#4f4f4f',
        'mdb-primary': '#1266f1',
        'mdb-secondary': '#b23cfd',
        'mdb-success': '#00b74a',
        'mdb-info': '#39c0ed',
        'mdb-warning': '#ffa900',
        'mdb-danger': '#f93154',
        'mdb-light': '#fbfbfb',
        'mdb-dark': '#262626',

        red: {
          0: '#ff0000',
        },
        blue: {
          0: '#0000ff',
        },
        green: {
          0: '#00ff00',
        },
        yellow: {
          0: '#ffff00',
        },
        cyan: {
          ...colors.cyan,
          0: '#00ffff',
        },
        fuchsia: {
          ...colors.fuchsia,
          0: '#ff00ff'
        },
      },
      width: {
        'fit-content': 'fit-content',
      }
    },
  },
  plugins: [
    require('daisyui'),
  ],
  daisyui: {},
}
