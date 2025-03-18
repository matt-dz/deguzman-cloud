import type { PageServerLoad } from './$types';
import type { Cookies } from '@sveltejs/kit';
import { VITE_BASE_URL } from '$env/static/private';
import { redirect } from '@sveltejs/kit';
import { getSessionCookie } from '$lib/auth/utils';
import { type Deployment } from '$lib/types/deployment';
import { type User } from '$lib/types/user';

export const load: PageServerLoad = async ({ cookies }: { cookies: Cookies }) => {
	const sessionCookie = getSessionCookie(cookies);
	const [userResponse, deploymentsResponse] = await Promise.all([
		fetch(`${VITE_BASE_URL}/api/user`, {
			method: 'GET',
			headers: {
				Cookie: sessionCookie
			}
		}),
		fetch(`${VITE_BASE_URL}/api/deployment/list?list=active&order=desc`, {
			method: 'GET',
			headers: {
				Cookie: sessionCookie
			}
		})
	]);

	if (!userResponse.ok) {
		if (userResponse.status === 401) {
			console.error(userResponse.statusText);
			redirect(303, '/login');
		}
		console.error(userResponse.statusText);
	}

	if (!deploymentsResponse.ok) {
		if (deploymentsResponse.status === 401) {
			console.error(deploymentsResponse.statusText);
			redirect(303, '/login');
		}
		console.error(deploymentsResponse.statusText);
	}

	if (userResponse.ok && deploymentsResponse.ok) {
		const [user, deployments] = await Promise.all([
			userResponse.json(),
			deploymentsResponse.json()
		]);
		return {
			user,
			deployments
		};
	} else if (userResponse.ok) {
		const user = await userResponse.json();
		const deployments: Deployment[] = [];
		return {
			user,
			deployments
		};
	} else if (deploymentsResponse.ok) {
		const deployments = await deploymentsResponse.json();
		const user: User = {
			first_name: '',
			last_name: '',
			email: ''
		};
		return {
			user,
			deployments
		};
	} else {
		const deployments: Deployment[] = [];
		const user: User = {
			first_name: '',
			last_name: '',
			email: ''
		};
		return {
			user,
			deployments
		};
	}
};
