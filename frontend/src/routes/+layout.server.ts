import type { LayoutServerLoad } from './$types';
import type { Cookies } from '@sveltejs/kit';
import { sessionCookieName } from '$lib/auth';

export const load: LayoutServerLoad = ({ url, cookies }: { url: URL; cookies: Cookies }) => {
	return {
		pathname: url.pathname,
		loggedIn: !!cookies.get(sessionCookieName)
	};
};
