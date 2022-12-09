<script lang="ts">
    import {onMount} from "svelte";
    import 'xterm/css/xterm.css';

    let terminalElement: HTMLElement;
    let isResizing = false;
    let terminal: any;

    async function initTerminal() {
        // import package dynamically
        const xterm = await import("xterm");
        const fit = await import("xterm-addon-fit");
        const webLinksAddon = await import("xterm-addon-web-links");

        // import classes
        const { Terminal } = xterm;
        const { FitAddon } = fit;
        const { WebLinksAddon } = webLinksAddon;

        // init terminal & addons
        terminal = new Terminal({allowProposedApi: true});
        const termFit = new FitAddon();
        terminal.loadAddon(new WebLinksAddon());
        terminal.loadAddon(termFit);

        // open terminal in DOM
        terminal.open(terminalElement);
        termFit.fit();

        // resize event handler
        window.addEventListener ("resize", (event: Event) => {
            if (isResizing)
                return;

            isResizing = true;
            setTimeout (() => {
                isResizing = false;
                termFit.fit();
            }, 500);
        });

        // paste (CMD-V) data
        terminal.onData((e) => {
            terminal.write(e);
        })

        for (let i=0; i<100; i++) {
            terminal.write('Hello from \x1B[1;3;31mxterm.js\x1B[0m $ ')
            terminal.write("\n\r");
        }
    }

    onMount(() => {
        initTerminal();
    });
</script>

<svelte:head>
    <title>Home</title>
</svelte:head>

<div class="output-layer" bind:this={terminalElement}></div>
<div class="w-full h-2 bg-gray-400"></div>
<div class="input-layer">
    <input type="text" class="input input-bordered w-full bg-black h-20"/>
</div>

<style>
    .output-layer {
        @apply w-full;
        @apply h-full;
        @apply bg-black;
        @apply text-white;
        padding: 1rem 0 1rem 1rem;
    }

    .input-layer {
        @apply flex;
        @apply flex-col;
        @apply p-1.5;
        @apply w-full;
        @apply h-fit;
        background-color: #333;
        color: white;
        @apply overflow-auto;
        @apply rounded-none;
    }
</style>
