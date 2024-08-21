<script>
	import Header from '$lib/components/Header.svelte';
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';
</script>

<svelte:head>
	<title>Blog</title>
</svelte:head>

<ModeWatcher />

<Header />
<slot></slot>
