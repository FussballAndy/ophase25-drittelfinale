<script lang="ts">
    let {isChanging, func} = $props();
</script>

<div class="select">
    <button class="sel-btn" disabled={isChanging} onclick={func}>
        {#if isChanging}
            Ausgewählt
        {:else}
            Wechseln
        {/if}
    </button>
</div>

<style>
    .select {
        display: flex;
        place-items: center;
    }
    .sel-btn {
        border: 1px solid var(--main-fg-color);
        padding: 0.3em;
        margin: 0.2em;
        background-color: transparent;
        color: var(--main-fg-color);
        border-radius: 0.5em;
        font-size: inherit;
        flex-grow: 1;
    }
    .sel-btn:hover:not(:disabled) {
        --main-fg-color: rgb(200 200 200 / 87%);
    }
    .sel-btn:active:not(:disabled) {
        background-color: var(--main-fg-color);
        color: var(--main-bg-color)
    }
    .sel-btn:disabled {
        --main-fg-color: rgb(150 150 150 / 87%);
    }
</style>