import adapter from '@sveltejs/adapter-node'; // replace with '@sveltejs/adapter-node' if you're not self-hosting
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: vitePreprocess(),

	kit: {
		adapter: adapter({
			out: 'build' // remove if you're not self hosting
		})
	}
};

export default config;
