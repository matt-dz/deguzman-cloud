import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		'https://docs.google.com/forms/d/e/1FAIpQLSd_1wq1Zng_OemUXnVgW1qWhvrlRZoCsRqJrtgLl_gG6PbFvA/viewform'
	);
}
