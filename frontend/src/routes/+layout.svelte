<script>
	import Header from '$lib/components/Header.svelte';
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';
</script>

<ModeWatcher />

<Header />
<slot></slot>
