<script lang="ts">
	import { API_URL } from '$lib/constant';
	import { ScrollArea } from '$lib/components/ui/scroll-area';
	import { Separator } from '$lib/components/ui/separator';
	import { Skeleton } from './ui/skeleton';
	async function getAllUsers() {
		const res = await fetch(`${API_URL}/api/auth/dashboard`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			},
			credentials: 'include'
		});

		return res.json();
	}
</script>

{#await getAllUsers()}
	<ScrollArea class="h-72 w-[95%] rounded-md border">
		<div class="p-4">
			<h4 class="mb-4 text-sm font-medium leading-none">Tags</h4>
			{#each Array(20) as i}
				<div class="text-sm">
					<div>
						<Skeleton class="h-4 w-20" />
					</div>
				</div>
				<Separator class="my-2" />
			{/each}
		</div>
	</ScrollArea>
{:then users}
	<ScrollArea class="h-72 w-[95%] rounded-md border">
		<div class="p-4">
			<h4 class="mb-4 text-sm font-medium leading-none">Tags</h4>
			{#each users.data as user}
				<div class="text-sm">
					<div>{user.username}</div>
				</div>
				<Separator class="my-2" />
			{/each}
		</div>
	</ScrollArea>
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}
