<script lang="ts">
	import '../app.css';
	import { fly, fade } from 'svelte/transition';
	import { quadOut } from 'svelte/easing';
	import type { LayoutData } from './$types';
	import type { Snippet } from 'svelte';

	let { data, children }: { data: LayoutData; children: Snippet } = $props();
	const pathname = $derived(data.pathname);

	const description =
		'QR Code Generator for events, clubs, and classes at UF - The University of Florida';
</script>

<svelte:head>
	<title>QR Code Generator</title>
	<meta name="author" content="Matthew DeGuzman" />
	<meta name="description" content={description} />

	<meta property="og:title" content="QR Code Generator" />
	<meta property="og:description" content={description} />
	<meta property="og:image" content="https://qr.deguzman.cloud/og-image.png" />
	<meta property="og:url" content="https://qr.deguzman.cloud" />
	<meta property="og:type" content="website" />

	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:site" content="@matthew_d13" />
</svelte:head>
{#key pathname}
	<div in:fly={{ delay: 250, duration: 500, y: 50, easing: quadOut }} out:fade={{ duration: 200 }}>
		{@render children()}
	</div>
{/key}
