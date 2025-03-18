<script lang="ts">
	import '../app.css';
	import Header from '$lib/components/header/Header.svelte';
	let { children } = $props();

	const description =
		'Discussion board for COP3530 - Data Structures and Algorithms at the University of Florida';
</script>

<Header />
<svelte:head>
	<title>COP3530 | Discussion</title>
	<meta name="author" content="Matthew DeGuzman" />
	<meta name="description" content={description} />

	<meta property="og:title" content="COP3530 Discussion" />
	<meta property="og:description" content={description} />
	<meta property="og:image" content="https://dsa.deguzman.cloud/og-image.png" />
	<meta property="og:url" content="https://dsa.deguzman.cloud" />
	<meta property="og:type" content="website" />

	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:site" content="@matthew_d13" />
</svelte:head>
{@render children()}
