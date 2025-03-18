<script lang="ts">
	import { createBubbler, passive } from 'svelte/legacy';

	const bubble = createBubbler();
	import type { HTMLInputAttributes } from 'svelte/elements';
	import type { InputEvents } from '$lib/components/shadcn/ui/input/index.js';
	import { cn } from '$lib/components/shadcn/utils.js';

	type $$Props = HTMLInputAttributes & { error: boolean };
	type $$Events = InputEvents;

	// Workaround for https://github.com/sveltejs/svelte/issues/9305

	interface Props {
		class?: $$Props['class'];
		error?: boolean;
		value?: $$Props['value'];
		// Fixed in Svelte 5, but not backported to 4.x.
		readonly?: $$Props['readonly'];
		[key: string]: any;
	}

	let {
		class: className = undefined,
		error = false,
		value = $bindable(undefined),
		readonly = undefined,
		...rest
	}: Props = $props();
</script>

<input
	class={cn(
		'placeholder:text-gray-400 focus-visible:ring-ring flex h-9 w-full rounded-[10px] border border-[#777777] px-3 py-1 text-sm transition-colors file:border-0 bg-[#353535] file:bg-transparent file:text-sm file:font-medium focus-visible:outline-none focus-visible:ring-1 disabled:cursor-not-allowed disabled:opacity-50',
		className
	)}
	class:error
	bind:value
	{readonly}
	onblur={bubble('blur')}
	onchange={bubble('change')}
	onclick={bubble('click')}
	onfocus={bubble('focus')}
	onfocusin={bubble('focusin')}
	onfocusout={bubble('focusout')}
	onkeydown={bubble('keydown')}
	onkeypress={bubble('keypress')}
	onkeyup={bubble('keyup')}
	onmouseover={bubble('mouseover')}
	onmouseenter={bubble('mouseenter')}
	onmouseleave={bubble('mouseleave')}
	onmousemove={bubble('mousemove')}
	onpaste={bubble('paste')}
	oninput={bubble('input')}
	use:passive={['wheel', () => bubble('wheel')]}
	{...rest}
/>

<style lang="postcss">
	.error {
		@apply border-red-700;
	}
</style>
