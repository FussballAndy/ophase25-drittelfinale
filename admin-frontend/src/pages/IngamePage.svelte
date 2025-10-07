<script lang="ts">
    import { drittelReady, pageStore, questionsStore, scalingStore, scoreStore } from "../lib/stores";

    let allEntries = $derived($scoreStore.flatMap(x => [x[0], x[1], x[2]]))
    let studs = $derived(allEntries.filter(x => x == 1).length)
    let total = $derived(allEntries.filter(x => x).length)

    let intraWinner = $state(-1)

    $effect(() => {
        let allIntra = $scoreStore.map(x => x[3])
        let counts = Array(25).fill(0)
        for(const c of allIntra) {
            if(c) {
                counts[c]++
            }
        }

        let max = counts[0];
        let maxIndex = 0;

        for (let i = 1; i < counts.length; i++) {
            if (counts[i] > max) {
                maxIndex = i;
                max = counts[i];
            }
        }
        intraWinner = maxIndex
    })

    function exportScores(e: Event) {
        e.preventDefault();
        const link = document.createElement('a');
        link.download = 'data.json';
        const blob = new Blob([JSON.stringify($scoreStore)], {type: 'application/json'});
        link.href = window.URL.createObjectURL(blob);
        link.click();
    }
</script>

<div class="flex:column surr">
    <table>
        <thead>
            <tr>
                <th>Station</th>
                <th>Iteration 1</th>
                <th>Iteration 2</th>
                <th>Iteration 3</th>
                <th>Intra</th>
            </tr>
        </thead>
        <tbody>
            {#each $scoreStore as row, idx (idx)}
                <tr>
                    <td>Station {idx+1}</td>
                    {#each row as cell, idx2 (idx2)}
                        <td>{cell}</td>
                    {/each}
                </tr>
            {/each}
        </tbody>
    </table>
    <button onclick={exportScores}>Export</button>
    <div class="flex:row result-wrap">
        <span>Studies: {studs}</span>
        <span>Tutoren: {total - studs}</span>
    </div>
    <div class="flex:column">
        <span>Intra: {intraWinner+1}</span>
    </div>
    <div class="flex:row" style="gap: 0.6em;">
        <span>Scaling:</span>
        <input type="number" bind:value={$scalingStore} step="1">
    </div>
    {#if $drittelReady}
        <div class="flex:column">
            <span>Aktuelle Frage:</span>
            <span>{$questionsStore[0].prompt}</span>
        </div>
    {/if}
</div>

<style>
    table, th, td {
        border: 1px solid white;
    }
    th, td {
        padding: 0.1em;
    }
    .surr {
        margin: 0 12em;
        gap: 1em;
        align-items: center;
    }
    .result-wrap {
        justify-content: center;
        align-items: center;
        gap: 1em;
    }
</style>