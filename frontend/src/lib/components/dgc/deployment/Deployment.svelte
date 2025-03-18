<script lang="ts">
	import { type Deployment } from '$lib/types/deployment';
	import Status from './Status.svelte';
	let { deployment }: { deployment: Deployment } = $props();
	const created_at = new Date(deployment.created_at);
	const updated_at = new Date(deployment.updated_at);
</script>

<a href={`/dashboard/deployment/${deployment.id}`}>
	<div
		class="w-[350px] h-[180px] hover:bg-gray-100 border border-gray-200 border-solid p-3 rounded-xl shadow-md transition-color duration-150"
	>
		<div class="flex flex-col h-full">
			<div class="flex flex-grow items-start">
				<div class="flex items-center w-full">
					<h1 class="text-xl flex-grow truncate font-light">{deployment.config.name}</h1>
					<Status status={deployment.status} />
				</div>
			</div>

			<div class="flex">
				<div class="flex-grow">
					<h1 class="text-sm">Created at</h1>
					<p class="font-light text-sm">
						{created_at.getMonth()}-{created_at.getDate()}-{created_at.getFullYear()}
					</p>
				</div>

				<div>
					<h1 class="text-sm">Updated at</h1>
					<p class="font-light text-sm">
						{updated_at.getMonth()}-{updated_at.getDate()}-{updated_at.getFullYear()}
					</p>
				</div>
			</div>
		</div>
	</div>
</a>

<style lang="postcss">
	.container-shadow {
		transition: box-shadow 0.2s ease-out;
	}

	.container-shadow:hover {
		box-shadow: 0 0 8px #888;
	}
</style>
