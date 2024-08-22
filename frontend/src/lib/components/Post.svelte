<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import Button from '$lib/components/ui/button/button.svelte';
	import { API_URL } from '$lib/constant';

	export let post: any;

	let deleted = false;

	async function deletepost(id: string) {
		const res = await fetch(`${API_URL}/api/deleteBlogPost/${id}`);

		if (res.ok) {
			deleted = true;
		}
	}
</script>

{#if !deleted}
	<Card.Root class="mx-auto my-4 w-1/3">
		<Card.Header>
			<Card.Title>
				<a href={`/blog/${post.id}`}>{post.title}</a>
			</Card.Title>
		</Card.Header>
		<Card.Content class="flex justify-between">
			<Button href={`/blog/${post.id}`}>Read more</Button>

			<Button on:click={() => deletepost(post.id)} variant="destructive">Delete</Button>
		</Card.Content>
	</Card.Root>
{/if}
