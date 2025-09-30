<script lang="ts">
    interface Props {
        yes(e: Event): void;
        no(e: Event): void;
    }
    let {yes,no}: Props = $props();
</script>

<div class="yesno flex:row">
    <button class="filled-btn green-btn" onclick={yes}>Ja</button>
    <button class="filled-btn red-btn" onclick={no}>Nein</button>
</div>

<style>
    .yesno {
        gap: 0.6em;
    }
</style>