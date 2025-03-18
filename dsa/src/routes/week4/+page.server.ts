import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		'https://docs.google.com/presentation/d/1SGjaeApZLrveO4CkKODq5qO72WVHeOAvq4IB9hM2E-I/edit?usp=sharing'
	);
}
