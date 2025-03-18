<script lang="ts">
	import '../app.css';
	import Footer from '$lib/components/dgc/footer/Footer.svelte';

	interface Props {
		data: { pathname: string; isLoggedIn: boolean };
		children?: import('svelte').Snippet;
	}

	let { data, children }: Props = $props();
	const pathname = $derived(data.pathname);
	// const isLoggedIn = $derived(data.isLoggedIn);
</script>

<svelte:head>
	<!-- SEO Meta Tags -->
	<meta
		name="description"
		content="deguzman.cloud is the latest in cloud innovations and is the only solution that can handle the most demanding and cutting edge applications with no downtime with a simple intuitive ui."
	/>
	<meta name="author" content="Matthew DeGuzman" />

	<!-- Open Graph / Facebook -->
	<meta property="og:title" content="deguzman.cloud" />
	<meta
		property="og:description"
		content="deguzman.cloud is the latest in cloud innovations and is the only solution that can handle the most demanding and cutting edge applications with no downtime with a simple intuitive ui."
	/>
	<meta property="og:image" content="https://deguzman.cloud/assets/logos/og-image.png" />
	<meta property="og:url" content="https://deguzman.cloud" />
	<meta property="og:type" content="website" />

	<!-- Twitter Card -->
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:site" content="@matthew_d13" />
</svelte:head>

<div class="flex flex-col items-center h-dvh">
	<div class="w-full max-w-[1000px] h-full flex flex-col pb-8">
		{#key pathname}
			<div class="flex-grow">
				{@render children?.()}
			</div>
		{/key}
		<Footer />
	</div>
</div>
