<script lang="ts">
	import { API_URL } from '$lib/constant';
	import * as Card from '$lib/components/ui/card';
	import Button from '$lib/components/ui/button/button.svelte';

	let rerender = 0;

	async function fetchPosts() {
		const res = await fetch(`${API_URL}/api/blogPosts`);
		const data = await res.json();
		return data;
	}

	async function deletepost(id: string) {
		const res = await fetch(`${API_URL}/api/deleteBlogPost/${id}`);

		if (res.ok) {
			rerender++;
		}
	}
</script>

{#key rerender}
	{#await fetchPosts() then posts}
		{#if posts.data == null}
			<p>No posts found</p>
		{:else}{#each posts.data as post}
				<div>
					<Card.Root class="mx-auto my-4 w-1/3">
						<Card.Header>
							<Card.Title>
								<a href={`/blog/${post.id}`}>{post.title}</a>
							</Card.Title>
						</Card.Header>
						<Card.Content class="flex justify-between">
							<Button href={`/blog/${post.id}`}>Read more</Button>

							<Button
								on:click={() => deletepost(post.id)}
								href={`/api/deleteBlogPost/${post.id}`}
								variant="destructive">Delete</Button
							>
						</Card.Content>
					</Card.Root>
				</div>
			{/each}
		{/if}
	{:catch error}
		<p>{error.message}</p>
	{/await}
{/key}
