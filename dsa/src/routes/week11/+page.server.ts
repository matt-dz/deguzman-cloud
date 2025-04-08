import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		'https://docs.google.com/presentation/d/1G_G0vepCkmJLTQ0Dtu3gWbgPFdS8tOYQEVWVk6usKrE/edit?usp=sharing'
	);
}
