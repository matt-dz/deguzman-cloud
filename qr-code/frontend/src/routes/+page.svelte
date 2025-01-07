<script lang="ts">
	import type { QRCode } from '$lib/types';
	import { goto } from '$app/navigation';

	let url = $state('');
	let title = $state('');
	let description = $state('');

	async function submitURL(event: Event) {
		event.preventDefault();

		try {
			const resp = await fetch(`${import.meta.env.VITE_BASE_URL}/api/qr`, {
				method: 'POST',
				body: JSON.stringify({ url }),
				headers: {
					'Content-Type': 'application/json'
				}
			});

			if (resp.status != 200) {
				alert('Error creating QR code');
				console.error(resp.statusText);
			}

			const qrCode: QRCode = await resp.json();
			await goto(`/view/${qrCode.id}?title=${title}&description=${description}`);
		} catch (error) {
			alert('Error creating QR code');
			console.error(error);
		}
	}
</script>

<div class="flex h-[100vh] items-center">
	<div class="flex w-full flex-col items-center">
		<form
			onsubmit={async (event) => submitURL(event)}
			class="flex w-full max-w-[700px] flex-col items-center gap-2"
		>
			<h1 class="mb-3 text-center text-5xl font-light text-white">QR Code Generator</h1>
			<input bind:value={url} type="text" placeholder="URL*" />
			<input bind:value={title} type="text" placeholder="Title" />
			<input bind:value={description} type="text" placeholder="Description" />
			<button
				class="border-base-content mt-4 rounded-xl border border-solid px-3 py-1 text-lg font-extralight"
				>Create</button
			>
		</form>
	</div>
</div>

<style lang="postcss">
	input {
		@apply input input-bordered h-10 w-full max-w-sm !bg-transparent;
	}
</style>
