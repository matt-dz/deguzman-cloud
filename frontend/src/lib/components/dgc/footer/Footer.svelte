<script lang="ts">
	import { env } from '$env/dynamic/public';
	interface Props {
		loggedIn: boolean;
	}

	let { loggedIn }: Props = $props();

	async function logout() {
		try {
			const res = await fetch(`${env.PUBLIC_AUTH_BACKEND_URL}/api/logout`, {
				method: 'POST',
				credentials: 'include'
			});
			if (!res.ok) {
				console.error('Error logging out', res.status, res.text);
			}
			window.location.href = '/';
		} catch (e) {
			alert('Uh-oh! Something went wrong...');
			console.error('Error logging out', e);
		}
	}
</script>

<div class="flex flex-col px-4">
	<a class="nav-container footer-underline w-fit" href="/">home</a>
	<a class="nav-container footer-underline w-fit" href="/services">services</a>
	<a class="nav-container footer-underline" target="_blank" href="https://about.deguzman.cloud"
		>about</a
	>
	{#if loggedIn}
		<button class="nav-container footer-underline" onclick={logout}>logout</button>
	{:else}
		<a class="nav-container footer-underline" href={env.PUBLIC_AUTH_FRONTEND_URL}>login</a>
	{/if}
</div>

<style lang="postcss">
	a,
	button {
		@apply text-lg md:text-2xl font-light;
	}

	.footer-underline {
		@apply underline decoration-1 underline-offset-[3px];
	}

	.nav-container {
		@apply hover:bg-white w-fit hover:text-black p-1 md:p-2 transition-colors;
	}
</style>
