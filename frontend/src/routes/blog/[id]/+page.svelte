<script>
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { API_URL } from '$lib/constant';

	export let data;

	async function load() {
		const res = await fetch(`${API_URL}/api/blogPosts/${data.params.id}`);
		const json = await res.json();
		return json;
	}
</script>

{#await load() then post}
	<Card.Root class="mx-auto mt-10 w-3/4">
		<Card.Header>
			<Card.Title>{post.data.title}</Card.Title>
			<Card.Description>{post.data.description}</Card.Description>
		</Card.Header>
		<Card.Content>
			<p>{post.data.content}</p>

			<Button class="my-4" href="/">Back to blog</Button>
		</Card.Content>
		<Card.Footer>
			<p>{post.data.created_at.split('.')[0]}</p>
		</Card.Footer>
	</Card.Root>
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}
