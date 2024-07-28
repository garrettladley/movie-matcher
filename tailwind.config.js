/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./**/*.html",
        "./**/*.templ",
        "./**/*.go",
        "./node_modules/flowbite/**/*.js",
    ],
    safelist: [],
    plugins: [
        require('flowbite/plugin')({
            charts: true,
        })
    ],
 }