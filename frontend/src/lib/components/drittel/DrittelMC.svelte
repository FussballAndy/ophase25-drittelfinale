<script lang="ts">
    interface Props {
        answers: string[];
        submit (e: Event, value: number): void;
    }
    let {answers, submit}: Props = $props();

</script>

<div class="mc flex:column">
    {#each answers as answer, idx (idx)}
        <button class="mc-answer" onclick={e => submit(e, idx)}>{answer}</button>
    {/each}
</div>

<style>
    .mc {
        gap: 0.6em;
        padding: 0 0.2em;
    }
    .mc-answer {
        background-color: transparent;
        border: 1px solid var(--main-fg-color);
        padding: 1em;
        font-size: inherit;
        color: inherit;
        border-radius: 0.5em;
        transition: background-color 0.1s ease-in-out;
    }
    .mc-answer:active {
        background-color: var(--main-fg-color);
        color: var(--main-bg-color)
    }
</style>