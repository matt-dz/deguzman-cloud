<script lang="ts">
	import type { Snippet } from 'svelte';
	// import type { Resource } from '$lib/types';
	import Link from '$lib/components/link/Link.svelte';

	let {
		name = '',
		// resources = [],
		// problems = [],
		extra,
		title,
		description,
		pptLink
	}: {
		name?: string;
		// resources?: Resource[];
		// problems?: Resource[];
		extra?: Snippet;
		title: string;
		description: Snippet;
		pptLink: string;
	} = $props();

	if (!name) {
		name = title;
	}
</script>

<div class="collapse collapse-arrow bg-base-200">
	<input type="checkbox" {name} />
	<div class="accordion-header collapse-title text-xl font-extralight">{title}</div>
	<div class="collapse-content flex flex-col gap-2">
		{@render description()}

		<Link href={pptLink}
			>{#snippet text()}
				<p class="font-light">Discussion Slides</p>
			{/snippet}
		</Link>

		<!-- {#if resources.length > 0}
			<h1 class="text-xl">Resources</h1>
			<ul class="list-disc">
				{#each resources as resource}
					<Link href={resource.link}>
						{#snippet text()}
							{resource.title}
						{/snippet}
					</Link>
				{/each}
			</ul>
		{/if}

		{#if problems.length > 0}
			<h1>Practice Problems</h1>
			<ul class="list-disc">
				{#each problems as problem}
					<Link href={problem.link}>
						{#snippet text()}
							{`${problem.title} (${problem.difficulty})`}
						{/snippet}
					</Link>
				{/each}
			</ul>
		{/if} -->

		{#if extra}
			{@render extra()}
		{/if}
	</div>
</div>
