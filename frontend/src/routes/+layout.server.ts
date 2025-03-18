import type { LayoutServerLoad } from './$types';
import type { Cookies } from '@sveltejs/kit';
import { sessionCookieName } from '$lib/stores/metadata';

export const load: LayoutServerLoad = ({ url, cookies }: { url: URL; cookies: Cookies }) => {
	let isLoggedIn = false;
	sessionCookieName.subscribe((name: string) => {
		isLoggedIn = !!cookies.get(name);
	});
	const { pathname } = url;
	return {
		pathname,
		isLoggedIn
	};
};
