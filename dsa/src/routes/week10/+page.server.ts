import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		'https://docs.google.com/presentation/d/17oYH9KxF9JlVbaPuDo5fABLkLoG2IJocHQd8B63mmuY/edit?usp=sharing'
	);
}
