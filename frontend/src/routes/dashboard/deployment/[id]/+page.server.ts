import { error } from '@sveltejs/kit';
import { getSessionCookie } from '$lib/auth/utils';
import { VITE_BASE_URL } from '$env/static/private';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies, params }) => {
	const sessionCookie = getSessionCookie(cookies);

	let resp: Response;
	try {
		resp = await fetch(`${VITE_BASE_URL}/api/deployment/${params.id}`, {
			method: 'GET',
			headers: {
				Cookie: sessionCookie
			}
		});
	} catch (err) {
		console.error(err);
		error(500, 'An error occurred');
	}

	if (resp.status === 404) {
		console.error(resp.statusText);
		error(404, 'Not found');
	} else if (!resp.ok) {
		console.error(resp.statusText);
		error(500, 'An error occurred');
	}

	try {
		const deployment = await resp.json();
		return {
			deployment
		};
	} catch (err) {
		console.error(err);
		error(500, 'Unable to parse response');
	}
};
