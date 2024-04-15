/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{vue,js,ts,svg}'],
  plugins: [require('daisyui')],
  daisyui: {
    themes: [
      {
        light: {
          ...require("daisyui/src/theming/themes")["light"],
          "--rounded-btn": ".25rem",
          "--rounded-box": ".5rem",
          "--tab-radius": ".5rem",
        },
        dark: {
          ...require("daisyui/src/theming/themes")["dark"],
          "--rounded-btn": ".25rem",
          "--rounded-box": ".5rem",
          "--tab-radius": ".5rem",
        },
      }
    ]
  }
}