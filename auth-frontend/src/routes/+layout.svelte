<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<svelte:head>
	<!-- SEO Meta Tags -->
	<meta name="description" content="deguzman.cloud log in page." />
	<meta name="author" content="Matthew DeGuzman" />

	<!-- Open Graph / Facebook -->
	<meta property="og:title" content="deguzman.cloud" />
	<meta property="og:description" content="deguzman.cloud log in page." />
	<meta property="og:image" content="https://deguzman.cloud/assets/logos/og-image.png" />
	<meta property="og:url" content="https://auth.deguzman.cloud" />
	<meta property="og:type" content="website" />

	<!-- Twitter Card -->
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:site" content="@matthew_d13" />
</svelte:head>

{@render children()}
