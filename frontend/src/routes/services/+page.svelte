<script lang="ts">
	import Service from '$lib/components/dgc/service/Service.svelte';
</script>

<svelte:head>
	<title>deguzman | services</title>
</svelte:head>

<div class="flex flex-col">
	<div class="px-4 pt-2 sm:py-12 w-full col-start-2 flex flex-col">
		<div class="flex flex-col gap-1">
			<h1 class="text-4xl md:text-7xl font-[550]">services</h1>
			<div class="grid lg:grid-cols-2 gap-6 lg:gap-10 mt-12">
				<Service title={'cloud'} url={'deguzman.cloud'}>
					{#snippet description()}
						<p class="text-lg md:text-xl font-light">
							Use the <span class="svc-underline">deguzman cloud</span> to level up your projects for
							web hosting or for running compute jobs.
						</p>
					{/snippet}
				</Service>

				<Service title={'qr code'} url={'qr.deguzman.cloud'}>
					{#snippet description()}
						<p class="text-lg md:text-xl font-light">
							Easily generate QR codes for your all your events at lightning speed.
						</p>
					{/snippet}
				</Service>

				<Service title={'storage'} url={'dss.deguzman.cloud'}>
					{#snippet description()}
						<p class="text-lg md:text-xl font-light">
							Store your static files and witness blazingly fast speeds with deguzman simple storage
							(dss).
						</p>
					{/snippet}
				</Service>

				<Service title={'blog'} url={'blog.deguzman.cloud'}>
					{#snippet description()}
						<p class="text-lg md:text-xl font-light">
							Read up on the amazing thoughts, stories, and ideas from the creators of the <span
								class="svc-underline">deguzman cloud</span
							>.
						</p>
					{/snippet}
				</Service>
			</div>
		</div>
	</div>
</div>

<style lang="postcss">
	.svc-underline {
		@apply underline decoration-1 underline-offset-[3px];
	}
</style>
