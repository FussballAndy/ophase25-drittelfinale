<script lang="ts">
    let {students, tutors} = $props();
</script>

<div class="flex:column card">
    <h2>Punktestand</h2>
    <div class="points">
        <strong>Studenten</strong>
        <strong>Tutoren</strong>
        <span>{students}</span>
        <span>{tutors}</span>
    </div>
</div>

<style>
    .points {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 0.2em 1em;
        font-size: large;
    }
</style>