import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		"https://docs.google.com/presentation/d/10UCZmsEmkvm1EDEbkpLQU3r-4Sm7zPDUzkLD2kaJcZg/"
	);
}
