<script lang="ts">
    import { onMount } from "svelte";
    import DrittelMC from "../components/drittel/DrittelMC.svelte";
    import DrittelNumber from "../components/drittel/DrittelNumber.svelte";
    import DrittelYesNo from "../components/drittel/DrittelYesNo.svelte";
    import { wsStart, wsClose, questionStore, wsSubmitAnswer } from "../websocket";
    import { get as getStoreValue } from "svelte/store";

    onMount(() => {
        wsStart()
        return wsClose
    })

    /*let questions = [
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
    let curQ = $derived(questions[currentQuestion]);*/
    function submit(e: Event, value: number) {
        e.preventDefault()
        // submit using currentQuestion and value
        const question = getStoreValue(questionStore)
        if(question) {
            const questionNr = question.num;
            wsSubmitAnswer(questionNr, value)
            console.log(`Answered ${value} to Question ${questionNr}`)
        }
    }
</script>

<div class="surr flex:column">
    {#if $questionStore !== null}
        <div class="question-box">
            <span class="question">{$questionStore.prompt}</span>
        </div>
        <div class="answer-box flex:column">
            <span>Antwort:</span>
            <DrittelMC answers={$questionStore.answers} submit={submit} />
        </div>
    {:else}
        <span>Warte auf die nächste Frage...</span>
    {/if}
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