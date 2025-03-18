<script lang="ts">
	import Input from '$lib/components/dgc/forms/input/input.svelte';
	import { Label } from '$lib/components/shadcn/ui/label';
	import type { SignupErrors, SignupForm } from '$lib/types/forms';

	let fields: SignupForm = $state({
		email: '',
		password: '',
		firstName: '',
		lastName: ''
	});

	let errors: SignupErrors = $state({
		email: false,
		password: false,
		firstName: false,
		lastName: false
	});

	const handleSubmit = async (event: Event) => {
		event.preventDefault();
		errors.firstName = fields.firstName === '';
		errors.lastName = fields.lastName === '';
		errors.password = fields.password === '';
		errors.email = fields.email === '';

		if (errors.password || errors.email || errors.firstName || errors.lastName) return;

		// const response = await fetch('/api/login', {
		// 	method: 'POST',
		// 	headers: {
		// 		'Content-Type': 'application/json'
		// 	},
		// 	body: JSON.stringify({ email, password })
		// });

		// if (response.ok) {
		// 	const { token } = await response.json();
		// 	localStorage.setItem('token', token);
		// 	window.location.href = '/';
		// } else {
		// 	alert('Invalid email or password');
		// }
	};
</script>

<section class="grid grid-cols-1 items-center justify-items-center h-screen mx-4">
	<form class="w-full max-w-[400px] text-black flex flex-col items-center gap-4">
		<h1 class="text-center text-3xl font-medium">
			Get ready to enter <br />The DeGuzman Cloud
		</h1>

		<div class="w-full">
			<Label for="email">First Name</Label>
			<Input
				bind:value={fields.firstName}
				error={errors.firstName}
				id="first-name"
				type="text"
				placeholder="Matthew"
			/>
		</div>

		<div class="w-full">
			<Label for="email">Last Name</Label>
			<Input
				bind:value={fields.lastName}
				error={errors.lastName}
				id="last-name"
				type="text"
				placeholder="DeGuzman"
			/>
		</div>

		<div class="w-full">
			<Label for="email">Email</Label>
			<Input
				bind:value={fields.email}
				error={errors.email}
				id="email"
				type="text"
				placeholder="matthew@deguzman.cloud"
			/>
		</div>

		<div class="w-full">
			<Label for="password">Password</Label>
			<Input
				bind:value={fields.password}
				error={errors.password}
				id="password"
				type="password"
				placeholder="password"
			/>
		</div>

		<button onclick={handleSubmit} class="mt-4">Sign Up</button>
	</form>
</section>
