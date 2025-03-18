import { redirect, type Handle } from '@sveltejs/kit';
import { VITE_BASE_URL } from '$env/static/private';
import { getSessionCookie } from '$lib/auth/utils';

export const handle: Handle = async ({ event, resolve }) => {
	const path = event.url.pathname;
	console.log(path);

	if (path.includes('/dashboard')) {
		const sessionCookie = getSessionCookie(event.cookies);

		if (sessionCookie) {
			const resp = await fetch(`${VITE_BASE_URL}/api/auth`, {
				method: 'GET',
				headers: {
					Cookie: sessionCookie
				}
			});

			if (resp.status === 200) {
				return await resolve(event);
			}
			console.error(resp.statusText);
		}

		redirect(303, '/login');
	}

	return await resolve(event);
};
