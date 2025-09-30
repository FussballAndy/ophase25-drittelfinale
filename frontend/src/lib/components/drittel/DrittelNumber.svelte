<script lang="ts">
    interface Props {
        submit (e: Event, value: number): void;
    }

    let {submit}: Props = $props();

    let numberInput: string = $state("");
    function onChange() {
        numberInput = numberInput.replace(/[^0-9]/g, '')
    }
</script>

<div class="number-wrapper flex:row">
    <input type="text" name="answer-number" inputmode="numeric" placeholder="Zahl" autocomplete="off" bind:value={numberInput} oninput={onChange}>
    <button class="filled-btn submit" onclick={e => numberInput && submit(e, parseInt(numberInput))}>Abgeben</button>
</div>

<style>
    input[type="text"] {
        flex: 2;
        border: 1px solid black;
    }
    .number-wrapper {
        align-items: stretch;
        gap: 0.4em;
    }
    .submit {
        padding: 0.4em 0em;
        margin: 0;
    }
</style>