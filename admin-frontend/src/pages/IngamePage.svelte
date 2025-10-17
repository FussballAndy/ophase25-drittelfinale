<script lang="ts">
    import QuestionResult from "../components/QuestionResult.svelte";
    import { drittelReady, pageStore, questionActiveStore, questionResultStore, questionsStore, scalingStore, scoreStore } from "../lib/stores";
    import { wsSend } from "../lib/websocket";

    let allEntries = $derived($scoreStore.flatMap(x => [x[0], x[1], x[2]]))
    let studs = $derived(allEntries.filter(x => x == 1).length)
    let total = $derived(allEntries.filter(x => x !== undefined).length)

    let intraWinner = $state(-1)

    let makeFrontValue = $state("")

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

    function exportData(data: any) {
        const link = document.createElement('a');
        link.download = 'data.json';
        const blob = new Blob([JSON.stringify(data)], {type: 'application/json'});
        link.href = window.URL.createObjectURL(blob);
        link.click();
    }

    function exportScores(e: Event) {
        e.preventDefault();
        exportData($scoreStore)
    }

    function startQuestion(num: number) {
        $questionActiveStore = num;
        wsSend({
            kind: "next",
            data: num
        })
        setTimeout(() => {
            $questionActiveStore = null
        }, 30 * 1000)
    }

    function tryMakeFront(e: Event) {
        e.preventDefault()
        wsSend({
            kind: "front",
            data: makeFrontValue
        })
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
            <span>Toggle Front</span>
            <div class="flex:row">
                <input type="text" bind:value={makeFrontValue}>
                <button onclick={tryMakeFront}>Submit</button>
            </div>
        </div>
        <div class="flex:column" style="border: 1px solid white; padding: 0.3em;">
            {#each $questionsStore as question, idx (idx)}
                <div class="flex:row" style="gap: 0.6em;">
                    <span>{idx+1}. {question.prompt}</span>
                    <button onclick={() => startQuestion(idx)} disabled={$questionActiveStore !== null}>Starten</button>
                </div>
            {/each}
        </div>
        <div class="flex:column" style="text-align: left;">
            {#each $questionResultStore as result, idx (idx)}
                <div class="flex:column">
                    <span style="font-size: large;">Runde {idx+1}</span>
                    <div class="flex:row">
                        <QuestionResult submissions={Object.values(result).filter(x => !x.group)} name="Studies" question={idx} />
                        <QuestionResult submissions={Object.values(result).filter(x => x.group)} name="Tuts" question={idx} />
                    </div>
                    <button onclick={() => exportData(result)}>Export</button>
                </div>
            {/each}
        </div>
        <button onclick={() => exportData($questionResultStore)}>Export All</button>
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