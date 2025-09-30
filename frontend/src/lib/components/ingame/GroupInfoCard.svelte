<script lang="ts">
    import { TIME_SLOTS } from "../../consts";
    import { stationsStore } from "../../stores";
    import type { Group } from "../../types";

    let {groups}: {groups: Group[]} = $props();
    let selected: number = $state(0);
    let group = $derived(groups[selected]);
    let stations = $derived(group.stations.map(idx => stationsStore.stations[idx]));
</script>

<div class="flex:column card">
    <h2>Gruppen Info</h2>
    <div class="flex:column">
        <select id="group-sel" bind:value={selected}>
            {#each groups as _, idx (idx)}
                <option value={idx}>Gruppe {idx + 1}</option>
            {/each}
        </select>
    </div>
    <div class="flex:column" style="text-align: left; padding: 0 0.4em">
        <span class="info-title">Punkte: {group.points}</span>
        <span class="info-title">Stationen:</span>
        <div class="flex:column" style="padding: 0.4em; gap: 0.4em;">
            {#each stations as st, idx (idx)}
                <div class="flex:column">
                    <span style="font-size: small;">{TIME_SLOTS[idx][0]} &ndash; {TIME_SLOTS[idx][1]}</span>
                    <strong>{st.name}</strong>
                    <span>{st.location}</span>
                </div>
            {/each}
        </div>
    </div>
</div>

<style>
    #group-sel {
        flex: 1;
        outline: none;
        background-color: transparent;
        font-size: inherit;
        border-top: 1px solid var(--main-fg-color);
        border-bottom: 1px solid var(--main-fg-color);
        border-left: none;
        border-right: none;
        color: inherit;
        padding: 0.2em;
        margin: 0.2em 0;
    }
    #group-sel option {
        color: black;
    }
</style>