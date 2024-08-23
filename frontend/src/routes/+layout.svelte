<script>
	import Header from '$lib/components/Header.svelte';
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';
	import { onNavigate } from '$app/navigation';
	import { Toaster } from '$lib/components/ui/sonner';
	import Auther from '$lib/components/Auther.svelte';
	import { injectSpeedInsights } from '@vercel/speed-insights/sveltekit';
	import { onMount } from 'svelte';

	injectSpeedInsights();

	onNavigate((navigation) => {
		// @ts-ignore
		if (!document.startViewTransition) return;

		return new Promise((resolve) => {
			// @ts-ignore
			document.startViewTransition(async () => {
				resolve();
				await navigation.complete;
			});
		});
	});

	let loaded = false;

	onMount(() => {
		setTimeout(() => {
			loaded = true;
		}, 1000);
	});
</script>

<svelte:head>
	<title>Blog</title>
</svelte:head>
{#if !loaded}
	<div class="fixed inset-0 flex items-center justify-center bg-white dark:bg-black">
		<div class="flex items-center gap-2">
			<p>Loading</p>
			<div class="h-4 w-4 animate-bounce rounded-full bg-black dark:bg-white"></div>
		</div>
	</div>
{:else}
	<ModeWatcher />
	<Toaster />
	<Auther />

	<Header />
	<slot></slot>
{/if}
