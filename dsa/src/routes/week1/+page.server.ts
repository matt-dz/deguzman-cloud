import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		'https://ufl.instructure.com/courses/525732/files/folder/Slides%20for%20Inperson%20Lectures/Discussion%20Presentations?preview=93479607'
	);
}
