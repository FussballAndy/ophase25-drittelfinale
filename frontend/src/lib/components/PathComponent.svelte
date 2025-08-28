<script lang="ts">
    import Card from "./Card.svelte";
    let {title, room, google, apple, start, end, prev_end} = $props();
    let now = Date.now();
    let isCurrent = prev_end <= now && now < end;
    let start_time = new Date(start);
    let end_time = new Date(end);
    let date_format = new Intl.DateTimeFormat("de-DE", {hour: "2-digit", minute: "2-digit"});
</script>

<Card background={isCurrent}>
    <div class="inner-wrapper flex:column">
        <div class="path-info flex:column">
            <span class="path-title">{title}</span>
            <span class="path-room">{room}</span>
            <span class="path-time">{date_format.format(start_time)} &ndash; {date_format.format(end_time)}</span>
        </div>
        {#if isCurrent}
            <div class="path-links flex:column">
                <span>Apple Karten: <a href={apple}>Link &#128279;</a></span>
                <span>Google Maps: <a href={google}>Link &#128279;</a></span>
            </div>
        {/if}
    </div>
</Card>

<style>
    .inner-wrapper {
        padding: 0.3em;
        flex: 1;
        gap: 1em;
    }
    .path-info {
        flex: 1;
    }
    .path-title {
        font-weight: 500;
        font-size: 1.1em;
    }
    .path-links {
        flex: 1;
        font-size: 0.9em;
    }
    a {
        color: hsl(210, 100%, 50%);
        text-decoration: none;
    }
</style>