/** @type {import('tailwindcss').Config} */
export default {
  corePlugins: {
    preflight: false
  },
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}'
  ],
  theme: {
    extend: {
      animation: {
        blink: 'blink 1.25s ease-in-out infinite',
      },
      keyframes: {
        blink: {
          '0%, 100%': { 'transform': 'scaleX(1)' },
          '50%': { 'transform': 'scale3d(1.25, 1.25, 1)' },
        },
      }
    },
  },
  plugins: []
}

