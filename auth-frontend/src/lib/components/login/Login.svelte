<script lang="ts">
	import type { LoginResponse } from '$lib/types';

	interface Props {
		redirectUrl?: string;
	}

	let { redirectUrl }: Props = $props();

	let email = $state('');
	let password = $state('');
	let errorText = $state('');

	async function onsubmit(e: Event) {
		e.preventDefault();
		let resBody: LoginResponse | null = null;
		try {
			const loginEndpoint =
				`${import.meta.env.VITE_BASE_URL}/api/login` +
				(redirectUrl ? `?redirect=${encodeURIComponent(redirectUrl)}` : '');
			const res = await fetch(loginEndpoint, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email, password })
			});
			console.log(res);

			if (!res.ok) {
				errorText = res.statusText;
				return;
			}

			resBody = await res.json();
		} catch (e) {
			alert('Uh-oh! Something went wrong...');
			console.error(e);
		}

		window.location.href = resBody ? resBody.redirect : '/home';
	}
</script>

<div class="flex w-full flex-col items-center px-12 py-8">
	<div class="w-full max-w-[700px]">
		<h1 class="text-2xl font-semibold">Log in to the DeGuzman Cloud</h1>
		<form class="mt-8 flex flex-col items-center gap-4" onsubmit={async (e) => await onsubmit(e)}>
			<label class="flex w-full flex-col gap-2">
				<h1 class="text-sm">Email</h1>
				<input
					required
					type="email"
					autocomplete="email"
					bind:value={email}
					class="rounded-lg border-2 border-gray-700 bg-black p-2 focus:border-blue-300 focus:drop-shadow-[0_0_6px_var(--color-blue-300)] focus:outline-none"
				/>
			</label>

			<label class="flex w-full flex-col gap-2">
				<h1 class="text-sm">Password</h1>
				<input
					required
					autocomplete="current-password"
					type="password"
					bind:value={password}
					class="rounded-lg border-2 border-gray-700 bg-black p-2 focus:border-blue-300 focus:drop-shadow-[0_0_6px_var(--color-blue-300)] focus:outline-none"
				/>
			</label>

			<button
				class="transition-duration-200 mt-4 w-full cursor-pointer rounded-lg border-2 border-solid border-white bg-black px-4 py-2 text-violet-300 drop-shadow-[0_0_6px_#fff] transition-colors hover:bg-stone-700"
			>
				<span class="btn-text text-sm text-white sm:text-base">Log In</span>
			</button>
		</form>
	</div>
</div>
