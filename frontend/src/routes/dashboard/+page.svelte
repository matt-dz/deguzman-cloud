<script lang="ts">
	import type { PageData } from './$types';
	import type { Deployment as DeploymentType } from '$lib/types/deployment';
	import { IconTablePlus } from '@tabler/icons-svelte';
	import Deployment from '$lib/components/dgc/deployment/Deployment.svelte';
	let { data }: { data: PageData } = $props();

	const user = $derived(data.user);
	const deployments: DeploymentType[] = $derived(data.deployments);
	// const deployments: DeploymentType[] = [
	// 	{
	// 		config: {
	// 			build: {
	// 				arguments: [],
	// 				image: 'matthew10125/p4'
	// 			},
	// 			env: {},
	// 			mounts: [],
	// 			name: 'test-build3',
	// 			services: []
	// 		},
	// 		created_at: new Date('2024-12-22T05:23:39.754653Z'),
	// 		id: '02fe013a-7b2a-4ce1-afc2-f9150b590835',
	// 		name: 'test-build3',
	// 		status: 'RUNNING',
	// 		updated_at: new Date('2024-12-22T05:23:39.754653Z'),
	// 		user_id: '882a9486-9648-461b-bef7-60efa2c93bc1'
	// 	},
	// 	{
	// 		config: {
	// 			build: {
	// 				arguments: [],
	// 				image: 'matthew10125/p4'
	// 			},
	// 			env: {},
	// 			mounts: [],
	// 			name: 'test-build',
	// 			services: []
	// 		},
	// 		created_at: new Date('2024-12-22T04:25:00.913101Z'),
	// 		id: '2fe857bc-72c3-4a03-86ea-e4080824ca86',
	// 		name: 'test-build',
	// 		status: 'SCHEDULED',
	// 		updated_at: new Date('2024-12-22T04:25:00.913101Z'),
	// 		user_id: '882a9486-9648-461b-bef7-60efa2c93bc1'
	// 	},
	// 	{
	// 		config: {
	// 			build: {
	// 				arguments: [],
	// 				image: 'matthew10125/p4'
	// 			},
	// 			env: {},
	// 			mounts: [],
	// 			name: 'test-build',
	// 			services: []
	// 		},
	// 		created_at: new Date('2024-12-22T04:25:00.913101Z'),
	// 		id: '2fe857bc-72c3-4a03-86ea-e4080824ca84',
	// 		name: 'test-build',
	// 		status: 'IDLE',
	// 		updated_at: new Date('2024-12-22T04:25:00.913101Z'),
	// 		user_id: '882a9486-9648-461b-bef7-60efa2c93bc1'
	// 	},
	// 	{
	// 		config: {
	// 			build: {
	// 				arguments: [],
	// 				image: 'matthew10125/p4'
	// 			},
	// 			env: {},
	// 			mounts: [],
	// 			name: 'test-build',
	// 			services: []
	// 		},
	// 		created_at: new Date('2024-12-22T04:25:00.913101Z'),
	// 		id: '2fe857bc-72c3-4a03-86ea-e4080824ca83',
	// 		name: 'test-build',
	// 		status: 'ARCHIVED',
	// 		updated_at: new Date('2024-12-22T04:25:00.913101Z'),
	// 		user_id: '882a9486-9648-461b-bef7-60efa2c93bc1'
	// 	},
	// 	{
	// 		config: {
	// 			build: {
	// 				arguments: [],
	// 				image: 'matthew10125/p4'
	// 			},
	// 			env: {},
	// 			mounts: [],
	// 			name: 'test-build',
	// 			services: []
	// 		},
	// 		created_at: new Date('2024-12-22T04:25:00.913101Z'),
	// 		id: '2fe857bc-72c3-4a03-86ea-e4080824ca82',
	// 		name: 'test-build',
	// 		status: 'ARCHIVED',
	// 		updated_at: new Date('2024-12-22T04:25:00.913101Z'),
	// 		user_id: '882a9486-9648-461b-bef7-60efa2c93bc1'
	// 	},
	// 	{
	// 		config: {
	// 			build: {
	// 				arguments: [],
	// 				image: 'matthew10125/p4'
	// 			},
	// 			env: {},
	// 			mounts: [],
	// 			name: 'test-build',
	// 			services: []
	// 		},
	// 		created_at: new Date('2024-12-22T04:25:00.913101Z'),
	// 		id: '2fe857bc-72c3-4a03-86ea-e4080824ca81',
	// 		name: 'test-build',
	// 		status: 'ARCHIVED',
	// 		updated_at: new Date('2024-12-22T04:25:00.913101Z'),
	// 		user_id: '882a9486-9648-461b-bef7-60efa2c93bc1'
	// 	}
	// ];
</script>

<section class="mt-16">
	<div class="w-full max-w-[1080px]">
		<h1 class="text-4xl font-medium">
			Welcome back, <br /><span class="text-blue-400 capitalize"
				>{user.first_name}<span>.</span></span
			>
		</h1>
	</div>
</section>

<section class="mt-32">
	<div class="w-full max-w-[1080px]">
		<div class="flex items-center">
			<h1 class="flex-grow text-2xl font-medium">Deployments</h1>
			<a
				href="/deployments/new"
				class="px-3 py-2 border border-solid border-gray-200 shadow-md rounded-full font-light hover:bg-gray-100 transition-colors duration-150 flex gap-2 items-center"
			>
				<IconTablePlus stroke={1} width={8} color="green" /> Create
			</a>
		</div>
		<div class="flex flex-col items-center w-full mt-4">
			<div class="grid grid-cols-3 gap-4">
				{#each deployments as deployment (deployment.id)}
					<Deployment {deployment} />
				{/each}
			</div>
		</div>
	</div>
</section>

<style lang="postcss">
	section {
		@apply w-full flex flex-col items-center;
	}
</style>
