<script lang="ts">
	interface Props {
		title: string;
		description: string;
		direction?: string;
		children?: import('svelte').Snippet;
	}

	let {
		title,
		description,
		direction = 'left',
		children
	}: Props = $props();
</script>

<div
	class="max-w-[430px] w-full flex items-center gap-16"
	class:flex-row-reverse={direction === 'right'}
>
	<div
		class="flex-grow"
		class:text-left={direction === 'left'}
		class:text-right={direction !== 'left'}
	>
		<h1 class="text-2xl font-bold text-black">{title}</h1>
		<p>
			{description}
		</p>
	</div>
	{@render children?.()}
</div>
