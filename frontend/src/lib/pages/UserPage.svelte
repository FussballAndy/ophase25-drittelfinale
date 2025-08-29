<script lang="ts">
    import { writable } from "svelte/store";
    import TeamComponent from "../components/TeamComponent.svelte";
    let change = writable(null);

</script>
<div class="surr flex:column">
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
        <div class="change flex:row">
            <button class="filled-btn green-btn">Ja</button>
            <button class="filled-btn red-btn" onclick={() => change.set(null)}>Nein</button>
        </div>
    {/if}
</div>

<style>
    .surr {
        text-align: left;
    }
    .change {
        justify-content: center;
        gap: 0.6em;
    }
</style>