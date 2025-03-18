<script lang="ts">
	import { goto } from '$app/navigation';
	import Input from '$lib/components/dgc/forms/input/input.svelte';
	import { Label } from '$lib/components/shadcn/ui/label';
	import type { LoginErrors, LoginForm } from '$lib/types/forms';

	let fields: LoginForm = $state({
		email: '',
		password: ''
	});

	let errors: LoginErrors = $state({
		email: false,
		password: false
	});

	let errMsg: string = $state('');
	const invisibleText = 'invisible';

	const handleSubmit = async (event: Event) => {
		event.preventDefault();
		errors.password = fields.password === '';
		errors.email = fields.email === '';

		if (errors.password || errors.email) return;
		const response = await fetch(`${import.meta.env.VITE_BASE_URL}/api/login`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			credentials: 'include',
			body: JSON.stringify(fields)
		});

		if (response.status === 401) {
			errMsg = 'Invalid email or password';
		} else if (!response.ok) {
			errMsg = 'Error logging in. Try again.';
			console.error(response.statusText);
		} else {
			errMsg = '';
			await goto('/dashboard');
		}
	};
</script>

<svelte:head>
	<title>deguzman | login</title>
</svelte:head>

<section class="grid grid-cols-1 items-center justify-items-center h-full mx-4">
	<form class="w-full max-w-[400px] flex flex-col items-center gap-4">
		<h1 class="text-center text-6xl font-light">login</h1>

		<div class="w-full">
			<Label for="email">email</Label>
			<Input
				bind:value={fields.email}
				error={errors.email}
				id="email"
				type="text"
				placeholder="matthew@deguzman.cloud"
			/>
		</div>

		<div class="w-full">
			<Label for="password">password</Label>
			<Input
				bind:value={fields.password}
				error={errors.password}
				id="password"
				type="password"
				placeholder="password"
			/>
		</div>

		<button onclick={handleSubmit} class="mt-4 text-lg relative"> Enter </button>
		<h1 class:text-transparent={!errMsg} class="text-red-400">
			{errMsg ? errMsg : invisibleText}
		</h1>
	</form>
</section>

<style lang="postcss">
	button {
		@apply hover:bg-white w-fit hover:text-black px-2 py-1 transition-colors;
	}
</style>
