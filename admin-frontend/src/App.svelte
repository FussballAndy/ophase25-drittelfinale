<script lang="ts">
    import { onMount } from "svelte";
    import { wsClose, wsSend, wsStart } from "./lib/websocket";
    import { pageStore } from "./lib/stores";
    import RequestPage from "./pages/RequestPage.svelte";
    import IngamePage from "./pages/IngamePage.svelte";

    let token = $state("");
    let init = $state(false)

    function submit(e: Event) {
        e.preventDefault();
        wsStart(token);
        init = true;
    }

    onMount(() => {

        return wsClose
    })
</script>

{#if !init}
    <input type="text" bind:value={token}>
    <button onclick={submit}>Submit</button>
{:else}
    {#if $pageStore === 0}
        <span>Hallo</span>
    {:else if $pageStore === 1}
        <RequestPage />
    {:else if $pageStore === 2}
        <button onclick={() => wsSend(0)}>Bestätigen</button>
    {:else if $pageStore === 3}
        <IngamePage />
    {/if}
{/if}

<style>
  
</style>
