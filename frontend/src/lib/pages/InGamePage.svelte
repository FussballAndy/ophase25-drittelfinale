<script lang="ts">
    import Card from "../components/Card.svelte";
    import TeamComponent from "../components/TeamComponent.svelte";

</script>

<div class="surr">
    <div class="team-wrapper">
        <span>Dein Team:</span>
        <TeamComponent name="1" showPlayers={false} isCurrent={true} />
    </div>
    <div class="path-wrapper">
        <span>Dein Pfad:</span>
        <Card background={true}>
            <div class="inner-wrapper">
                <div class="path-info">
                    <span class="path-title">Raketen</span>
                    <span class="path-room">S1|03 205</span>
                </div>
                <div class="path-more">
                    &gt;
                </div>
            </div>
        </Card>
    </div>
</div>

<style>
    .surr {
        display: flex;
        flex-direction: column;
        text-align: left;
    }
    .inner-wrapper {
        display: flex;
        flex-direction: row;
        padding: 0.3em;
        justify-content: space-between;
        align-items: center;
        gap: 8em;
        flex: 1;
    }
    .path-info {
        display: flex;
        flex-direction: column;
    }
    .team-wrapper {
        display: flex;
        flex-direction: column;
    }
    .path-title {
        font-weight: 500;
        font-size: 1.1em;
    }
    .path-more {
        display: flex;
        border: 1px solid var(--main-fg-color);
        border-radius: 100%;
        width: 2em;
        height: 2em;
        justify-content: center;
        align-items: center;
    }
</style>