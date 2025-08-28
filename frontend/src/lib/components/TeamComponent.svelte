<script lang="ts">
    import Card from "./Card.svelte";
    import SelectButton from "./SelectButton.svelte";

    let {name, changer = undefined, isChanging = false, isCurrent = false, showPlayers = true} = $props();
    function changeTeam(e: Event) {
        e.preventDefault();
        changer && changer.set(name);
    }
</script>

<Card background={isCurrent}>
    <div class="inner-wrapper flex:row">
        <div class="spacer flex:row">
            <div class="icon"></div>
            <div class="flex:column">
                <span class="name">Team {name}</span>
                {#if showPlayers}
                    <span class="play">X / Y Spieler</span>
                {/if}
            </div>
        </div>
        {#if !isCurrent}
            <SelectButton isChanging={isChanging} func={changeTeam} />
        {/if}
    </div>
</Card>

<style>
    .inner-wrapper {
        align-items: center;
        justify-content: space-between;
        flex: 1;
    }
    .spacer {
        align-items: center;
        gap: 0.6em;
    }
    .icon {
        border: 1px solid black;
        border-radius: 100%;
        width: 2.5em;
        height: 2.5em;
        background-color: green;
        margin: 0.1em;
    }
    .name {
        font-size: 1.1em;
        font-weight: 500;
    }
</style>