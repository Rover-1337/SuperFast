<script lang="ts">
	import { API_URL } from '$lib/constant';
	import * as Card from '$lib/components/ui/card';
	import Button from '$lib/components/ui/button/button.svelte';
	import { Skeleton } from '$lib/components/ui/skeleton';

	async function fetchPosts() {
		const res = await fetch(`${API_URL}/api/blogPosts`);
		const data = await res.json();
		return data;
	}
</script>

{#await fetchPosts()}
	{#each Array(1, 2) as _}
		<div>
			<Card.Root class="w-1/3 mx-auto my-4">
				<Card.Header>
					<Card.Title>
						<Skeleton class="w-1/2" />
					</Card.Title>
				</Card.Header>
				<Card.Content>
					<Skeleton class="w-1/2" />
				</Card.Content>
			</Card.Root>
		</div>
	{/each}
{:then posts}
	{#if posts.data == null}
		<p>No posts found</p>
	{:else}{#each posts.data as post}
			<div>
				<Card.Root class="w-1/3 mx-auto my-4">
					<Card.Header>
						<Card.Title>
							<a href={`/blog/${post.id}`}>{post.title}</a>
						</Card.Title>
					</Card.Header>
					<Card.Content>
						<Button href={`/blog/${post.id}`}>Read more</Button>
					</Card.Content>
				</Card.Root>
			</div>
		{/each}
	{/if}
{:catch error}
	<p>{error.message}</p>
{/await}
