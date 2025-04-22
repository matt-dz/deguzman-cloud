import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		'https://docs.google.com/presentation/d/1BclCucahDNZLPs_EwIz0J-SW8XRB2sZf2x9y5TjDUyo/edit?usp=sharing'
	);
}
