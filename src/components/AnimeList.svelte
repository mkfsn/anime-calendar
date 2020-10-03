<script>
    import Anime from "./Anime/Anime.svelte";
    import {fetchAnime} from "../utils";
    import {animeStore} from "../store";
    export let show = true;
    $: $animeStore.loaded ? $animeStore.data : fetchAnime().then(data => {
        animeStore.update(o => {
            o.data = data;
            o.loaded = true;
            return o;
        })
    })
</script>

<style>
    .outer {
        background-color: #e2e2e2;
        display: none;
    }
    .outer.show {
        display: block;
    }
    .scrollable {
        width: 100%;
        overflow-x: scroll;
    }
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

<div class="outer" class:show={show}>
    {#if $animeStore.loaded}
        <div class="scrollable">
            <div class="container">
                {#each $animeStore.data as anime, i}
                    <Anime {anime} />
                {/each}
            </div>
        </div>
    {:else}
        <p class="loading">Loading ...</p>
    {/if}
</div>
