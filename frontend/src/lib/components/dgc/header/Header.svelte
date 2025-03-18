<script lang="ts">
	import { goto } from '$app/navigation';

	async function logout() {
		const resp = await fetch(`${import.meta.env.VITE_BASE_URL}/api/logout`, {
			method: 'GET',
			credentials: 'include'
		});
		if (!resp.ok) {
			console.error('Error logging out', resp.statusText);
		}
		await goto('/login');
	}

	let { isLoggedIn = false } = $props();
</script>

<div class="w-full flex flex-col items-center mt-2">
	<header class="w-full max-w-[1080px]">
		<div class="flex flex-row-reverse grow gap-6 text-lg">
			{#if isLoggedIn}
				<button onclick={logout}>Log out</button>
				<a href="/dashboard">Dashboard</a>
			{:else}
				<a href="/login">Login</a>
			{/if}
			<a href="/" class="flex-grow">deguzman.cloud</a>
		</div>
	</header>
</div>
