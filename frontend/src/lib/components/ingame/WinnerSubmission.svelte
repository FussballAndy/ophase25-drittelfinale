<script lang="ts">
    import { TIME_SLOTS } from "../../consts";

    const stationName = "Peter"
    let submitted = $state(0)

    function submitResult(e: Event) {
        e.preventDefault();
    }
</script>

<div class="flex:column card">
    <h2>Station {stationName}</h2>
    {#if submitted < TIME_SLOTS.length}
        <span>{TIME_SLOTS[submitted][0]} &ndash; {TIME_SLOTS[submitted][1]}</span>
        <div class="flex:column surr">
            <span style="text-align: left;">Gewinner:</span>
            <div class="flex:row wrapper">
                <div class="radio-wrapper">
                    <input type="radio" name="studtut" id="inputStuds">
                    <label for="inputStuds">Studenten</label>
                </div>
                <div class="radio-wrapper">
                    <input type="radio" name="studtut" id="inputTuts" checked>
                    <label for="inputTuts">Tutoren</label>
                </div>
            </div>
        </div>
        <button class="filled-btn white-btn" style="margin: 0.2em 4em" onclick={submitResult}>Absenden</button>
    {:else}
    <span>Alle Stationen fertig!</span>
    {/if}
</div>

<style>
    .surr {
        padding: 0 0.5em;
        margin: 0.6em 0;
        gap: 0.1em;
    }
    input[type="radio"] {
        appearance: none;
        width: 0;
    }
    label {
        cursor: pointer;
        padding: 1em;
        border: 1px solid var(--main-fg-color);
        border-radius: 0.3em;
        flex: 1;
    }
    input[type="radio"]:checked + label {
        background-color: var(--main-fg-color);
        color: var(--main-bg-color);
    }
    .wrapper {
        gap: 0.5em;
    }
    .radio-wrapper {
        display: flex;
        flex: 1;
        align-items: center;
        justify-content: center;
    }
</style>