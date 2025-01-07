import type { PageLoad } from './$types';
export const load: PageLoad = async ({ url, params }) => {
	const title = url.searchParams.get('title') ?? '';
	const description = url.searchParams.get('description') ?? '';
	const id = params.id;
	return {
		id,
		title,
		description
	};
};
