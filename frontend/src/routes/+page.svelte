<svelte:head>
	<title>deguzman | home </title>
</svelte:head>

<div class="flex flex-col">
	<div class="px-4 pt-2 sm:py-12 w-full col-start-2 flex flex-col">
		<div class="flex flex-col md:gap-1">
			<h1 class="text-4xl md:text-7xl font-[550]">deguzman.cloud</h1>
			<p class="text-lg md:text-3xl font-light">/deguzman dät kloud/</p>
			<p class="text-lg md:text-3xl font-light italic">noun</p>
			<div class="text-lg md:text-3xl pl-2 md:pl-4 definition">
				<p class="col-start-1">1.</p>
				<p class="col-start-2">
					The <span class="italic text-[#6AA9FA]">ultimate</span> cloud service provider.
				</p>
				<p class="col-start-2 definition-sentence">
					"I use <span class="definition-underline md:lg-definition-underline">deguzman.cloud</span>
					to run my unicorn b2b saas."
				</p>
			</div>

			<div class="text-lg md:text-3xl pl-2 md:pl-4 definition">
				<p class="col-start-1">2.</p>
				<p class="col-start-2">
					A <span class="italic text-[#6AFA9A]">vast</span> suite of tools and information hub.
				</p>
				<p class="col-start-2 definition-sentence">
					"I used <span class="definition-underline md:lg-definition-underline">deguzman.cloud</span
					> to store all my files!"
				</p>
			</div>
		</div>
	</div>
</div>

<style lang="postcss">
	.definition {
		@apply grid grid-cols-[2.5rem_auto] font-light;
	}

	.definition-sentence {
		@apply text-[#ADADAD];
	}

	.definition-underline {
		@apply underline decoration-1 underline-offset-[3px];
	}

	.lg-definition-underline {
		@apply underline decoration-2 underline-offset-[6px];
	}
</style>
