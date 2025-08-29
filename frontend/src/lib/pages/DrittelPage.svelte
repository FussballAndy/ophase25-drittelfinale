<script lang="ts">
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
    let currentQuestion = $state(2)
    function submit(e: Event, value: boolean | number | string) {
        e.preventDefault()
        // submit using currentQuestion and value
        console.log(`Answered ${value} to Question ${currentQuestion}`)
    }

    let numberInput: string = $state("");

    function onChange() {
        numberInput = numberInput.replace(/[^0-9]/g, '')
    }
</script>

<div class="surr flex:column">
    <div class="question-box">
        <span class="question">{questions[currentQuestion].text}</span>
    </div>
    <div class="answer-box flex:column">
        <span>Antwort:</span>
        {#if questions[currentQuestion].type === "yesno"}
            <div class="yesno flex:row">
                <button class="filled-btn green-btn" onclick={e => submit(e, true)}>Ja</button>
                <button class="filled-btn red-btn" onclick={e => submit(e, false)}>Nein</button>
            </div>
        {:else if questions[currentQuestion].type === "number"}
            <div class="number-wrapper flex:row">
                <input type="text" name="answer-number" inputmode="numeric" placeholder="Zahl" autocomplete="off" bind:value={numberInput} oninput={onChange}>
                <button class="filled-btn submit" onclick={e => numberInput && submit(e, parseInt(numberInput))}>Abgeben</button>
            </div>
        {:else if questions[currentQuestion].type === "mc"}
            <div class="mc flex:column">
                {#each questions[currentQuestion].answers || [] as answer, idx (idx)}
                    <button class="mc-answer" onclick={e => submit(e, idx)}>{answer}</button>
                {/each}
            </div>
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
    input[type="text"] {
        flex: 2;
        font-size: inherit;
        margin: 0;
        padding: 0;
        border-radius: 0.25em;
        border: 1px solid black;
        padding: 0 0.1em;
        outline: none;
        text-align: center;
    }
    .number-wrapper {
        align-items: stretch;
        gap: 0.4em;
    }
    .question {
        font-size: 1.1em;
        font-weight: 500;
    }
    .yesno {
        gap: 0.6em;
    }
    .submit {
        padding: 0.4em 0em;
        margin: 0;
    }
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