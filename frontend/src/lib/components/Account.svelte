<script lang="ts">
	import { User, UserCheck, UserX } from 'lucide-svelte';

	import { Button } from '$lib/components/ui/button';
	import * as AlertDialog from '$lib/components/ui/alert-dialog';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label/';
	import { toast } from 'svelte-sonner';

	import { API_URL } from '$lib/constant';
	import { authed } from '$lib/store';
	import * as HoverCard from '$lib/components/ui/hover-card';
	import Cookies from 'js-cookie';

	let username: string, password: string;

	async function addPost() {
		const res = await fetch(`${API_URL}/api/login`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			credentials: 'include',
			body: JSON.stringify({
				username: username,
				password: password
			})
		});

		if ((await res.json()).loggedin) {
			authed.set(true);
			toast.success('Login success');
		} else {
			toast.error('Login failed', {
				description: 'Please check your username and password'
			});
		}
	}
</script>

{#if $authed}
	<HoverCard.Root>
		<HoverCard.Trigger
			><Button class="flex items-center gap-2" disabled
				><User />
				Account</Button
			></HoverCard.Trigger
		>
		<HoverCard.Content
			><div class="flex items-center justify-center gap-4">
				<UserCheck />
				<div>Logged in</div>
			</div>
			<div class="mt-5">
				<Button
					variant="secondary"
					class="mx-auto flex items-center justify-center gap-8"
					on:click={() => Cookies.remove('session')}
					on:click={() => authed.set(false)}
					on:click={() => toast.success('Logout success')}
					><UserX />
					<div>Logout</div></Button
				>
			</div>
		</HoverCard.Content>
	</HoverCard.Root>{:else}
	<AlertDialog.Root>
		<AlertDialog.Trigger
			><Button class="flex items-center gap-2"
				><User />
				Account</Button
			>
		</AlertDialog.Trigger>
		<AlertDialog.Content>
			<AlertDialog.Header>
				<AlertDialog.Title>Login</AlertDialog.Title>
				<AlertDialog.Description>
					<div class="my-2 mt-8">
						<Label for="username">Username</Label>
						<Input
							bind:value={username}
							id="username"
							class="mb-4 mt-1"
							type="text"
							placeholder="Username"
						/>

						<Label for="password">Password</Label>
						<Input
							bind:value={password}
							id="password"
							class="mb-4 mt-1"
							type="password"
							placeholder="Password"
						/>
					</div>
				</AlertDialog.Description>
			</AlertDialog.Header>

			<AlertDialog.Footer>
				<AlertDialog.Cancel>Annuleer</AlertDialog.Cancel>
				<AlertDialog.Action on:click={() => addPost()}>Login</AlertDialog.Action>
			</AlertDialog.Footer>
		</AlertDialog.Content>
	</AlertDialog.Root>
{/if}
