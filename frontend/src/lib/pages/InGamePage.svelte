<script lang="ts">
    import { onDestroy } from "svelte";
    import PathComponent from "../components/PathComponent.svelte";
    import TeamComponent from "../components/TeamComponent.svelte";

    let path = [{
        title: "Raketen",
        room: "S1|03 205",
        google: "https://maps.app.goo.gl/MNNjYW1AUyPqFXUFA",
        apple: "https://maps.apple/p/5GWcLBK9CdcD6f",
        start: Date.UTC(2025, 7, 28, 20, 15, 0, 0),
        end: Date.UTC(2025, 7, 28, 20, 45, 0, 0),
        prev_end: 0,
    },
    {
        title: "Musikraten",
        room: "S1|03 205",
        google: "https://maps.app.goo.gl/MNNjYW1AUyPqFXUFA",
        apple: "https://maps.apple/p/5GWcLBK9CdcD6f",
        start: Date.UTC(2025, 7, 28, 20, 50, 0, 0),
        end: Date.UTC(2025, 7, 28, 21, 20, 0, 0),
        prev_end: Date.UTC(2025, 7, 28, 20, 45, 0, 0),
    },
    {
        title: "Musikraten",
        room: "S1|03 205",
        google: "https://maps.app.goo.gl/MNNjYW1AUyPqFXUFA",
        apple: "https://maps.apple/p/5GWcLBK9CdcD6f",
        start: Date.UTC(2025, 7, 28, 20, 50, 0, 0),
        end: Date.UTC(2025, 7, 28, 21, 20, 0, 0),
        prev_end: Date.UTC(2025, 7, 28, 20, 45, 0, 0),
    },
    {
        title: "Musikraten",
        room: "S1|03 205",
        google: "https://maps.app.goo.gl/MNNjYW1AUyPqFXUFA",
        apple: "https://maps.apple/p/5GWcLBK9CdcD6f",
        start: Date.UTC(2025, 7, 28, 20, 50, 0, 0),
        end: Date.UTC(2025, 7, 28, 21, 20, 0, 0),
        prev_end: Date.UTC(2025, 7, 28, 20, 45, 0, 0),
    }]

    let ms = 1000 * 60
    let time = $state(Date.now())
    let inter = setInterval(() => time = Date.now(), ms);
    onDestroy(() => clearInterval(inter))
</script>

<div class="surr">
    <div class="team-wrapper">
        <span>Dein Team:</span>
        <TeamComponent name="1" showPlayers={false} isCurrent={true} />
    </div>
    <div class="path-wrapper">
        <span>Dein Pfad:</span>
        <div class="path">
            {#each path as p, idx (idx)}
                <PathComponent {...p} />
            {/each}
        </div>
    </div>
</div>

<style>
    .surr {
        display: flex;
        flex-direction: column;
        text-align: left;
    }
    .path {
        display: flex;
        flex-direction: column;
    }
    .team-wrapper {
        display: flex;
        flex-direction: column;
    }
</style>