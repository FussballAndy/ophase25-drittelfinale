<script lang="ts">
    import { requestNameStore, requestDataStore } from "../lib/stores";
    import { wsSend } from "../lib/websocket";

    let file: FileList | null = $state(null);

    function uploadFile(e: Event) {
        e.preventDefault();
        if(file) {
            let f = file.item(0)
            f?.text().then(t => JSON.parse(t)).then(d => $requestDataStore = d).then(_ => file = null);
        }
    }

    function exportFile(e: Event) {
        e.preventDefault()
        const link = document.createElement('a');
        link.download = 'data.json';
        const blob = new Blob([JSON.stringify($requestDataStore)], {type: 'application/json'});
        link.href = window.URL.createObjectURL(blob);
        link.click();
    }
    function submit(e: Event) {
        e.preventDefault()
        wsSend($requestDataStore)
        requestDataStore.set([])
        requestNameStore.set("")
    }
</script>

<div class="window flex:column">
    <h2 style="margin: 0.2em;">Aktuelle eingabe: {$requestNameStore}</h2>
    <div class="flex:column" style="padding: 0 0.5em;">
        {#each $requestDataStore as entry, idx (idx)}
            <div class="flex:column entry" style="text-align: left;">
                {#each Object.keys(entry) as name, idx2 (idx2)}
                    <div><strong>{name}:</strong> {entry[name]}</div>
                {/each}
            </div>
        {/each}
    </div>
    <div class="flex:row flex:centered" style="gap: 0.3em; padding: 0.3em;">
        <input type="file" accept="application/json" bind:files={file}>
        <button onclick={uploadFile}>Upload JSON</button>
        <button onclick={exportFile}>Export JSON</button>
        <button onclick={submit}>Submit</button>
    </div>
</div>

<style>
    .window {
        background-color: #343434;
        margin: 1em;
        border-radius: 1em;
    }
    .entry {
        border-top: 1px solid black;
        border-bottom: 1px solid black;
        padding: 0.1em;
    }
    .entry:nth-child(odd) {
        background-color: #20202010;
    }
</style>