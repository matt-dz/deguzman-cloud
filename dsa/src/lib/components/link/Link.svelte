<script lang="ts">
	import type { Snippet } from 'svelte';

	let {
		borderColor = '#ba9ee8',
		hoverColor = '#3776ff',
		href,
		text
	}: { borderColor?: string; hoverColor?: string; href: string; text: Snippet } = $props();
</script>

<a
	{href}
	target="_bank"
	class="underline underline-offset-2 transition-colors duration-300 ease-in-out"
	style="text-decoration-color: {borderColor};"
	onmouseover={(event) => (event.currentTarget.style.textDecorationColor = hoverColor)}
	onmouseout={(event) => (event.currentTarget.style.textDecorationColor = borderColor)}
	onfocus={(event) => (event.currentTarget.style.textDecorationColor = hoverColor)}
	onblur={(event) => (event.currentTarget.style.textDecorationColor = borderColor)}
	onclick={(event) => (event.currentTarget.style.textDecorationColor = borderColor)}
>
	{@render text()}
</a>
