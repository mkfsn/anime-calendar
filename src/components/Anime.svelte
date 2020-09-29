<script>
    export let anime;
</script>

<style>
    .anime {
        margin-right: 1em;
        width: 320px;
        white-space: initial;
        height: 420px;
        margin-top: 1em;
        margin-bottom: 1em;
        background-color: white;
    }
    .cover {
        height: 180px;
        background-size: cover;
        background-repeat: no-repeat;
        background-position: center;
        width: 320px;
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
    .title {
        padding: .5em 0;
        margin: 0;
        background-color: #8e8e8e;
        color: #eaeaea;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
    .title > a:link {
        color: #eaeaea;
        text-decoration: none;
    }
    .title > a:visited {
        color: #eaeaea;
        text-decoration: none;
    }
    .title > a:hover {
        color: #eaeaea;
        /*text-decoration: underline;*/
        border-bottom: 1px solid #eaeaea;
    }
    .title > a:active {
        color: #eaeaea;
        /*text-decoration: underline;*/
        border-bottom: 1px solid #eaeaea;
    }
    .title svg {
        height: 16px;
        width: 16px;
        line-height: 16px;
        fill: #eaeaea;
        vertical-align: text-top;
    }
    .row {
        display: flex;
        gap: .3em;
        padding: 0 .3em;
    }
    .staff,
    .cast {
        width: 50%;
        font-size: .7em;
        display: -webkit-box;
        -webkit-line-clamp: 9;
        -webkit-box-orient: vertical;
        overflow: hidden;
        text-overflow: ellipsis;
        text-align: left;
    }
</style>

<div class="anime">
    {#if anime.OpenGraph && anime.OpenGraph.Images && anime.OpenGraph.Images.length > 0}
        <div class="cover" style="background-image: url('{anime.OpenGraph.Images[0].URL}')">
        </div>
    {:else}
        <div class="cover broken" style="background-color: rgb(239, 239, 239);">
            <svg class="svg-icon" viewBox="0 0 20 20">
                <path d="M18.555,15.354V4.592c0-0.248-0.202-0.451-0.45-0.451H1.888c-0.248,0-0.451,0.203-0.451,0.451v10.808c0,0.559,0.751,0.451,0.451,0.451h16.217h0.005C18.793,15.851,18.478,14.814,18.555,15.354 M2.8,14.949l4.944-6.464l4.144,5.419c0.003,0.003,0.003,0.003,0.003,0.005l0.797,1.04H2.8z M13.822,14.949l-1.006-1.317l1.689-2.218l2.688,3.535H13.822z M17.654,14.064l-2.791-3.666c-0.181-0.237-0.535-0.237-0.716,0l-1.899,2.493l-4.146-5.42c-0.18-0.237-0.536-0.237-0.716,0l-5.047,6.598V5.042h15.316V14.064z M12.474,6.393c-0.869,0-1.577,0.707-1.577,1.576s0.708,1.576,1.577,1.576s1.577-0.707,1.577-1.576S13.343,6.393,12.474,6.393 M12.474,8.645c-0.371,0-0.676-0.304-0.676-0.676s0.305-0.676,0.676-0.676c0.372,0,0.676,0.304,0.676,0.676S12.846,8.645,12.474,8.645"></path>
            </svg>
        </div>
    {/if}
    <p class="title">
        <a target="_blank" href={anime.Website}>
            {anime.Title}
            <svg class="svg-icon" viewBox="0 0 20 20">
                <path d="M16.469,8.924l-2.414,2.413c-0.156,0.156-0.408,0.156-0.564,0c-0.156-0.155-0.156-0.408,0-0.563l2.414-2.414c1.175-1.175,1.175-3.087,0-4.262c-0.57-0.569-1.326-0.883-2.132-0.883s-1.562,0.313-2.132,0.883L9.227,6.511c-1.175,1.175-1.175,3.087,0,4.263c0.288,0.288,0.624,0.511,0.997,0.662c0.204,0.083,0.303,0.315,0.22,0.52c-0.171,0.422-0.643,0.17-0.52,0.22c-0.473-0.191-0.898-0.474-1.262-0.838c-1.487-1.485-1.487-3.904,0-5.391l2.414-2.413c0.72-0.72,1.678-1.117,2.696-1.117s1.976,0.396,2.696,1.117C17.955,5.02,17.955,7.438,16.469,8.924 M10.076,7.825c-0.205-0.083-0.437,0.016-0.52,0.22c-0.083,0.205,0.016,0.437,0.22,0.52c0.374,0.151,0.709,0.374,0.997,0.662c1.176,1.176,1.176,3.088,0,4.263l-2.414,2.413c-0.569,0.569-1.326,0.883-2.131,0.883s-1.562-0.313-2.132-0.883c-1.175-1.175-1.175-3.087,0-4.262L6.51,9.227c0.156-0.155,0.156-0.408,0-0.564c-0.156-0.156-0.408-0.156-0.564,0l-2.414,2.414c-1.487,1.485-1.487,3.904,0,5.391c0.72,0.72,1.678,1.116,2.696,1.116s1.976-0.396,2.696-1.116l2.414-2.413c1.487-1.486,1.487-3.905,0-5.392C10.974,8.298,10.55,8.017,10.076,7.825"></path>
            </svg>
        </a>
    </p>
    <div class="row">
        <div class="staff">
            {#if anime.Staff}
            {#each Array.from(Object.entries(anime.Staff)) as [key, value]}
                <b>{key}</b>
                <i style="padding-right: .5em">{value}</i>
            {/each}
            {/if}
        </div>
        <div class="cast">
            {#if anime.Cast}
            {#each Array.from(Object.entries(anime.Cast)) as [key, value]}
                <b>{key}</b>
                <i style="padding-right: .5em">{value}</i>
            {/each}
            {/if}
        </div>
    </div>
</div>