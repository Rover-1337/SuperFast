<script lang="ts">
	import { Moon, Sun } from 'lucide-svelte';
	import { toggleMode } from 'mode-watcher';
	import { Button } from '$lib/components/ui/button';
	import AddPost from './AddPost.svelte';
	import Account from './Account.svelte';
	import { API_URL } from '$lib/constant';
	import { Skeleton } from './ui/skeleton';
	import { authed } from '$lib/store';
	import * as Dialog from '$lib/components/ui/dialog';
	import Dash from './Dash.svelte';

	async function checkifLoggedIn() {
		const res = await fetch(`${API_URL}/api/auth/status`, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			},
			credentials: 'include'
		});

		if ((await res.json()).loggedin) {
			return true;
		} else {
			return false;
		}
	}
</script>

<div class="flex items-center justify-between border-b p-4">
	<div class="flex items-center gap-2">
		<h1>
			<a href="/">Rover's Hub</a>
		</h1>
		<Dialog.Root>
			<Dialog.Trigger><Button variant="ghost">Blog</Button></Dialog.Trigger>
			<Dialog.Content>
				<Dash />
			</Dialog.Content>
		</Dialog.Root>
	</div>
	<div class="flex items-center gap-4">
		<AddPost />
		<Account />
		{#key $authed}
			{#await checkifLoggedIn()}
				<Button disabled variant="outline"><Skeleton class="h-4 w-[60px]" /></Button>
			{:then loggedIn}
				{#if loggedIn}
					<Button disabled variant="outline">Loggin in</Button>
				{:else}
					<Button disabled variant="outline">Logged out</Button>
				{/if}
			{:catch error}
				<p style="color: red">{error.message}</p>
			{/await}
		{/key}
		<Button on:click={toggleMode} variant="outline" size="icon">
			<Sun
				class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
			/>
			<Moon
				class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
			/>
			<span class="sr-only">Toggle theme</span>
		</Button>
	</div>
</div>
