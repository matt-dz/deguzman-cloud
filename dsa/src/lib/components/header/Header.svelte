<script lang="ts">
	let innerWidth = 0;
</script>

<svelte:window bind:innerWidth />
<header class="mt-4 flex flex-col items-center">
	<a href="/" class="group flex items-center text-2xl">
		{#if innerWidth >= 500}
			<h1 class="font-bold transition-colors group-hover:text-orange-400">COP3530</h1>
			<div
				class="mx-2 h-7 w-0 border-l border-solid border-base-content transition-colors group-hover:border-[#E97451]"
			></div>
		{/if}
		<h1 class="font-extralight transition-colors group-hover:text-blue-500">
			Data Structures & Algorithms
		</h1>
	</a>
</header>
