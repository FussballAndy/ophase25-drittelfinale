<script lang="ts">
    import { onMount } from "svelte";
    import DrittelMC from "../components/drittel/DrittelMC.svelte";
    import DrittelNumber from "../components/drittel/DrittelNumber.svelte";
    import DrittelYesNo from "../components/drittel/DrittelYesNo.svelte";

    onMount(() => {
        
    })

    let questions = [
        {
            text: "Ist die Erde flach?",
            type: "yesno"
        },
        {
            text: "Was ist die Antwort auf alles?",
            type: "number"
        },
        {
            text: "Wofür steht das F in TU Darmstadt?",
            type: "mc",
            answers: [
                "Freizeit", "Fortschritt", "Funktional"
            ]
        }
    ]
    let currentQuestion = $state(2);
    let curQ = $derived(questions[currentQuestion]);
    function submit(e: Event, value: boolean | number | string) {
        e.preventDefault()
        // submit using currentQuestion and value
        console.log(`Answered ${value} to Question ${currentQuestion}`)
    }
</script>

<div class="surr flex:column">
    <div class="question-box">
        <span class="question">{curQ.text}</span>
    </div>
    <div class="answer-box flex:column">
        <span>Antwort:</span>
        {#if curQ.type === "yesno"}
            <DrittelYesNo yes={e => submit(e,true)} no={e => submit(e,false)} />
        {:else if curQ.type === "number"}
            <DrittelNumber submit={submit} />
        {:else if curQ.type === "mc"}
            <DrittelMC answers={curQ.answers || []} submit={submit} />
        {/if}
    </div>
</div>

<style>
    .surr {
        gap: 1em;
    }
    .answer-box {
        text-align: left;
    }
    .question-box {
        padding: 2em 0.3em;
        border: 1px solid var(--main-fg-color);
        display: flex;
        flex: 1;
        justify-content: center;
        align-items: center;
    }
    .question {
        font-size: 1.1em;
        font-weight: 500;
    }
</style>