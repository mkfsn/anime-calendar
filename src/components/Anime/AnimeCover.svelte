<script>
    import AnimeActionStar from "./AnimeActionStar.svelte";
    export let anime;
</script>

{#if anime.OpenGraph && anime.OpenGraph.Images && anime.OpenGraph.Images.length > 0}
    <div class="cover" style="background-image: url('{anime.OpenGraph.Images[0].URL}')">
        <div class="star">
            <AnimeActionStar {anime} />
        </div>
    </div>
{:else}
    <div class="cover broken" style="background-color: rgb(239, 239, 239);">
        <div class="star">
            <AnimeActionStar {anime} />
        </div>
        <svg class="svg-icon" viewBox="0 0 20 20">
            <path d="M18.555,15.354V4.592c0-0.248-0.202-0.451-0.45-0.451H1.888c-0.248,0-0.451,0.203-0.451,0.451v10.808c0,0.559,0.751,0.451,0.451,0.451h16.217h0.005C18.793,15.851,18.478,14.814,18.555,15.354 M2.8,14.949l4.944-6.464l4.144,5.419c0.003,0.003,0.003,0.003,0.003,0.005l0.797,1.04H2.8z M13.822,14.949l-1.006-1.317l1.689-2.218l2.688,3.535H13.822z M17.654,14.064l-2.791-3.666c-0.181-0.237-0.535-0.237-0.716,0l-1.899,2.493l-4.146-5.42c-0.18-0.237-0.536-0.237-0.716,0l-5.047,6.598V5.042h15.316V14.064z M12.474,6.393c-0.869,0-1.577,0.707-1.577,1.576s0.708,1.576,1.577,1.576s1.577-0.707,1.577-1.576S13.343,6.393,12.474,6.393 M12.474,8.645c-0.371,0-0.676-0.304-0.676-0.676s0.305-0.676,0.676-0.676c0.372,0,0.676,0.304,0.676,0.676S12.846,8.645,12.474,8.645"></path>
        </svg>
    </div>
{/if}

<style>
    .star {
        position: absolute;
        top: .3em;
        left: .3em;
        display: inline-block;
    }
    .cover {
        height: 90px;
        background-size: cover;
        background-repeat: no-repeat;
        background-position: center;
        width: 160px;
        position: relative;
    }
    .cover.broken {
        display: grid;
        justify-content: center;
        align-content: center;
    }
    .cover.broken > svg {
        width: 50px;
        height: 50px;
        fill: #9f9f9f;
    }
</style>