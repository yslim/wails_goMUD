<script>
	import { page } from '$app/stores';
    import {onMount} from "svelte";

    onMount(() => {
        console.log ('pathname = ', $page.url.pathname);
    });
</script>

<header>
</header>

<style>
</style>
