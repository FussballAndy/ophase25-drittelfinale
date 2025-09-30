<script lang="ts">
    import PointsCard from "../components/ingame/PointsCard.svelte";
    import GroupPointsCard from "../components/ingame/GroupPointsCard.svelte";
    import GroupInfoCard from "../components/ingame/GroupInfoCard.svelte";

    let groups = [{
        points: 0,
        stations: [0, 0, 0]
    },{
        points: 1,
        stations: []
    }];

</script>

<div class="surr flex:column">
    <PointsCard students="50" tutors="20" />
    <GroupPointsCard />
    <GroupInfoCard groups={groups} />
</div>

<style>
    .surr {
        gap: 0.5em;
    }
</style>