<script lang="ts">
    import PointsCard from "../components/ingame/PointsCard.svelte";
    import GroupPointsCard from "../components/ingame/GroupPointsCard.svelte";
    import GroupInfoCard from "../components/ingame/GroupInfoCard.svelte";
    import { userStore } from "../stores";
    import WinnerSubmission from "../components/ingame/WinnerSubmission.svelte";

</script>

<div class="surr flex:column">
    <PointsCard students="50" tutors="20" />
    <GroupPointsCard />
    {#if $userStore === null}
    <GroupInfoCard />
    {:else}
    <WinnerSubmission />
    {/if}
</div>

<style>
    .surr {
        gap: 0.5em;
    }
</style>