<script lang="ts">
	import Signup from '$lib/components/signup/Signup.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<svelte:head>
	<title>deguzman | sign up</title>
</svelte:head>

<div class="flex h-dvh items-center">
	<div class="flex w-full flex-col items-center">
		<div class="flex w-full max-w-[525px] items-center">
			<Signup redirectUrl={data.redirectUrl || undefined} />
		</div>
	</div>
</div>
