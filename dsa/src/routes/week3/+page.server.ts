import { redirect } from '@sveltejs/kit';

export function load() {
	throw redirect(
		307,
		'https://docs.google.com/presentation/d/1BJCV0xbhYQweHNKt5Cx7r-8tIpKw9Ndt/edit?usp=sharing&ouid=101706383218741169050&rtpof=true&sd=true'
	);
}
