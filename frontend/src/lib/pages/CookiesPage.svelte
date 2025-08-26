<script lang="ts">
    import pageStore from "../pageStore";

    function acceptClick(e: Event) {
        e.preventDefault();
        pageStore.set(1);
    }
</script>

<p>Diese Seite nutzt Cookies um euch zu
identifizieren. Bitte nicht die Website-Daten
löschen oder Cookies zurücksetzen.</p>

<button class="accept" on:click={acceptClick}>Akzeptieren</button>

<style>
    .accept {
        border: none;
        padding: 0.5em 0.75em;
        background-color: hsl(120 100% 37%);
        color: inherit;
        border-radius: 0.25em;
        font-weight: bold;
        font-size: inherit;
        letter-spacing: 0.5px;
        
    }
    .accept:hover {
        background-color: hsl(120 100% 33%);
    }
    .accept:active {
        background-color: hsl(120 100% 25%);
    }
</style>