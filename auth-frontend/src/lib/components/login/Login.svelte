<script lang="ts">
	import type { LoginResponse } from '$lib/types';

	interface Props {
		redirectUrl?: string;
	}

	let { redirectUrl }: Props = $props();

	let isErrorShaking = $state(false);

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
			console.log(loginEndpoint);
			const res = await fetch(loginEndpoint, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				credentials: 'include',
				body: JSON.stringify({ email, password })
			});

			if (!res.ok) {
				errorText = await res.text();
				isErrorShaking = true;
				setTimeout(() => {
					isErrorShaking = false;
				}, 1000);
				return;
			}

			resBody = await res.json();
			window.location.href = redirectUrl ?? import.meta.env.VITE_HOME_URL;
		} catch (e) {
			alert('Uh-oh! Something went wrong...');
			console.error(e);
		}
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
			<p class:shake-text={isErrorShaking} class="text-red-400" class:opacity-0={!errorText}>
				{errorText || 'placeholder'}
			</p>
		</form>
	</div>
</div>

<style>
	@keyframes shake {
		0% {
			transform: translateX(0);
		}
		25% {
			transform: translateX(-10px);
		}
		50% {
			transform: translateX(10px);
		}
		75% {
			transform: translateX(-10px);
		}
		100% {
			transform: translateX(0);
		}
	}

	.shake-text {
		display: inline-block;
		animation: shake 0.5s ease-in-out;
	}
</style>
