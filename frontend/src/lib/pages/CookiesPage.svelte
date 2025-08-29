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

<button class="filled-btn accept" on:click={acceptClick}>Akzeptieren</button>

<style>
    .accept {
        background-color: hsl(120 100% 37%);
        color: inherit;
        font-weight: bold;
        letter-spacing: 0.5px;
        
    }
    .accept:hover {
        background-color: hsl(120 100% 33%);
    }
    .accept:active {
        background-color: hsl(120 100% 25%);
    }
</style>