<script lang="ts">
	import Login from '$lib/components/login/Login.svelte';
</script>

<div class="flex h-dvh">
	<div class="flex w-full flex-col">
		<div class="grid h-full grid-cols-1 md:grid-cols-2">
			<div class="hidden md:visible md:flex md:items-center"></div>
			<div class="flex w-full items-center border-l-2 border-solid border-gray-700">
				<Login />
			</div>
		</div>
	</div>
</div>
