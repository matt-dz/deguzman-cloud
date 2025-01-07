<script lang="ts">
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	const id = $derived(data.id);
	const title = $derived(data.title);
	const description = $derived(data.description);
</script>

<div class="flex h-screen items-center">
	<div class="flex w-full flex-col items-center">
		<div class="flex w-full flex-col items-center gap-2 px-8">
			<h1 class="text-center text-8xl font-light">{title}</h1>
			<p class="mb-8 text-lg font-extralight">{description}</p>
			<img alt="qrcode" src={`${import.meta.env.VITE_BASE_URL}/api/qr/${id}`} />
		</div>
	</div>
</div>
