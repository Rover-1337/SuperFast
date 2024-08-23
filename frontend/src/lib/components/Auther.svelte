<script lang="ts">
	import { authed } from '$lib/store';
	import Cookies from 'js-cookie';

	if (Cookies.get('session')) {
		authed.set(true);
	} else {
		authed.set(false);
	}
</script>
