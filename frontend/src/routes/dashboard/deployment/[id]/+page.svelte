<script lang="ts">
	import type { PageData } from './$types';
	import type { Deployment } from '$lib/types/deployment';
	import {
		IconBlocks,
		IconBrandCodesandbox,
		IconBrandAppleArcade,
		IconDatabase
	} from '@tabler/icons-svelte';
	import Status from '$lib/components/dgc/deployment/Status.svelte';

	let { data }: { data: PageData } = $props();
	const deployment: Deployment = data.deployment;
	const iconWidth = 27;

	function parseArgs(args: string[]) {
		return `[${args
			.map((arg: string) => {
				return `"${arg.replaceAll('"', '\\"')}"`;
			})
			.join(', ')}]`;
	}
</script>

<section class="mt-16">
	<div class="w-full max-w-[700px] flex">
		<h1 class="text-3xl font-medium flex-grow">{deployment.config.name}</h1>
		<Status status={deployment.status} />
	</div>
</section>

<section class="mt-32">
	<div class="w-full max-w-[700px] flex flex-col gap-4">
		<div>
			<div class="deployment-header-container">
				<IconBlocks stroke={1} color="green" size={iconWidth} />
				<h1 class="deployment-header">Build</h1>
			</div>
			<div class="deployment-line"></div>

			<h2 class="text-lg font-light">Image</h2>
			<p class="font-extralight">{deployment.config.build.image}</p>

			<h2 class="text-lg font-light">Arguments</h2>
			<p class="font-extralight">{parseArgs(deployment.config.build.arguments)}</p>
		</div>

		<div>
			<div class="deployment-header-container">
				<IconBrandCodesandbox stroke={1} color="blue" size={iconWidth} />
				<h1 class="deployment-header">Environment Variables</h1>
			</div>
			<div class="deployment-line"></div>
			<div class="grid grid-cols-[1fr_2fr]">
				<h1 class="text-lg font-light">Key</h1>
				<h1 class="text-lg font-light">Value</h1>
				{#each Object.entries(deployment.config.env) as [key, value]}
					<p class="text-sm font-extralight">{key}</p>
					<p class="text-sm font-extralight">{value}</p>
				{/each}
			</div>
		</div>

		<div>
			<div class="deployment-header-container">
				<IconBrandAppleArcade stroke={1} color="orange" size={iconWidth} />
				<h1 class="deployment-header">Services</h1>
			</div>
			<div class="deployment-line"></div>
			<div class="grid grid-cols-[1fr_2fr]">
				<h1 class="text-lg font-light">Port</h1>
				<h1 class="text-lg font-light">Target Port</h1>
				{#each Object.entries(deployment.config.services) as [key, value]}
					<p class="text-sm font-extralight">{key}</p>
					<p class="text-sm font-extralight">{value}</p>
				{/each}
			</div>
		</div>

		<div>
			<div class="deployment-header-container">
				<IconDatabase stroke={1} color="purple" size={iconWidth} />
				<h1 class="deployment-header">Mounts</h1>
			</div>
			<div class="deployment-line"></div>
			{#each deployment.config.mounts as mount}
				<h1 class="text-lg font-light">{mount}</h1>
			{/each}
		</div>
	</div>
</section>

<style lang="postcss">
	section {
		@apply w-full flex flex-col items-center;
	}

	.deployment-header {
		@apply text-2xl font-medium;
	}

	.deployment-line {
		@apply w-full border-t-0 border-solid border-black;
	}

	.deployment-header-container {
		@apply flex items-center gap-2 ml-[-38px];
	}
</style>
