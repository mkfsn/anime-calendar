<script>
    import AnimeActions from "./Anime/AnimeActions.svelte";
    import AnimeSummary from "./Anime/AnimeSummary.svelte";
    import AnimeCover from "./Anime/AnimeCover.svelte";

    export let anime;
</script>

<style>
    .anime {
        margin-right: 1em;
        width: 160px;
        white-space: initial;
        height: 237px;
        margin-top: 1em;
        margin-bottom: 1em;
        background-color: white;
    }
    .anime:first-child {
        margin-left: 1em;
    }
</style>

<div class="anime">
    <AnimeCover {anime} />
    <AnimeSummary {anime} />
    <AnimeActions {anime} />
</div>