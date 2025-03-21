import type { PageLoad } from './$types';

export const load: PageLoad = ({ url }: { url: URL }) => {
	const redirect = url.searchParams.get('redirect');
	return {
		redirectUrl: redirect
	};
};
