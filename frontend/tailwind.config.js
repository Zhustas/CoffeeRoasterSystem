import containerQueries from '@tailwindcss/container-queries';
import forms from '@tailwindcss/forms';
import typography from '@tailwindcss/typography';

/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
			colors: {
				page: '#FFFDBB',
				// main: '#F0DCD0',
				main: '#FCEDE3',
				button: '#EABB9F',
				// 'hover-button': '#C79B81'
				'hover-button': '#D4A587'
			},
			keyframes: {
				slide: {
					'0%': { transform: 'translate(-100%)' },
					'100%': { transform: 'translate(100%)' }
				},
				'slide-message': {
					'0%': { transform: 'translate(200%)' },
					'100%': { transform: '0' }
				}
			},
			animation: {
				slide: 'slide 3s linear infinite',
				'slide-message': 'slide-message 0.5s ease-out'
			}
		}
	},

	plugins: [typography, forms, containerQueries]
};
