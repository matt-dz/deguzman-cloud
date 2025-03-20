<div class="flex flex-col px-4">
	<a class="nav-container footer-underline w-fit" href="/">home</a>
	<a class="nav-container footer-underline w-fit" href="/services">services</a>
	<a class="nav-container footer-underline" href="https://auth.deguzman.cloud">login</a>
	<a class="nav-container footer-underline" target="_blank" href="https://about.deguzman.cloud"
		>about</a
	>
</div>

<style lang="postcss">
	a {
		@apply text-lg md:text-2xl font-light;
	}

	.footer-underline {
		@apply underline decoration-1 underline-offset-[3px];
	}

	.nav-container {
		@apply hover:bg-white w-fit hover:text-black p-1 md:p-2 transition-colors;
	}
</style>
