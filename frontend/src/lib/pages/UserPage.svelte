<script lang="ts">
    import { writable } from "svelte/store";
    import TeamComponent from "../components/TeamComponent.svelte";
    import pageStore from "../pageStore";
    let change = writable(null);

    function advance(e: Event) {
        e.preventDefault();
        pageStore.set(2);
    }
</script>
<div class="surr">
    <span>Dein Team:</span>
    <TeamComponent name="1" isCurrent={true} changer={change} />
    {#if $change === null}
        <span>Weitere Teams:</span>
        {#each {length: 24}, idx}
            <TeamComponent name={idx} changer={change} />
        {/each}
    {:else}
        <span>Ausgewählt:</span>
        <TeamComponent name={$change} isChanging={true} changer={change} />
        <span>Wirklich wechseln?</span>
        <div class="change">
            <button class="change-btn green-btn">Ja</button>
            <button class="change-btn red-btn" onclick={() => change.set(null)}>Nein</button>
        </div>
    {/if}
    <button onclick={advance}>Advance</button>
</div>

<style>
    .surr {
        text-align: left;
        display: flex;
        flex-direction: column;
    }
    .change {
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 0.6em;
    }
    .change-btn {
        flex: 1;
        border: none;
        border-radius: 0.25em;
        font-size: inherit;
        padding: 0.4em;
        margin: 0.3em 0;
    }
</style>