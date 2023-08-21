/** @type {import('tailwindcss').Config} */

const withMT = require("@material-tailwind/html/utils/withMT");

module.exports =  withMT({
  content: ["./groceryhelper/*.html"],
  theme: {
    extend: {},
  },
  plugins: [],
});

