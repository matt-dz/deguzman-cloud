import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		'https://docs.google.com/presentation/d/1pXlepCty9XjII8XISoeBy_vM0A4USPznhSYxAY0lA0M/edit?usp=sharing'
	);
}
