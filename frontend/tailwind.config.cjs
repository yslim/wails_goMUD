/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./index.html', './src/**/*.{svelte,js,ts,html}'],
    theme: {
        extend: {
            colors: {
                cyan: '#00ffff',
            },
        },
    },
    plugins: [
        require('daisyui'),
    ],
    daisyui: {},
}
