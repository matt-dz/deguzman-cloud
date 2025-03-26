<script lang="ts">
	import Link from '$lib/components/link/Link.svelte';
	import Discussion from '$lib/components/discussion/Discussion.svelte';
</script>

<div class="flex items-center">
	<div class="flex w-full flex-col items-center">
		<div class="flex w-full max-w-[610px] flex-col px-4">
			<div class="group collapse collapse-arrow mt-24 bg-base-200">
				<input type="checkbox" name="about" />
				<div class="accordion-header collapse-title group-hover:text-blue-400">About Your TA</div>
				<div class="collapse-content">
					<p>
						Hello! My name is Matthew DeGuzman, I am a fourth year computer science major and this
						is my fourth semester being a TA for this class.
					</p>
					<br />
					<p>
						I'm the <Link href="https://x.swamphacks.com" borderColor="red"
							>{#snippet text()}SwampHacks{/snippet}</Link
						> Technical Director, <Link href="https://uf.colorstack.org" borderColor="orange"
							>{#snippet text()}ColorStack{/snippet}</Link
						> Technical Director, and software developer for the <Link href="https://ufcsu.org"
							>{#snippet text()}Computing Student Union{/snippet}</Link
						>. If you have any questions about these organizations or want to get involved, feel
						free to ask me!
					</p>
					<br />
					<p>
						I have 3 previous internships at Microsoft in Seattle and I will be joining Datadog as a
						software engineer in New York after I graduate. In my free time, I enjoy <Link
							href="https://food.matthewdeguzman.info"
							borderColor="orange">{#snippet text()}cooking{/snippet}</Link
						>, travelling, and playing pickleball.
					</p>
					<br />
					<p>
						Learn more about me on my <Link href="https://matthewdeguzman.info" borderColor="green"
							>{#snippet text()}personal website{/snippet}</Link
						> or connect with me on <Link
							href="https://linkedin.com/in/matthew-deguzman"
							borderColor="purple">{#snippet text()}LinkedIn{/snippet}</Link
						> :)
					</p>
				</div>
			</div>

			<div class="divider-header divider divider-primary" id="general">General Information</div>
			<div class="flex w-full flex-col gap-2">
				<div class="font-extralight">
					<h1 class="font-bold">Section 10696</h1>
					<div class="flex flex-wrap items-center">
						<p>NRN 1037</p>
						<div class="my-divider"></div>
						<p>Tuesday</p>
						<div class="my-divider"></div>
						<p>11:45 am - 12:35 pm</p>
					</div>
				</div>
				<div class="font-extralight">
					<h1 class="font-bold">Office Hours</h1>
					<div class="flex flex-wrap items-center">
						<p>MALA 5200 (Malachowsky Hall)</p>
						<div class="my-divider"></div>
						<p>Thursday</p>
						<div class="my-divider"></div>
						<p>8:30 am - 9:20 am & 10:40 am - 11:30 am</p>
					</div>
				</div>

				<div>
					<h1 class="font-bold">Email</h1>
					<div class="w-fit">
						<Link href="mailto:matthewdeguzman@ufl.edu" borderColor="green"
							>{#snippet text()}<p class="font-extralight">
									matthewdeguzman@ufl.edu
								</p>{/snippet}</Link
						>
					</div>
				</div>
			</div>

			<div class="divider-header divider divider-secondary mt-12" id="discussions">Discussions</div>
			<div class="mb-4 flex flex-col gap-4">
				<Discussion
					name="c++"
					title="1. C++ (01/14)"
					pptLink="https://ufl.instructure.com/courses/525732/files/folder/Slides%20for%20Inperson%20Lectures/Discussion%20Presentations?preview=93479607"
				>
					{#snippet description()}
						<p>
							Welcome to the first week of discussion! We're doing a quick overview of C++. Get
							ready to brush up on pointers, data types, classes, and more!
						</p>
					{/snippet}
				</Discussion>
				<Discussion
					name="algorithmic-analysis"
					title="2. Algorithmic Analysis and Pseudocode (01/21)"
					pptLink="https://docs.google.com/presentation/d/1SKauPvsXgH6VXktNGxwmlc8C0m5byvu_IYKe1cP9v9Y/edit?usp=sharing"
				>
					{#snippet description()}
						<p>
							This week, we're diving into the world of algorithmic analysis and pseudocode. Get
							ready to learn how to analyze the efficiency of algorithms and write pseudocode to
							plan out your code.
						</p>
					{/snippet}
				</Discussion>
				<Discussion
					name="algorithmic-analysis"
					title="3. Lists, Stacks, and Queues (01/28)"
					pptLink="https://docs.google.com/presentation/d/1BJCV0xbhYQweHNKt5Cx7r-8tIpKw9Ndt/edit?usp=sharing&ouid=101706383218741169050&rtpof=true&sd=true"
				>
					{#snippet description()}
						<p>
							This week, we review stacks and queues as well as common algorithmic strategies such
							as the slow pointer, fast pointer technique on Linked Lists.
						</p>
					{/snippet}
				</Discussion>
				<Discussion
					name="algorithmic-analysis"
					title="4. Trees and Traversals (02/04)"
					pptLink="https://docs.google.com/presentation/d/1SGjaeApZLrveO4CkKODq5qO72WVHeOAvq4IB9hM2E-I/edit?usp=sharing"
				>
					{#snippet description()}
						<p>
							This week, we review binary trees and the different tree traversals. We'll also cover
							a special case called binary search trees.
						</p>
					{/snippet}
				</Discussion>
				<Discussion
					name="algorithmic-analysis"
					title="5. Balanced Trees (02/11)"
					pptLink="https://docs.google.com/presentation/d/10UCZmsEmkvm1EDEbkpLQU3r-4Sm7zPDUzkLD2kaJcZg/"
				>
					{#snippet description()}
						<p>
							This week, we go over balanced trees and their advantages over binary search trees.
							We'll mainly focus on AVL trees and touch on red-black, splay, B, and B+ trees.
						</p>
					{/snippet}
				</Discussion>
				<Discussion
					name="algorithmic-analysis"
					title="6. Mock Interview (02/18)"
					pptLink="https://docs.google.com/presentation/d/1g95fl4lzjhVttD3vFRIj9rGXw_teVy0nj7af_hHynXM/edit?usp=sharing"
				>
					{#snippet description()}
						<p>
							This week, we dive into a mock interview to prepare you for the job market! We'll go
							over resumes, behavioral questions, and technical questions.
						</p>
					{/snippet}
				</Discussion>
				<Discussion
					name="algorithmic-analysis"
					title="7. Exam 1 Review (02/25)"
					pptLink="https://docs.google.com/presentation/d/1DXYDmENGEAaiuQMrsta6MlvZ4v3MFMH2q3dk8q7CwV4/edit"
				>
					{#snippet description()}
						<p>Exam 1 Review!!! Review trees, BSTs, and heaps in preparation for the midterm.</p>
					{/snippet}
				</Discussion>
				<Discussion
					name="sorting"
					title="8. Sorting (03/11)"
					pptLink="https://docs.google.com/presentation/d/1ONDt_dJn_ISCR3iWgABeoauVqaY9W80JxOgAckExL5o/edit?usp=sharing"
				>
					{#snippet description()}
						<p>
							This week we discuss sorting algorithms! We'll go over the basic quadratic sorts, then
							the more advanced linearithmic sorts, and some niche sorts such as shell sort.
						</p>
					{/snippet}
				</Discussion>
				<Discussion
					name="sets-maps-hash-tables"
					title="8. Sets, Maps, and Hash Tables (03/25)"
					pptLink="https://docs.google.com/presentation/d/1pXlepCty9XjII8XISoeBy_vM0A4USPznhSYxAY0lA0M/edit?usp=sharing"
				>
					{#snippet description()}
						<p>
							We go over sets, maps, and their breathtaking implementations using hash tables. We'll
							also discuss the different strategies of hashing and their advantages/disadvantages.
						</p>
					{/snippet}
				</Discussion>
			</div>
		</div>
	</div>
</div>

<style lang="postcss">
	.accordion-header {
		@apply text-xl font-extralight;
	}

	.divider-header {
		@apply text-xl font-extralight;
	}

	.my-divider {
		@apply mx-2 h-[1.1rem] w-0 border-l border-base-content;
	}
</style>
