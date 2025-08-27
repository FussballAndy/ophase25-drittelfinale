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
    <div class="inner-wrapper">
        <div class="spacer">
            <div class="icon"></div>
            <div class="info">
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
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: space-between;
        flex: 1;
    }
    .spacer {
        display: flex;
        flex-direction: row;
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
    .info {
        display: flex;
        flex-direction: column;
    }
    .name {
        font-size: 1.1em;
        font-weight: 500;
    }
</style>