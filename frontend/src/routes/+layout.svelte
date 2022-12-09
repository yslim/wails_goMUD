<script lang="ts">
    import "../app.css";
    import Header from "./Header.svelte";
    // import { page } from '$app/stores';
    // import {onMount} from "svelte";
    //
    // onMount(() => {
    //     page.subscribe((page) => {
    //         console.log ('page.url.pathname = ', page.url.pathname);
    //     })
    // });
</script>

<div class="app">
    <Header/>

    <div class="body-wrapper">
        <slot/>
    </div>
</div>

<style>
</style>
