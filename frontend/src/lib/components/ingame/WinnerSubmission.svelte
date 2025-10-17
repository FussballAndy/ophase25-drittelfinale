<script lang="ts">
    import { get as getStore } from "svelte/store";
    import { API_BASE_URL, TIME_SLOTS } from "../../consts";
    import type { WinnerSub } from "../../types";
    import { groupsStore, stationsStore, userStationStore, userStore } from "../../stores";

    const groups = Array(3).fill(0).map((_, idx) => 
        $groupsStore
            .map((g, gIdx) => [g.stations[idx], gIdx])
            .filter((g, _) => g[0] == $userStationStore)
            .map((g) => g[1]).at(0)
    );
    console.log(groups);
    let submitted = $state(0); // CHECK
    let buttonsDisabled = $state(false);
    let tutorChecked = $state(1);

    $effect(() => {
        console.log($stationsStore[$userStationStore])
    })

    async function sendResult() {
        const data: WinnerSub = {
            token: getStore(userStore)!,
            iteration: submitted,
            score: tutorChecked,
        };
        let result = await fetch(API_BASE_URL + "/api/winner", {
            method: "POST",
            body: JSON.stringify(data)
        }).then(res => res.json()).catch(res => {status: false});
        if(result.status) {
            submitted++;
        } else {
            alert("Something went wrong sending the result!");
        }
        buttonsDisabled = false;
    }

    function submitResult(e: Event) {
        e.preventDefault();
        buttonsDisabled = true;
        sendResult()
    }
</script>

<div class="flex:column card">
    <h2>Station {$stationsStore[$userStationStore].name}</h2>
    {#if submitted < TIME_SLOTS.length}
        <span>{TIME_SLOTS[submitted][0]} &ndash; {TIME_SLOTS[submitted][1]}</span>
        <div class="flex:column surr">
            <span style="text-align: left;">Gewinner:</span>
            <div class="flex:row wrapper" class:blurred={buttonsDisabled}>
                <div class="radio-wrapper">
                    <input type="radio" name="studtut" id="inputStuds" bind:group={tutorChecked} disabled={buttonsDisabled} value={0}>
                    <label for="inputStuds">Studenten</label>
                </div>
                <div class="radio-wrapper">
                    <input type="radio" name="studtut" id="inputTuts" checked bind:group={tutorChecked} disabled={buttonsDisabled} value={1}>
                    <label for="inputTuts">Tutoren</label>
                </div>
            </div>
        </div>
        <button class="filled-btn white-btn" style="margin: 0.2em 4em" onclick={submitResult} disabled={buttonsDisabled}>Absenden</button>
    {:else if submitted == TIME_SLOTS.length}
        <span>Intrawettbewerb</span>
        <div class="flex:column surr">
            <span style="text-align: left;">Gewinner:</span>
            <div class="flex:column wrapper" class:blurred={buttonsDisabled}>
                {#each groups as group, idx (idx)}
                    <div class="radio-wrapper">
                        <input type="radio" name="intra" id={"inputIntra" + idx} bind:group={tutorChecked} disabled={buttonsDisabled} value={group!}>
                        <label for={"inputIntra"+idx}>Gruppe {group!+1}</label>
                    </div>
                {/each}
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
    .blurred {
        filter: blur(5px);
    }
</style>