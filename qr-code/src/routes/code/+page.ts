import type { PageLoad } from './$types';

export const load: PageLoad = async ({ url }: { url: URL }) => {
	const title = url.searchParams.get('title') ?? '';
	const description = url.searchParams.get('description') ?? '';
	const urlQuery = url.searchParams.get('url') ?? '';
	return {
		url: urlQuery,
		title,
		description
	};
};
