<script lang="ts">
	import { CirclePlus } from 'lucide-svelte';

	import { Button } from '$lib/components/ui/button';
	import * as AlertDialog from '$lib/components/ui/alert-dialog';
	import { Input } from '$lib/components/ui/input';
	import { Textarea } from '$lib/components/ui/textarea';
	import { API_URL } from '$lib/constant';
	import { goto } from '$app/navigation';

	let titel: string, description: string, content: string;

	async function addPost() {
		const res = await fetch(`${API_URL}/api/blogPosts`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				title: titel,
				description: description,
				content: content
			})
		});

		if (res.ok) {
			goto(`/api/blogPosts/${(await res.json()).id}`);
		}
	}
</script>

<AlertDialog.Root>
	<AlertDialog.Trigger
		><Button class="flex items-center gap-2"
			><CirclePlus />
			Add post</Button
		>
	</AlertDialog.Trigger>
	<AlertDialog.Content>
		<AlertDialog.Header>
			<AlertDialog.Title>Voeg een post toe</AlertDialog.Title>
			<AlertDialog.Description>
				<p>Probeer het elke dag vol te houden</p>
				<div class="my-2 mt-8">
					<Input bind:value={titel} class="mb-4" type="text" placeholder="Titel" />
					<Input bind:value={description} class="mb-4" type="text" placeholder="Description" />
					<Textarea bind:value={content} placeholder="Content hier" />
				</div>
			</AlertDialog.Description>
		</AlertDialog.Header>

		<AlertDialog.Footer>
			<AlertDialog.Cancel>Annuleer</AlertDialog.Cancel>
			<AlertDialog.Action on:click={() => addPost()}>Voeg toe</AlertDialog.Action>
		</AlertDialog.Footer>
	</AlertDialog.Content>
</AlertDialog.Root>
