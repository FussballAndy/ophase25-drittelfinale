<script lang="ts">
    import { onMount } from "svelte";
    import "./app.css";
    import DrittelPage from "./lib/pages/DrittelPage.svelte";
    import InGamePage from "./lib/pages/InGamePage.svelte";
    import TutorLogin from "./lib/components/TutorLogin.svelte";
    import { userStore } from "./lib/stores";

    const INGAME_START_TIME = Date.UTC(2025, 9, 8, 15);
    const INGAME_END_TIME = Date.UTC(2025, 9, 8, 16, 20);
    const DRITTEL_START_TIME = Date.UTC(2025, 9, 8, 16, 38);
    const DRITTEL_END_TIME = Date.UTC(2025, 9, 8, 17);

    let currentTime = $state(DRITTEL_START_TIME);

    /*onMount(() => {
        let interval = setInterval(() => currentTime = Date.now(), 30 * 1000);
        return () => clearInterval(interval)
    });*/
</script>

<main>
    <h1 class="sticky-title">Geländespiel</h1>
    {#if currentTime < INGAME_START_TIME}
        Das Geländespiel startet bald!
        Noch {Math.ceil((INGAME_START_TIME - currentTime) / (60 * 1000))} Minuten!
    {:else if INGAME_START_TIME <= currentTime && currentTime < INGAME_END_TIME}
        <InGamePage />
    {:else if INGAME_END_TIME <= currentTime && currentTime < DRITTEL_START_TIME}
        Geht gleich weiter.
    {:else if DRITTEL_START_TIME <= currentTime && currentTime < DRITTEL_END_TIME}
        <DrittelPage />
    {:else}
        Das Geländespiel ist nun vorbei! &#x1F44B; <br> Es hat uns viel Spaß gemacht. 
    {/if}
    {#if $userStore === null}
        <TutorLogin />
    {/if}
</main>

<style>
    .sticky-title {
        position: sticky;
        top: 20px;
        /* background-color: var(--main-bg-color); */
        backdrop-filter: blur(5px);
    }
    main {
        padding: 2em 0px;
        display: flex;
        flex-direction: column;
        gap: 2em;
    }
</style>
