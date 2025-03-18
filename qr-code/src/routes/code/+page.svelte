<script lang="ts">
	import type { PageProps } from './$types';
	import { createQrPngDataUrl } from '@svelte-put/qr';
	import { onMount } from 'svelte';

	let { data }: PageProps = $props();

	const qrConfig = {
		data: data.url,
		width: 256,
		height: 256,
		backgroundFill: '#fff'
	};

	let qrcodeData = $state('');
	onMount(async () => {
		qrcodeData = await createQrPngDataUrl(qrConfig);
	});
</script>

<div class="flex h-screen items-center">
	<div class="flex w-full flex-col items-center">
		<div class="flex w-full flex-col items-center gap-2 px-8">
			<h1 class="text-center text-8xl font-light">{data.title}</h1>
			<p class="mb-8 text-lg font-extralight">{data.description}</p>
			<img alt="qrcode" src={qrcodeData} width="256" height="256" />
		</div>
	</div>
</div>
