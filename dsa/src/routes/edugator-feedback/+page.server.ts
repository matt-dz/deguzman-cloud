import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(307, 'https://forms.gle/YYKBZGmW2A7jypLL6');
}
