<script lang="ts">
	import type { Snippet } from 'svelte';

	let {
		title,
		link,
		href,
		description
	}: { title: string; link: string; href: string; description: Snippet } = $props();
</script>

<div>
	<div class="flex items-center">
		<h1 class="text-2xl md:text-3xl font-semibold">{title}</h1>
		<span class="text-3xl md:text-4xl font-light px-2">|</span>
		<a class="italic text-2xl tablet:text-3xl font-light hover:underline" {href}>{link}</a>
	</div>
	{@render description()}
</div>
