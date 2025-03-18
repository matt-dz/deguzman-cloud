import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		'https://docs.google.com/presentation/d/1g95fl4lzjhVttD3vFRIj9rGXw_teVy0nj7af_hHynXM/edit?usp=sharing'
	);
}
