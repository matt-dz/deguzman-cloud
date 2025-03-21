import { redirect, type Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	const path = event.url.pathname;
	console.log(path);

	if (path.includes('/home')) {
		const loginUrl = `${import.meta.env.VITE_AUTH_URL}?redirect=${encodeURIComponent(import.meta.env.VITE_BASE_URL + '/home')}`;
		const sessionCookie = event.cookies.get('session');

		// Redirect unauthenticated users to the login page
		if (!sessionCookie) {
			redirect(303, loginUrl);
		}

		const resp = await fetch(`${import.meta.env.VITE_AUTH_BACKEND_URL}/api/auth`, {
			method: 'POST',
			credentials: 'include',
			headers: {
				Origin: import.meta.env.VITE_BASE_URL,
				cookie: `session=${sessionCookie}`
			}
		});

		if (resp.status === 200) {
			return await resolve(event);
		}
		console.error(await resp.text());
		redirect(303, loginUrl);
	}

	return await resolve(event);
};
