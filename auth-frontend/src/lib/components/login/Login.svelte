<script lang="ts">
	let email = $state('');
	let password = $state('');
</script>

<div class="flex w-full flex-col items-center px-12 py-8">
	<div class="w-full max-w-[700px]">
		<h1 class="text-2xl font-semibold">Log in to the DeGuzman Cloud</h1>
		<form class="mt-8 flex flex-col items-center gap-4">
			<label class="flex w-full flex-col gap-2">
				<h1 class="text-sm">Email</h1>
				<input
					type="email"
					bind:value={email}
					class="rounded-lg border-2 border-gray-700 bg-black p-2 focus:border-blue-300 focus:drop-shadow-[0_0_6px_var(--color-blue-300)] focus:outline-none"
				/>
			</label>

			<label class="flex w-full flex-col gap-2">
				<h1 class="text-sm">Password</h1>
				<input
					type="password"
					bind:value={password}
					class="rounded-lg border-2 border-gray-700 bg-black p-2 focus:border-blue-300 focus:drop-shadow-[0_0_6px_var(--color-blue-300)] focus:outline-none"
				/>
			</label>

			<button
				class="transition-duration-200 mt-4 w-full cursor-pointer rounded-lg border-2 border-solid border-white bg-black px-4 py-2 text-violet-300 drop-shadow-[0_0_6px_#fff] transition-colors hover:bg-stone-700"
			>
				<span class="btn-text text-sm text-white sm:text-base">Log In</span>
			</button>
			<!-- <button
			class="transition-duration-200 mt-4 ml-3 w-[150px] cursor-pointer rounded-full border-2 border-solid border-white bg-black px-4 py-2 text-violet-300 drop-shadow-[0_0_6px_#fff] transition-colors hover:bg-stone-700"
		>
			<span class="btn-text text-sm text-white sm:text-base">Log In</span>
		</button> -->
		</form>
	</div>
</div>
