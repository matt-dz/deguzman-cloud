import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		'https://docs.google.com/presentation/d/1SKauPvsXgH6VXktNGxwmlc8C0m5byvu_IYKe1cP9v9Y/edit?usp=sharing'
	);
}
