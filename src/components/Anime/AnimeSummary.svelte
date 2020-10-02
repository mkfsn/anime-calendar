<script>
    export let anime;
</script>

<p class="title">
    {anime.Title}
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
    <div class="divider"></div>
    <div class="cast">
        {#if anime.Cast}
            {#each Array.from(Object.entries(anime.Cast)) as [key, value]}
                <b>{key}</b>
                <i style="padding-right: .5em">{value}</i>
            {/each}
        {/if}
    </div>
</div>

<style>
    .title {
        padding: .5em .25em;
        margin: 0;
        font-size: .7em;
        background-color: #8e8e8e;
        color: #eaeaea;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
    .row {
        border-top: 1px solid #eaeaea;
        display: flex;
        gap: .15em;
        padding: 0 .3em;
    }
    .divider {
        width: 1px;
        background-color: #cccccc;
    }
    .staff,
    .cast {
        width: 50%;
        font-size: .7em;
        display: -webkit-box;
        -webkit-line-clamp: 5;
        -webkit-box-orient: vertical;
        overflow: hidden;
        text-overflow: ellipsis;
        text-align: left;
    }
</style>