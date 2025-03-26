<script lang="ts">
	import Login from '$lib/components/login/Login.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<svelte:head>
	<title>deguzman | log in</title>
</svelte:head>

<div class="flex h-dvh items-center">
	<div class="flex w-full flex-col items-center">
		<div class="flex w-full max-w-[525px] items-center">
			<Login redirectUrl={data.redirectUrl || undefined} />
		</div>
		<!-- <div class="grid h-full grid-cols-1 md:grid-cols-2">
			<div class="hidden md:visible md:flex md:items-center"></div>
			<div class="flex w-full items-center md:border-l-2 md:border-solid md:border-gray-700">
				<Login redirectUrl={data.redirectUrl || undefined} />
			</div>
		</div> -->
	</div>
</div>
