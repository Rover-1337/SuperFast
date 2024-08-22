<script lang="ts">
	import Post from '$lib/components/Post.svelte';
	import { API_URL } from '$lib/constant';

	async function fetchPosts() {
		const res = await fetch(`${API_URL}/api/blogPosts`);
		const data = await res.json();
		return data;
	}
</script>

{#await fetchPosts() then posts}
	{#if posts.data == null}
		<p>No posts found</p>
	{:else}{#each posts.data as post}
			<Post {post} />
		{/each}
	{/if}
{:catch error}
	<p>{error.message}</p>
{/await}
