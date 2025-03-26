import { redirect, type Handle } from '@sveltejs/kit';
import { type AuthPayload, AuthRole } from '$lib/auth';
import { env as privateEnv } from '$env/dynamic/private';
import { env as publicEnv } from '$env/dynamic/public';

export const handle: Handle = async ({ event, resolve }) => {
	const path = event.url.pathname;
	console.log(path);

	if (path.includes('/home')) {
		const loginUrl = `${publicEnv.PUBLIC_AUTH_FRONTEND_URL}?redirect=${encodeURIComponent(publicEnv.PUBLIC_BASE_URL + '/home')}`;
		const sessionCookie = event.cookies.get('session');

		// Redirect unauthenticated users to the login page
		if (!sessionCookie) {
			redirect(303, loginUrl);
		}

		const resp = await fetch(`${privateEnv.AUTH_INTERNAL_URL}/api/auth`, {
			method: 'POST',
			credentials: 'include',
			headers: {
				Origin: publicEnv.PUBLIC_BASE_URL,
				cookie: `session=${sessionCookie}`
			},
			body: JSON.stringify({ role: AuthRole.Admin } as AuthPayload)
		});

		if (resp.status === 200) {
			return await resolve(event);
		}
		console.error(await resp.text());
		redirect(303, loginUrl);
	}

	return await resolve(event);
};
