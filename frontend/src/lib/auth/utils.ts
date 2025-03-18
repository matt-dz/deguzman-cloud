import { sessionCookieName } from '$lib/stores/metadata';
import { type Cookies } from '@sveltejs/kit';

export function getSessionCookie(cookies: Cookies) {
	let cookieHeader = '';
	sessionCookieName.subscribe((name) => {
		const sessionId = cookies.get(name) ?? '';
		if (sessionId) {
			cookieHeader = `${name}=${sessionId}`;
		}
	});
	return cookieHeader;
}
