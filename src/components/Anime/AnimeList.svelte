<script>
    import Anime from "./Anime.svelte";
    import {fetchAnime} from "../../utils";
    import {animeStore} from "../../stores";
    $: $animeStore.loaded ? $animeStore.data : fetchAnime().then(data => {
        animeStore.update(o => {
            o.data = data;
            o.loaded = true;
            return o;
        })
    })
</script>

{#if $animeStore.loaded}
    <div class="scrollable">
        <div class="container">
            {#each $animeStore.data as anime}
                <Anime {anime} />
            {/each}
        </div>
    </div>
{:else}
    <p class="loading">Loading ...</p>
{/if}

<style>
    .container {
        display: flex;
        flex-direction: row;
        float: left;
    }
    .loading {
        margin: 0;
        padding: 1em;
    }
</style>