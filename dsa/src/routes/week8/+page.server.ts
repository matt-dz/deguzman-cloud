import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		'https://docs.google.com/presentation/d/1ONDt_dJn_ISCR3iWgABeoauVqaY9W80JxOgAckExL5o/edit?usp=sharing'
	);
}
