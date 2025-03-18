<script lang="ts">
	let { status }: { status: string } = $props();
	status = status.toUpperCase();
</script>

<div
	class="w-5 h-5 border border-solid border-gray-500 rounded-full"
	class:bg-green-400={status === 'RUNNING'}
	class:bg-white={status === 'IDLE'}
	class:bg-blue-400={status === 'SCHEDULED'}
	class:bg-black={status === 'ARCHIVED'}
></div>
