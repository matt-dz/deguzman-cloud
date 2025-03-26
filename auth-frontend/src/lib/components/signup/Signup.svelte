<script lang="ts">
	import type { SignupPayload } from '$lib/auth';
	import { env } from '$env/dynamic/public';

	interface Props {
		redirectUrl?: string;
	}

	let { redirectUrl }: Props = $props();

	let isErrorShaking = $state(false);

	let signupPayload: SignupPayload = $state({
		first_name: '',
		last_name: '',
		email: '',
		password: ''
	});
	let confirmedPassword = $state('');
	let errorText = $state('');

	function setErrorText(text: string) {
		errorText = text;
		isErrorShaking = true;
		setTimeout(() => {
			isErrorShaking = false;
		}, 1000);
	}

	function validatePasswords(): boolean {
		if (confirmedPassword !== signupPayload.password) {
			setErrorText('Passwords do not match');
			return false;
		} else {
			errorText = '';
			return true;
		}
	}

	function clearFields() {
		signupPayload = {
			first_name: '',
			last_name: '',
			email: '',
			password: ''
		};
		confirmedPassword = '';
		errorText = '';
	}

	async function onsubmit(e: Event) {
		e.preventDefault();

		if (!validatePasswords()) return;
		try {
			const loginEndpoint =
				`${env.PUBLIC_BASE_URL}/api/signup` +
				(redirectUrl ? `?redirect=${encodeURIComponent(redirectUrl)}` : '');

			const res = await fetch(loginEndpoint, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				credentials: 'include',
				body: JSON.stringify(signupPayload)
			});

			if (!res.ok) {
				setErrorText(await res.text());
				return;
			}

			clearFields();
			alert('Successfully created user!');
		} catch (e) {
			alert('Uh-oh! Something went wrong...');
			console.error(e);
		}
	}
</script>

<div class="flex w-full flex-col items-center px-12 py-8">
	<div class="w-full">
		<h1 class="text-2xl font-semibold">
			Sign up for the <span class="inline-block italic">DeGuzman Cloud</span>
		</h1>
		<form class="mt-8 flex flex-col items-center gap-4" onsubmit={async (e) => await onsubmit(e)}>
			<label class="flex w-full flex-col gap-2">
				<h1 class="text-sm">First Name</h1>
				<input
					required
					type="text"
					bind:value={signupPayload.first_name}
					class="rounded-lg border-2 border-gray-700 bg-black p-2 focus:border-blue-300 focus:drop-shadow-[0_0_6px_var(--color-blue-300)] focus:outline-none"
				/>
			</label>
			<label class="flex w-full flex-col gap-2">
				<h1 class="text-sm">Last Name</h1>
				<input
					required
					type="text"
					bind:value={signupPayload.last_name}
					class="rounded-lg border-2 border-gray-700 bg-black p-2 focus:border-blue-300 focus:drop-shadow-[0_0_6px_var(--color-blue-300)] focus:outline-none"
				/>
			</label>
			<label class="flex w-full flex-col gap-2">
				<h1 class="text-sm">Email</h1>
				<input
					required
					type="email"
					bind:value={signupPayload.email}
					autocomplete="email"
					class="rounded-lg border-2 border-gray-700 bg-black p-2 focus:border-blue-300 focus:drop-shadow-[0_0_6px_var(--color-blue-300)] focus:outline-none"
				/>
			</label>

			<label class="flex w-full flex-col gap-2">
				<h1 class="text-sm">Password</h1>
				<input
					required
					type="password"
					autocomplete="new-password"
					bind:value={signupPayload.password}
					class="rounded-lg border-2 border-gray-700 bg-black p-2 focus:border-blue-300 focus:drop-shadow-[0_0_6px_var(--color-blue-300)] focus:outline-none"
				/>
			</label>

			<label class="flex w-full flex-col gap-2">
				<h1 class="text-sm">Confirm Password</h1>
				<input
					required
					type="password"
					autocomplete="new-password"
					bind:value={confirmedPassword}
					class:!border-red-400={confirmedPassword && confirmedPassword !== signupPayload.password}
					onchange={validatePasswords}
					class="rounded-lg border-2 border-gray-700 bg-black p-2 focus:border-blue-300 focus:drop-shadow-[0_0_6px_var(--color-blue-300)] focus:outline-none"
				/>
			</label>

			<button
				class="transition-duration-200 mt-4 w-full cursor-pointer rounded-lg border-2 border-solid border-white bg-black px-4 py-2 text-violet-300 drop-shadow-[0_0_6px_#fff] transition-colors hover:bg-stone-700"
			>
				<span class="btn-text text-sm text-white sm:text-base">Sign up</span>
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
