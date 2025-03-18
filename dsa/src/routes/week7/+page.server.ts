import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		'https://docs.google.com/presentation/d/1DXYDmENGEAaiuQMrsta6MlvZ4v3MFMH2q3dk8q7CwV4/edit'
	);
}
