import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ url }: { url: URL }) => {
	const { pathname } = url;

	return {
		pathname
	};
};
