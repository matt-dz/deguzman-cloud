import { redirect, type Handle } from '@sveltejs/kit';
import { AuthRole, sessionCookieName } from '$lib/auth';
import { env as privateEnv } from '$env/dynamic/private';
import { env as publicEnv } from '$env/dynamic/public';

export const handle: Handle = async ({ event, resolve }) => {
	const path = event.url.pathname;
	console.log(path);

	if (path.includes('/signup')) {
		const loginUrl = `/?redirect=${encodeURIComponent('/signup')}`;
		const sessionCookie = event.cookies.get(sessionCookieName);

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
			body: JSON.stringify({ role: AuthRole.Admin })
		});

		if (resp.status === 200) {
			return await resolve(event);
		}

		console.error(await resp.text());
		redirect(303, loginUrl);
	}

	return await resolve(event);
};
