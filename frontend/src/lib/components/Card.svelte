<script lang="ts">
    let {background = false, children = undefined} = $props();
</script>

<div class="card" class:card-bg={background}>
    {@render children?.()}
</div>

<style>
    .card {
        display: flex;
        padding: 0.5em;
        border-radius: 0.6em;
        margin: 0.3em 0.1em;
        flex: 1;
    }
    .card-bg {
        background-color: rgb(255 255 255 / 5%);
    }
</style>