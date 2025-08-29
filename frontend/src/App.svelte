<script lang="ts">
    import "./app.css";
    import CookiesPage from "./lib/pages/CookiesPage.svelte";
    import DrittelPage from "./lib/pages/DrittelPage.svelte";
    import InGamePage from "./lib/pages/InGamePage.svelte";
    import UserPage from "./lib/pages/UserPage.svelte";
    import page from "./lib/pageStore";

    function advance(e: Event) {
        e.preventDefault();
        page.update(n => n + 1);
    }
</script>

<main>
    <h1 class="sticky-title">Geländespiel</h1>
    {#if $page == 0}
        <CookiesPage />
    {:else if $page == 1}
        <UserPage />
    {:else if $page == 2}
        <InGamePage />
    {:else if $page == 3}
        <DrittelPage />
    {:else}
        Unknown page {$page}
    {/if}
    <button onclick={advance}>Advance</button>
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
    }
</style>
